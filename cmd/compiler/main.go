package main

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"

	"github.com/richardzmija/fusion-compiler/internal/ast"
	"github.com/richardzmija/fusion-compiler/internal/parser"
	"github.com/richardzmija/fusion-compiler/internal/semantic"

	"github.com/richardzmija/fusion-compiler/internal/parser/parsererror"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: compiler <file.c>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	inputBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %q: %v\n", filePath, err)
		os.Exit(1)
	}

	// Create the ANTLR stream from the read file.
	antlrInput := antlr.NewInputStream(string(inputBytes))

	// Create the lexer and parser.
	lexer := parser.NewCLexer(antlrInput)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewCParser(tokenStream)

	// Add a custom error listener.
	p.RemoveErrorListeners()
	listener := parsererror.NewErrorListener()
	p.AddErrorListener(listener)

	// Use parser for the top-level rule.
	parseTree := p.Program()

	if listener.HasErrors() {
		fmt.Fprintln(os.Stderr, "Parsing failed with the following errors:")
		for _, parseErr := range listener.GetErrors() {
			fmt.Fprintln(os.Stderr, "  ", parseErr)
		}
		fmt.Fprintf(os.Stderr, "Parsing failed with %d error(s).\n", len(listener.GetErrors()))
		os.Exit(1)
	}

	fmt.Println("Parsing completed successfully without syntax errors.")

	// Build the AST.
	builder := ast.NewASTBuilder()
	astResult := parseTree.Accept(builder)
	prog, ok := astResult.(*ast.Program)
	if !ok {
		fmt.Fprintln(os.Stderr, "Internal error: ASTBuilder did not return *ast.Program.")
		os.Exit(1)
	}

	fmt.Println("AST has been successfully built.")
	fmt.Println("AST representation:")
	fmt.Println(ast.StringifyProgram(prog))

	// Perform semantic analysis.
	analyzer := semantic.NewAnalyzer()
	semanticErrors := analyzer.Analyze(prog)
	if len(semanticErrors) > 0 {
		fmt.Fprintln(os.Stderr, "Semantic errors were found:")
		for _, semErr := range semanticErrors {
			fmt.Fprintln(os.Stderr, "  ", semErr)
		}
		os.Exit(1)
	}

	fmt.Println("No semantic errors found. Compilation may proceed.")
}

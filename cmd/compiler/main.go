package main

import (
	"fmt"
	"os"

	"github.com/richardzmija/fusion-compiler/internal/ast"
	"github.com/richardzmija/fusion-compiler/internal/parser"
	"github.com/richardzmija/fusion-compiler/internal/semantic"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	// Input file path must be provided.
	input, err := antlr.NewFileStream("")
	if err != nil {
		fmt.Println("Failed to read input file:", err)
		os.Exit(1)
	}

	lexer := parser.NewCLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewCParser(stream)

	// Parse tree
	tree := p.Program()

	// Build AST
	builder := ast.NewASTBuilder()
	prog := tree.Accept(builder).(*ast.Program)

	// Semantic analysis
	analyzer := semantic.NewAnalyzer()
	errors := analyzer.Analyze(prog)
	if len(errors) > 0 {
		for _, e := range errors {
			fmt.Println("Semantic error:", e)
		}
		os.Exit(1)
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/antlr4-go/antlr/v4"

	"github.com/richardzmija/fusion-compiler/internal/ast"
	"github.com/richardzmija/fusion-compiler/internal/codegen"
	"github.com/richardzmija/fusion-compiler/internal/parser"
	"github.com/richardzmija/fusion-compiler/internal/parser/parsererror"
	"github.com/richardzmija/fusion-compiler/internal/semantic"

	llvm "tinygo.org/x/go-llvm"
)

func main() {
	outputPathPtr := flag.String("o", "", "Specify the path of the output file")

	// Separate flags from positional arguments.
	var inputFiles []string
	for _, arg := range os.Args[1:] {
		if arg == "-o" {
			break
		}
		inputFiles = append(inputFiles, arg)
	}

	flag.CommandLine.Parse(os.Args[1+len(inputFiles):])

	if len(inputFiles) == 0 {
		log.Fatalln("Error: No input file specified. Usage: compiler <input_file> [-o <output_file>]")
	}

	inputPath := inputFiles[0]
	outputPath := "a"
	if *outputPathPtr != "" {
		outputPath = *outputPathPtr
	}

	// Extract directory from output file path.
	outputDir := filepath.Dir(outputPath)
	if outputDir == "." {
		absPath, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error: Failed to get current working directory: %v\n", err)
		}
		outputDir = absPath
	}

	// Ensure that the output directory exists.
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatalf("Error: Failed to create output directory %q: %v\n", outputDir, err)
	}

	bitCodeFilePath := filepath.Join(outputDir, "output.bc")
	irFilePath := filepath.Join(outputDir, "output.ll")

	inputBytes, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %q: %v\n", inputPath, err)
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

	// Generate the LLVM IR.
	codeGenerator := codegen.NewCodeGenerator("fusion-module")
	llvmModule := codeGenerator.Generate(prog)

	// Verify the generated IR.
	if err := llvm.VerifyModule(llvmModule, llvm.ReturnStatusAction); err != nil {
		fmt.Fprintf(os.Stderr, "LLVM module verification failed:\n%v\n", err)
		os.Exit(1)
	}

	// Write module to LLVM IR file.
	ir := llvmModule.String()
	err = os.WriteFile(irFilePath, []byte(ir), 0644)
	if err != nil {
		log.Fatalf("Error: Failed to write LLVM IR file: %v\n", err)
	}

	// Write module to bitcode file.
	bitCodeFile, err := os.Create(bitCodeFilePath)
	if err != nil {
		log.Fatalf("Error: Failed to create bitcode file: %v\n", err)
	}
	defer bitCodeFile.Close()

	fmt.Printf("Writing bitcode to %s...\n", bitCodeFile.Name())
	if err := llvm.WriteBitcodeToFile(llvmModule, bitCodeFile); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write bitcode: %v\n", err)
		os.Exit(1)
	}

	// Invoke clang to compile the bitcode into an executable. Here we pass
	// optimization flags and link against the C standard library which is automatic.
	fmt.Printf("Compiling bitcode with clang...\n")
	args := []string{
		"-o", outputPath,
		"-O3",
		bitCodeFilePath,
	}
	cmd := exec.Command("clang", args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute clang: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Compilation successful!\n")
}

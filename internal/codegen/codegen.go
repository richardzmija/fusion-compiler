package codegen

import (
	"github.com/richardzmija/fusion-compiler/internal/ast"
	llvm "tinygo.org/x/go-llvm"
)

// CodeGenerator translates high-level AST constructs into a lower-level,
// language-agnostic LLVM IR.
type CodeGenerator struct {
	// module is a central data structure in LLVM IR that encapsulates all
	// the information related to a single unit of code being compiled. It acts
	// as a container for functions, global variables, and other necessary components
	// required to represent a complete program or a given translation unit.
	module llvm.Module
	// builder is an object that exposes an interface for programatically creating
	// LLVM IR. It maintains the current insertion point within function's basic
	// blocks and provides methods for creating various types of instructions.
	builder llvm.Builder
	// context is an object that encapsulates and manages the unique instances of
	// LLVM's core data structures such as types, constants, and other immutable
	// entities.
	context llvm.Context
	// namedValues is a symbol table that keeps track of variable names and their
	// corresponding LLVM values within the current scope. It associates each variable
	// name with its allocated memory location.
	namedValues map[string]llvm.Value
	// printf is an LLVM represenatation of the external printf function. It is used to
	// make it possible to interoperate with existing C libraries to allow writing
	// to the standard output.
	printf llvm.Value
}

// NewCodeGenerator creates a new code generator with an empty module.
func NewCodeGenerator(moduleName string) *CodeGenerator {
	context := llvm.NewContext()
	module := context.NewModule(moduleName)
	builder := context.NewBuilder()

	codeGenerator := &CodeGenerator{
		module:      module,
		builder:     builder,
		context:     context,
		namedValues: make(map[string]llvm.Value),
	}

	codeGenerator.declarePrintf()

	return codeGenerator
}

// declarePrintf adds a function named printf to the module encapsulated
// in the code generator and stores a reference to it in the code generator.
func (c *CodeGenerator) declarePrintf() {
	printfType := llvm.FunctionType(
		c.context.Int32Type(),
		[]llvm.Type{llvm.PointerType(c.context.Int8Type(), 0)},
		true,
	)
	c.printf = llvm.AddFunction(c.module, "printf", printfType)
}

// Generate takes the root node of the AST and translates the entire
// AST into a set of constructs in the LLVM module.
func (c *CodeGenerator) Generate(program *ast.Program) llvm.Module {
	for _, function := range program.Functions {
		c.generateFunction(function)
	}

	return c.module
}

// generateFunction translates a FunctionDefinition node of the AST into an LLVM function
// and adds it to the module.
func (c *CodeGenerator) generateFunction(function *ast.FunctionDefinition) llvm.Value {
	// All parameter types are int in this subset.
	parameterTypes := make([]llvm.Type, len(function.Parameters))
	for i := range function.Parameters {
		parameterTypes[i] = c.context.Int32Type()
	}
	// A functions in this subset have return type int.
	returnType := c.context.Int32Type()

	// Build the function type and add to module.
	functionType := llvm.FunctionType(returnType, parameterTypes, false)
	llvmFunction := llvm.AddFunction(c.module, function.Name, functionType)

	// Create a new basic block named 'entry' within the newly added LLVM function
	// and set the builder insertion point to the end of this block.
	entryBasicBlock := c.context.AddBasicBlock(llvmFunction, "entry")
	c.builder.SetInsertPointAtEnd(entryBasicBlock)

	// At this point parameters could be stored in the symbol table and allocated
	// before we start processing the function body. For now this step is left out
	// as only the most basic functionality is implemented now.

	// Generate LLVM code for the function body.
	c.generateBlockStatement(function.Body)

	// It is assumed that every function that has a non-void return type
	// has a correct return statement. This should be handled by the semantic
	// analysis stage.
	return llvmFunction
}

func (c *CodeGenerator) generateBlockStatement(blockStatement *ast.BlockStatement) {
	// To be implemented.
}

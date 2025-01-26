package codegen

import (
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
		module: module,
		builder: builder,
		context: context,
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


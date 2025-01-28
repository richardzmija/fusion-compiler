package codegen

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"

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
	// name with its allocated memory location. Note: the way that namedValues is used
	// does not support lexical scopes in code generation.
	namedValues map[string]llvm.Value
	// stringLiterals tracks global string literals that were created in the LLVM module
	// and allows code generation methods to reuse these global strings for efficiency.
	stringLiterals map[string]llvm.Value
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
		module:         module,
		builder:        builder,
		context:        context,
		namedValues:    make(map[string]llvm.Value),
		stringLiterals: make(map[string]llvm.Value),
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

// generateBlockStatement translates a BlockStatement AST node into LLVM IR that is
// inserted into the basic block representing the entry point of the function's body
// and further linked basic blocks.
func (c *CodeGenerator) generateBlockStatement(blockStatement *ast.BlockStatement) {
	for _, declaration := range blockStatement.Declarations {
		c.generateDeclaration(declaration)
	}

	for _, statement := range blockStatement.Statements {
		c.generateStatement(statement)
	}
}

// generateDeclaration allocates memory on stack for the local variables
// and optionally stores the initializer at that memory location.
func (c *CodeGenerator) generateDeclaration(declaration *ast.Declaration) {
	for i, name := range declaration.Names {
		// Allocate memory on the stack for the local variable and store a pointer
		// to this memory location in the symbol table.
		allocation := c.builder.CreateAlloca(c.context.Int32Type(), name)
		c.namedValues[name] = allocation

		if declaration.Initializers[i] != nil {
			// If the initializer is provided, generate the LLVM value representing it
			// and store it at that memory location.
			initializationValue := c.generateExpression(declaration.Initializers[i])
			c.builder.CreateStore(initializationValue, allocation)
		}
	}
}

// generateStatement converts an AST statement node into its corresponding
// LLVM IR representation.
func (c *CodeGenerator) generateStatement(statement ast.Statement) {
	switch s := statement.(type) {
	case *ast.ExpressionStatement:
		if s.ContainedExpression != nil {
			c.generateExpression(s.ContainedExpression)
		} else {
			log.Fatal("Error: nil detected in an expression statement in codegen.")
		}

	case *ast.ReturnStatement:
		if s.ReturnValue != nil {
			val := c.generateExpression(s.ReturnValue)
			c.builder.CreateRet(val)
		} else {
			log.Printf("Warning: Return statement without return value detected. Defaulting to return value 0...")
			c.builder.CreateRet(llvm.ConstInt(c.context.Int32Type(), 0, false))
		}

	case *ast.PrintfStatement:
		c.generatePrintfStatement(s)

	case *ast.BlockStatement:
		// Generate a nested block inside of this compound statement.
		c.generateBlockStatement(s)

	default:
		log.Fatal("Fatal Error: Unhandled statement in codegen.")
	}
}

// getOrCreateGlobalString returns a pointer to the first element of the global string
// storing the provided string literal.
func (c *CodeGenerator) getOrCreateGlobalString(stringLiteral string) llvm.Value {
	if pointer, found := c.stringLiterals[stringLiteral]; found {
		return pointer
	}

	// Generate a unique name.
	hash := sha1.Sum([]byte(stringLiteral))
	hashString := hex.EncodeToString(hash[:])
	globalName := fmt.Sprintf(".str_%s", hashString)

	// Create a global string for the format. CreateGlobalString defines
	// a global array [N x i8] containing the format string. It returns an LLVM Value
	// referencing that global array.
	globalArray := c.builder.CreateGlobalString(stringLiteral, globalName)

	// These global strings are local to this LLVM module. Setting internal
	// linkage allows LLVM to perform additional optimizations.
	globalArray.SetLinkage(llvm.PrivateLinkage)

	// Get a pointer to the first character of this array.
	zero := llvm.ConstInt(c.context.Int32Type(), 0, false)
	pointer := c.builder.CreateInBoundsGEP(
		globalArray.Type().ElementType(),
		globalArray,
		[]llvm.Value{zero, zero},
		"strptr",
	)

	// Store the pointer for future access.
	c.stringLiterals[stringLiteral] = pointer

	return pointer
}

// generatePrintfStatement generates LLVM IR for a call to the external printf function.
func (c *CodeGenerator) generatePrintfStatement(printfStatement *ast.PrintfStatement) {
	formatPointer := c.getOrCreateGlobalString(printfStatement.Format)

	// Create the argument list.
	argumentValues := []llvm.Value{formatPointer}
	for _, expression := range printfStatement.Arguments {
		val := c.generateExpression(expression)
		argumentValues = append(argumentValues, val)
	}

	// Call external printf function and ignore the result.
	call := c.builder.CreateCall(c.printf.GlobalValueType(), c.printf, argumentValues, "printfcall")
	_ = call
}

// generateExpression converts an AST expression node into its corresponding
// LLVM IR value.
func (c *CodeGenerator) generateExpression(expression ast.Expression) llvm.Value {
	switch e := expression.(type) {

	case *ast.Literal:
		if e.Type == ast.IntLiteral {
			intValue, err := strconv.ParseInt(e.Value, 10, 64)
			if err != nil {
				log.Fatalf("Error: Cannot parse integer literal '%s': %v.", e.Value, err)
			}
			return llvm.ConstInt(
				c.context.Int32Type(),
				uint64(intValue),
				true, // sign-extend
			)
		} else if e.Type == ast.StringLiteral {
			return c.getOrCreateGlobalString(e.Value)
		}
		log.Fatalf("Error: Unknown literal type %v.", e.Type)

	case *ast.VariableExpression:
		variableAddress, exists := c.namedValues[e.Name]
		if !exists {
			log.Fatalf("Error: Variable '%s' not found in 'namedValues'. Semantic checking is incorrect.", e.Name)
		}
		return c.builder.CreateLoad(c.context.Int32Type(), variableAddress, e.Name)

	case *ast.BinaryExpression:
		left := c.generateExpression(e.Left)
		right := c.generateExpression(e.Right)
		switch e.Operator {
		case "+":
			return c.builder.CreateAdd(left, right, "addtmp")
		case "-":
			return c.builder.CreateSub(left, right, "subtmp")
		case "*":
			return c.builder.CreateMul(left, right, "multmp")
		case "/":
			return c.builder.CreateSDiv(left, right, "divtmp")

		default:
			log.Fatalf("Error: unknown binary operator '%s' in codegen.", e.Operator)
		}

	case *ast.UnaryExpression:
		operandVal := c.generateExpression(e.Operand)
		switch e.Operator {
		case "+":
			return operandVal
		case "-":
			return c.builder.CreateNeg(operandVal, "negtmp")

		default:
			log.Fatalf("Error: Unknown unary operator '%s' in codegen.", e.Operator)
		}

	case *ast.CallExpression:
		log.Fatal("Error: Function calls not implemented in codegen.")

	default:
		log.Fatalf("Unhandled expression type in codegen: %T\n", e)
	}

	return llvm.ConstInt(c.context.Int32Type(), 0, false) // Unreachable fallback.
}

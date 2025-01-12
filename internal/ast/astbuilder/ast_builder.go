package astbuilder

import (
	"log"

	"github.com/richardzmija/fusion-compiler/internal/ast"
	"github.com/richardzmija/fusion-compiler/internal/parser"
)

// The first placeholder specifies the node of the parse tree that is being processed.
//
// The second placeholder specifies the type that should have been returned when building
// the AST.
const wrongTypeMessageTemplate string = "Error processing %s node of the parse tree: " +
	"Expected result of type %s but got %T.\n"

// assertType checks whether the result is of type T. If this is the case it
// returns the asserted value with type T. Otherwise it terminates the program
// with an informational error message.
func assertType[T any](result interface{}, context, expectedType string) T {
	v, ok := result.(T)
	if !ok {
		log.Fatalf(wrongTypeMessageTemplate, context, expectedType, result)
	}
	return v
}

type ASTBuilder struct {
	*parser.BaseCVisitor
}

func NewASTBuilder() *ASTBuilder {
	return &ASTBuilder{
		BaseCVisitor: &parser.BaseCVisitor{},
	}
}

func (b *ASTBuilder) VisitProgram(ctx parser.ProgramContext) interface{} {
	program := &ast.Program{}

	for _, functionDefinitionCtx := range ctx.AllFunctionDefinition() {
		result := functionDefinitionCtx.Accept(b)
		function := assertType[*ast.FunctionDefinition](result, "Program", "*astFunctionDefinition")
		program.Functions = append(program.Functions, function)
	}

	return program
}

func (b *ASTBuilder) VisitFunctionDefinition(ctx parser.FunctionDefinitionContext) interface{} {
	function := &ast.FunctionDefinition{
		Name:       ctx.ID().GetText(),
		ReturnType: "int", // For the current C subset all functions have return type 'int'.
	}

	if parameterListCtx := ctx.ParameterList(); parameterListCtx != nil {
		result := parameterListCtx.Accept(b)
		function.Parameters = assertType[[]*ast.Parameter](result, "FunctionDefinition", "[]*Parameter")
	}

	compoundStatementCtx := ctx.CompoundStatement()
	result := compoundStatementCtx.Accept(b)
	function.Body = assertType[*ast.BlockStatement](result, "FunctionDefinition", "*BlockStatement")

	return function
}

func (b *ASTBuilder) VisitParameterList(ctx parser.ParameterListContext) interface{} {
	var parameters []*ast.Parameter

	for _, parameterDeclCtx := range ctx.AllParameterDeclaration() {
		result := parameterDeclCtx.Accept(b)
		parameter := assertType[*ast.Parameter](result, "ParameterList", "*Parameter")
		parameters = append(parameters, parameter)
	}

	return parameters
}

func (b *ASTBuilder) VisitParameterDeclaration(ctx parser.ParameterDeclarationContext) interface{} {
	return &ast.Parameter{
		Name:     ctx.ID().GetText(),
		DataType: "int", // For the current C subset all parameters have type 'int'.
	}
}

func (b *ASTBuilder) VisitCompoundStatement(ctx parser.CompoundStatementContext) interface{} {
	blockStatement := &ast.BlockStatement{}

	if declarationListCtx := ctx.DeclarationList(); declarationListCtx != nil {
		result := declarationListCtx.Accept(b)
		blockStatement.Declarations = assertType[[]*ast.Declaration](result, "CompoundStatement", "[]*Declaration")
	}

	if statementListCtx := ctx.StatementList(); statementListCtx != nil {
		result := statementListCtx.Accept(b)
		blockStatement.Statements = assertType[[]ast.Statement](result, "CompoundStatement", "[]Statement")
	}

	return blockStatement
}

func (b *ASTBuilder) VisitStatementList(ctx parser.StatementListContext) interface{} {
	var statements []ast.Statement

	for _, statementCtx := range ctx.AllStatement() {
		result := statementCtx.Accept(b)
		statement := assertType[ast.Statement](result, "StatementList", "Statement")
		statements = append(statements, statement)
	}

	return statements
}

func (b *ASTBuilder) VisitStatement(ctx parser.StatementContext) interface{} {
	const expectedType = "*BlockStatement|*ExpressionStatement|*IfStatement|" +
		"*WhileStatement|*ReturnStatement|*PrintfStatement"

	var result interface{}

	switch {
	case ctx.CompoundStatement() != nil:
		result = ctx.CompoundStatement().Accept(b)
	case ctx.ExpressionStatement() != nil:
		result = ctx.ExpressionStatement().Accept(b)
	case ctx.SelectionStatement() != nil:
		result = ctx.SelectionStatement().Accept(b)
	case ctx.IterationStatement() != nil:
		result = ctx.IterationStatement().Accept(b)
	case ctx.JumpStatement() != nil:
		result = ctx.JumpStatement().Accept(b)
	case ctx.PrintfStatement() != nil:
		result = ctx.PrintfStatement().Accept(b)
	default:
		log.Fatalf("Error processing Statement node of the parse tree: Missing valid child node context.")
	}

	return assertType[ast.Statement](result, "Statement", expectedType)
}

func (b *ASTBuilder) VisitExpressionStatement(ctx parser.ExpressionStatementContext) interface{} {
	expressionStatement := &ast.ExpressionStatement{}

	if expressionCtx := ctx.Expression(); expressionCtx != nil {
		result := expressionCtx.Accept(b)
		expressionStatement.ContainedExpression = assertType[ast.Expression](result, "ExpressionStatement", "Expression")
	}

	return expressionStatement
}

func (b *ASTBuilder) VisitSelectionStatement(ctx parser.SelectionStatementContext) interface{} {
	conditionResult := ctx.Expression().Accept(b)
	condition := assertType[ast.Expression](conditionResult, "SelectionStatement", "Expression")

	thenStatementResult := ctx.Statement(0).Accept(b)
	thenStatement := assertType[ast.Statement](thenStatementResult, "SelectionStatement", "Statement")

	var elseStatement ast.Statement
	if ctx.ELSE() != nil {
		elseStatementResult := ctx.Statement(1).Accept(b)
		elseStatement = assertType[ast.Statement](elseStatementResult, "SelectionStatement", "Statement")
	}

	return &ast.IfStatement{
		Condition: condition,
		Then:      thenStatement,
		Else:      elseStatement,
	}
}

func (b *ASTBuilder) VisitIterationStatement(ctx parser.IterationStatementContext) interface{} {
	conditionResult := ctx.Expression().Accept(b)
	condition := assertType[ast.Expression](conditionResult, "IterationStatement", "Expression")

	bodyResult := ctx.Statement().Accept(b)
	body := assertType[ast.Statement](bodyResult, "IterationStatement", "Statement")

	return &ast.WhileStatement{
		Condition: condition,
		Body:      body,
	}
}

func (b *ASTBuilder) VisitJumpStatement(ctx parser.JumpStatementContext) interface{} {
	var returnValue ast.Expression

	if expressionCtx := ctx.Expression(); expressionCtx != nil {
		returnValueResult := expressionCtx.Accept(b)
		returnValue = assertType[ast.Expression](returnValueResult, "JumpStatement", "Expression")
	}

	return &ast.ReturnStatement{
		ReturnValue: returnValue,
	}
}

func (b *ASTBuilder) VisitPrintfStatement(ctx parser.PrintfStatementContext) interface{} {
	printfStatement := &ast.PrintfStatement{
		Format: removeQuotes(ctx.STR().GetText()),
	}

	for _, expressionCtx := range ctx.AllExpression() {
		result := expressionCtx.Accept(b)
		argument := assertType[ast.Expression](result, "PrintfStatement", "Expression")
		printfStatement.Arguments = append(printfStatement.Arguments, argument)
	}

	return printfStatement
}

// removeQuotes removes quotes from a string literal
// and returns it contents.
func removeQuotes(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}

	return s
}

package ast

import (
	"log"

	"github.com/antlr4-go/antlr/v4"
	"github.com/richardzmija/fusion-compiler/internal/parser"
)

// The first placeholder specifies the node of the parse tree that is being processed.
//
// The second placeholder specifies the type that should have been returned when building
// the AST.
const wrongTypeMessageTemplate string = "Error processing %s node of the parse tree: " +
	"Expected result of type %s but got %T.\n"

// The placeholder specifies the node of the parse tree that is being processed.
const unknownExpressionMessageTemplate string = "Error processing %s node of the parse tree: " +
	"Unknown expression.\n"

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
		BaseCVisitor: &parser.BaseCVisitor{
			BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{},
		},
	}
}

func (b *ASTBuilder) VisitProgram(ctx *parser.ProgramContext) interface{} {
	program := &Program{}

	for _, functionDefinitionCtx := range ctx.AllFunctionDefinition() {
		result := functionDefinitionCtx.Accept(b)
		function := assertType[*FunctionDefinition](result, "Program", "*astFunctionDefinition")
		program.Functions = append(program.Functions, function)
	}

	return program
}

func (b *ASTBuilder) VisitFunctionDefinition(ctx *parser.FunctionDefinitionContext) interface{} {
	function := &FunctionDefinition{
		Name:       ctx.ID().GetText(),
		ReturnType: IntType, // For the current C subset all functions have return type 'int'.
	}

	if parameterListCtx := ctx.ParameterList(); parameterListCtx != nil {
		result := parameterListCtx.Accept(b)
		function.Parameters = assertType[[]*Parameter](result, "FunctionDefinition", "[]*Parameter")
	}

	compoundStatementCtx := ctx.CompoundStatement()
	result := compoundStatementCtx.Accept(b)
	function.Body = assertType[*BlockStatement](result, "FunctionDefinition", "*BlockStatement")

	return function
}

func (b *ASTBuilder) VisitParameterList(ctx *parser.ParameterListContext) interface{} {
	var parameters []*Parameter

	for _, parameterDeclCtx := range ctx.AllParameterDeclaration() {
		result := parameterDeclCtx.Accept(b)
		parameter := assertType[*Parameter](result, "ParameterList", "*Parameter")
		parameters = append(parameters, parameter)
	}

	return parameters
}

func (b *ASTBuilder) VisitParameterDeclaration(ctx *parser.ParameterDeclarationContext) interface{} {
	return &Parameter{
		Name:     ctx.ID().GetText(),
		BaseType: IntType, // For the current C subset all parameters have type 'int'.
	}
}

func (b *ASTBuilder) VisitCompoundStatement(ctx *parser.CompoundStatementContext) interface{} {
	blockStatement := &BlockStatement{}

	if declarationListCtx := ctx.DeclarationList(); declarationListCtx != nil {
		result := declarationListCtx.Accept(b)
		blockStatement.Declarations = assertType[[]*Declaration](result, "CompoundStatement", "[]*Declaration")
	}

	if statementListCtx := ctx.StatementList(); statementListCtx != nil {
		result := statementListCtx.Accept(b)
		blockStatement.Statements = assertType[[]Statement](result, "CompoundStatement", "[]Statement")
	}

	return blockStatement
}

func (b *ASTBuilder) VisitStatementList(ctx *parser.StatementListContext) interface{} {
	var statements []Statement

	for _, statementCtx := range ctx.AllStatement() {
		result := statementCtx.Accept(b)
		statement := assertType[Statement](result, "StatementList", "Statement")
		statements = append(statements, statement)
	}

	return statements
}

func (b *ASTBuilder) VisitStatement(ctx *parser.StatementContext) interface{} {
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

	return assertType[Statement](result, "Statement", expectedType)
}

func (b *ASTBuilder) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) interface{} {
	expressionStatement := &ExpressionStatement{}

	if expressionCtx := ctx.Expression(); expressionCtx != nil {
		result := expressionCtx.Accept(b)
		expressionStatement.ContainedExpression = assertType[Expression](result, "ExpressionStatement", "Expression")
	}

	return expressionStatement
}

func (b *ASTBuilder) VisitSelectionStatement(ctx *parser.SelectionStatementContext) interface{} {
	conditionResult := ctx.Expression().Accept(b)
	condition := assertType[Expression](conditionResult, "SelectionStatement", "Expression")

	thenStatementResult := ctx.Statement(0).Accept(b)
	thenStatement := assertType[Statement](thenStatementResult, "SelectionStatement", "Statement")

	var elseStatement Statement
	if ctx.ELSE() != nil {
		elseStatementResult := ctx.Statement(1).Accept(b)
		elseStatement = assertType[Statement](elseStatementResult, "SelectionStatement", "Statement")
	}

	return &IfStatement{
		Condition: condition,
		Then:      thenStatement,
		Else:      elseStatement,
	}
}

func (b *ASTBuilder) VisitIterationStatement(ctx *parser.IterationStatementContext) interface{} {
	conditionResult := ctx.Expression().Accept(b)
	condition := assertType[Expression](conditionResult, "IterationStatement", "Expression")

	bodyResult := ctx.Statement().Accept(b)
	body := assertType[Statement](bodyResult, "IterationStatement", "Statement")

	return &WhileStatement{
		Condition: condition,
		Body:      body,
	}
}

func (b *ASTBuilder) VisitJumpStatement(ctx *parser.JumpStatementContext) interface{} {
	var returnValue Expression

	if expressionCtx := ctx.Expression(); expressionCtx != nil {
		returnValueResult := expressionCtx.Accept(b)
		returnValue = assertType[Expression](returnValueResult, "JumpStatement", "Expression")
	}

	return &ReturnStatement{
		ReturnValue: returnValue,
	}
}

func (b *ASTBuilder) VisitPrintfStatement(ctx *parser.PrintfStatementContext) interface{} {
	printfStatement := &PrintfStatement{
		Format: removeQuotes(ctx.STR().GetText()),
	}

	for _, expressionCtx := range ctx.AllExpression() {
		result := expressionCtx.Accept(b)
		argument := assertType[Expression](result, "PrintfStatement", "Expression")
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

type declaratorExpressionPair struct {
	name        string
	initializer Expression
}

func (b *ASTBuilder) VisitDeclarationList(ctx *parser.DeclarationListContext) interface{} {
	var declarations []*Declaration

	for _, declarationCtx := range ctx.AllDeclaration() {
		result := declarationCtx.Accept(b)
		declaration := assertType[*Declaration](result, "DeclarationList", "*Declaration")
		declarations = append(declarations, declaration)
	}

	return declarations
}

func (b *ASTBuilder) VisitDeclaration(ctx *parser.DeclarationContext) interface{} {
	declaration := &Declaration{
		Type: IntType, // For the current C subset all variable declarations use type 'int'.
	}

	result := ctx.InitDeclaratorList().Accept(b)
	pairs := assertType[[]declaratorExpressionPair](result, "Declaration", "[]declaratorExpressionPair")

	for _, pair := range pairs {
		declaration.Names = append(declaration.Names, pair.name)
		declaration.Initializers = append(declaration.Initializers, pair.initializer)
	}

	return declaration
}

func (b *ASTBuilder) VisitInitDeclaratorList(ctx *parser.InitDeclaratorListContext) interface{} {
	var declarators []declaratorExpressionPair

	for _, declaratorCtx := range ctx.AllInitDeclarator() {
		result := declaratorCtx.Accept(b)
		pair := assertType[declaratorExpressionPair](result, "InitDeclaratorList", "declaratorExpressionPair")
		declarators = append(declarators, pair)
	}

	return declarators
}

func (b *ASTBuilder) VisitInitDeclarator(ctx *parser.InitDeclaratorContext) interface{} {
	declarator := declaratorExpressionPair{
		name: ctx.ID().GetText(),
	}

	if ctx.ASSIGN() != nil {
		result := ctx.Expression().Accept(b)
		declarator.initializer = assertType[Expression](result, "InitDeclarator", "Expression")
	}

	return declarator
}

func (b *ASTBuilder) VisitConstant(ctx *parser.ConstantContext) interface{} {
	literal := &Literal{}

	if numCtx := ctx.NUM(); numCtx != nil {
		literal.Type = IntLiteral
		literal.Value = numCtx.GetText()
	} else if strCtx := ctx.STR(); strCtx != nil {
		literal.Type = StringLiteral
		literal.Value = removeQuotes(strCtx.GetText())
	} else {
		log.Fatalf(unknownExpressionMessageTemplate, "Constant")
	}

	return literal
}

func (b *ASTBuilder) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	if constantCtx := ctx.Constant(); constantCtx != nil {
		result := constantCtx.Accept(b)
		return assertType[Expression](result, "PrimaryExpression", "*Literal")
	}

	if idCtx := ctx.ID(); idCtx != nil {
		return &VariableExpression{
			Name: idCtx.GetText(),
		}
	}

	if exprCtx := ctx.Expression(); exprCtx != nil {
		result := exprCtx.Accept(b)
		return assertType[Expression](result, "PrimaryExpression", "Expression")
	}

	log.Fatalf(unknownExpressionMessageTemplate, "PrimaryExpression")
	return nil
}

func (b *ASTBuilder) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{} {
	if primaryExprCtx := ctx.PrimaryExpression(); primaryExprCtx != nil {
		result := primaryExprCtx.Accept(b)
		return assertType[Expression](result, "PostfixExpression", "Expression")
	}

	if postfixExprCtx := ctx.PostfixExpression(); postfixExprCtx != nil {
		resultPostfix := postfixExprCtx.Accept(b)
		postfix := assertType[Expression](resultPostfix, "PostfixExpression", "Expression")

		var argumentExpressions []Expression
		if argumentExprListCtx := ctx.ArgumentExpressionList(); argumentExprListCtx != nil {
			resultArgumentExpr := argumentExprListCtx.Accept(b)
			argumentExpressions = assertType[[]Expression](resultArgumentExpr, "PostfixExpression", "[]Expression")
		}

		return &CallExpression{
			Callee:    postfix,
			Arguments: argumentExpressions,
		}
	}

	log.Fatalf(unknownExpressionMessageTemplate, "PostfixExpression")
	return nil
}

func (b *ASTBuilder) VisitArgumentExpressionList(ctx *parser.ArgumentExpressionListContext) interface{} {
	var argumentExpressions []Expression

	for _, exprCtx := range ctx.AllAssignmentExpression() {
		result := exprCtx.Accept(b)
		argumentExpression := assertType[Expression](result, "ArgumentExpressionList", "Expression")
		argumentExpressions = append(argumentExpressions, argumentExpression)
	}

	return argumentExpressions
}

func (b *ASTBuilder) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	if postfixExprCtx := ctx.PostfixExpression(); postfixExprCtx != nil {
		result := postfixExprCtx.Accept(b)
		return assertType[Expression](result, "UnaryExpression", "Expression")
	}

	if ctx.PLUS() != nil {
		result := ctx.UnaryExpression().Accept(b)
		operand := assertType[Expression](result, "UnaryExpression", "Expression")
		return &UnaryExpression{
			Operator: "+",
			Operand:  operand,
		}
	}

	if ctx.MINUS() != nil {
		result := ctx.UnaryExpression().Accept(b)
		operand := assertType[Expression](result, "UnaryExpression", "Expression")
		return &UnaryExpression{
			Operator: "-",
			Operand:  operand,
		}
	}

	log.Fatalf(unknownExpressionMessageTemplate, "UnaryExpression")
	return nil
}

func leftAssociativeReduction[T antlr.ParseTree](operands []T, ctx antlr.BaseParserRuleContext,
	nodeName string, visitor *ASTBuilder) Expression {

	if len(operands) == 1 {
		return assertType[Expression](operands[0].Accept(visitor), nodeName, "Expression")
	}

	// Begin with the leftmost operand and use it to construct the first BinaryExpression instance.
	left := assertType[Expression](operands[0].Accept(visitor), nodeName, "Expression")

	for opIndex := 1; opIndex < ctx.GetChildCount(); opIndex += 2 {
		token, ok := ctx.GetChild(opIndex).(antlr.TerminalNode)
		if !ok {
			log.Fatalf("Error processing node %s of the parse tree: "+
				"Expected terminal node for an operator.", nodeName)
		}
		operator := token.GetText()

		// Calculate the index for the corresponding right-hand operand based on the index
		// of the operator.
		rightOperandIndex := (opIndex + 1) / 2
		if rightOperandIndex >= len(operands) {
			log.Fatalf("Error processing node %s of the parse tree: "+
				"Index out of bounds!", nodeName)
		}

		right := assertType[Expression](operands[rightOperandIndex].Accept(visitor),
			nodeName, "Expression")

		left = &BinaryExpression{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}

	return left
}

func (b *ASTBuilder) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	unaryExpressions := ctx.AllUnaryExpression()
	return leftAssociativeReduction(unaryExpressions, ctx.BaseParserRuleContext, "MultiplicativeExpression", b)
}

func (b *ASTBuilder) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	multiplicativeExpressions := ctx.AllMultiplicativeExpression()
	return leftAssociativeReduction(multiplicativeExpressions, ctx.BaseParserRuleContext, "AdditiveExpression", b)
}

func (b *ASTBuilder) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) interface{} {
	additiveExpressions := ctx.AllAdditiveExpression()
	return leftAssociativeReduction(additiveExpressions, ctx.BaseParserRuleContext, "RelationalExpression", b)
}

func (b *ASTBuilder) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} {
	relationalExpressions := ctx.AllRelationalExpression()
	return leftAssociativeReduction(relationalExpressions, ctx.BaseParserRuleContext, "EqualityExpression", b)
}

func (b *ASTBuilder) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) interface{} {
	equalityExpressions := ctx.AllEqualityExpression()
	return leftAssociativeReduction(equalityExpressions, ctx.BaseParserRuleContext, "LogicalAndExpression", b)
}

func (b *ASTBuilder) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) interface{} {
	logicalAndExpressions := ctx.AllLogicalAndExpression()
	return leftAssociativeReduction(logicalAndExpressions, ctx.BaseParserRuleContext, "LogicalOrExpression", b)
}

func (b *ASTBuilder) VisitConditionalExpression(ctx *parser.ConditionalExpressionContext) interface{} {
	if logicalOrExpressionCtx := ctx.LogicalOrExpression(); logicalOrExpressionCtx != nil {
		result := logicalOrExpressionCtx.Accept(b)
		return assertType[Expression](result, "ConditionalExpression", "Expression")
	}

	log.Fatalf(unknownExpressionMessageTemplate, "ConditionalExpression")
	return nil
}

func (b *ASTBuilder) VisitAssignmentExpression(ctx *parser.AssignmentExpressionContext) interface{} {
	if conditionalExprCtx := ctx.ConditionalExpression(); conditionalExprCtx != nil {
		result := conditionalExprCtx.Accept(b)
		return assertType[Expression](result, "AssignmentExpression", "Expression")
	}

	if unaryExprCtx := ctx.UnaryExpression(); unaryExprCtx != nil {
		leftResult := unaryExprCtx.Accept(b)
		left := assertType[Expression](leftResult, "AssignmentExpression", "Expression")

		rightResult := ctx.AssignmentExpression().Accept(b)
		right := assertType[Expression](rightResult, "AssignmentExpression", "Expression")

		return &BinaryExpression{
			Left:     left,
			Operator: ctx.ASSIGN().GetText(),
			Right:    right,
		}
	}

	log.Fatalf(unknownExpressionMessageTemplate, "AssignmentExpression")
	return nil
}

func (b *ASTBuilder) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	if assignmentExprCtx := ctx.AssignmentExpression(); assignmentExprCtx != nil {
		return assertType[Expression](assignmentExprCtx.Accept(b), "Expression", "Expression")
	}

	log.Fatalf(unknownExpressionMessageTemplate, "Expression")
	return nil
}

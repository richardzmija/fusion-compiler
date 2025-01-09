// Code generated from ../grammar/C.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // C

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CParser.
type CVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by CParser#functionDefinition.
	VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{}

	// Visit a parse tree produced by CParser#parameterList.
	VisitParameterList(ctx *ParameterListContext) interface{}

	// Visit a parse tree produced by CParser#parameterDeclaration.
	VisitParameterDeclaration(ctx *ParameterDeclarationContext) interface{}

	// Visit a parse tree produced by CParser#declarationList.
	VisitDeclarationList(ctx *DeclarationListContext) interface{}

	// Visit a parse tree produced by CParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by CParser#initDeclaratorList.
	VisitInitDeclaratorList(ctx *InitDeclaratorListContext) interface{}

	// Visit a parse tree produced by CParser#initDeclarator.
	VisitInitDeclarator(ctx *InitDeclaratorContext) interface{}

	// Visit a parse tree produced by CParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by CParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by CParser#compoundStatement.
	VisitCompoundStatement(ctx *CompoundStatementContext) interface{}

	// Visit a parse tree produced by CParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by CParser#selectionStatement.
	VisitSelectionStatement(ctx *SelectionStatementContext) interface{}

	// Visit a parse tree produced by CParser#iterationStatement.
	VisitIterationStatement(ctx *IterationStatementContext) interface{}

	// Visit a parse tree produced by CParser#jumpStatement.
	VisitJumpStatement(ctx *JumpStatementContext) interface{}

	// Visit a parse tree produced by CParser#printfStatement.
	VisitPrintfStatement(ctx *PrintfStatementContext) interface{}

	// Visit a parse tree produced by CParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by CParser#assignmentExpression.
	VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{}

	// Visit a parse tree produced by CParser#conditionalExpression.
	VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{}

	// Visit a parse tree produced by CParser#logicalOrExpression.
	VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{}

	// Visit a parse tree produced by CParser#logicalAndExpression.
	VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{}

	// Visit a parse tree produced by CParser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by CParser#relationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by CParser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by CParser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by CParser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by CParser#postfixExpression.
	VisitPostfixExpression(ctx *PostfixExpressionContext) interface{}

	// Visit a parse tree produced by CParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by CParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by CParser#argumentExpressionList.
	VisitArgumentExpressionList(ctx *ArgumentExpressionListContext) interface{}
}

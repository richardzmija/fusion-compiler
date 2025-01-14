// Code generated from ../grammar/C.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // C

import "github.com/antlr4-go/antlr/v4"

type BaseCVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitParameterList(ctx *ParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitParameterDeclaration(ctx *ParameterDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitDeclarationList(ctx *DeclarationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitInitDeclaratorList(ctx *InitDeclaratorListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitInitDeclarator(ctx *InitDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitStatementList(ctx *StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitCompoundStatement(ctx *CompoundStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitSelectionStatement(ctx *SelectionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitIterationStatement(ctx *IterationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitJumpStatement(ctx *JumpStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitPrintfStatement(ctx *PrintfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitRelationalExpression(ctx *RelationalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitPostfixExpression(ctx *PostfixExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitArgumentExpressionList(ctx *ArgumentExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

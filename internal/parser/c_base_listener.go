// Code generated from C.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // C
import "github.com/antlr4-go/antlr/v4"

// BaseCListener is a complete listener for a parse tree produced by CParser.
type BaseCListener struct{}

var _ CListener = &BaseCListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseCListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseCListener) ExitProgram(ctx *ProgramContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *BaseCListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *BaseCListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseCListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseCListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterParameterDeclaration is called when production parameterDeclaration is entered.
func (s *BaseCListener) EnterParameterDeclaration(ctx *ParameterDeclarationContext) {}

// ExitParameterDeclaration is called when production parameterDeclaration is exited.
func (s *BaseCListener) ExitParameterDeclaration(ctx *ParameterDeclarationContext) {}

// EnterDeclarationList is called when production declarationList is entered.
func (s *BaseCListener) EnterDeclarationList(ctx *DeclarationListContext) {}

// ExitDeclarationList is called when production declarationList is exited.
func (s *BaseCListener) ExitDeclarationList(ctx *DeclarationListContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseCListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseCListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterInitDeclaratorList is called when production initDeclaratorList is entered.
func (s *BaseCListener) EnterInitDeclaratorList(ctx *InitDeclaratorListContext) {}

// ExitInitDeclaratorList is called when production initDeclaratorList is exited.
func (s *BaseCListener) ExitInitDeclaratorList(ctx *InitDeclaratorListContext) {}

// EnterInitDeclarator is called when production initDeclarator is entered.
func (s *BaseCListener) EnterInitDeclarator(ctx *InitDeclaratorContext) {}

// ExitInitDeclarator is called when production initDeclarator is exited.
func (s *BaseCListener) ExitInitDeclarator(ctx *InitDeclaratorContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseCListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseCListener) ExitStatementList(ctx *StatementListContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseCListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseCListener) ExitStatement(ctx *StatementContext) {}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *BaseCListener) EnterCompoundStatement(ctx *CompoundStatementContext) {}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *BaseCListener) ExitCompoundStatement(ctx *CompoundStatementContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseCListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseCListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterSelectionStatement is called when production selectionStatement is entered.
func (s *BaseCListener) EnterSelectionStatement(ctx *SelectionStatementContext) {}

// ExitSelectionStatement is called when production selectionStatement is exited.
func (s *BaseCListener) ExitSelectionStatement(ctx *SelectionStatementContext) {}

// EnterIterationStatement is called when production iterationStatement is entered.
func (s *BaseCListener) EnterIterationStatement(ctx *IterationStatementContext) {}

// ExitIterationStatement is called when production iterationStatement is exited.
func (s *BaseCListener) ExitIterationStatement(ctx *IterationStatementContext) {}

// EnterJumpStatement is called when production jumpStatement is entered.
func (s *BaseCListener) EnterJumpStatement(ctx *JumpStatementContext) {}

// ExitJumpStatement is called when production jumpStatement is exited.
func (s *BaseCListener) ExitJumpStatement(ctx *JumpStatementContext) {}

// EnterPrintfStatement is called when production printfStatement is entered.
func (s *BaseCListener) EnterPrintfStatement(ctx *PrintfStatementContext) {}

// ExitPrintfStatement is called when production printfStatement is exited.
func (s *BaseCListener) ExitPrintfStatement(ctx *PrintfStatementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *BaseCListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *BaseCListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *BaseCListener) EnterConditionalExpression(ctx *ConditionalExpressionContext) {}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *BaseCListener) ExitConditionalExpression(ctx *ConditionalExpressionContext) {}

// EnterLogicalOrExpression is called when production logicalOrExpression is entered.
func (s *BaseCListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production logicalOrExpression is exited.
func (s *BaseCListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// EnterLogicalAndExpression is called when production logicalAndExpression is entered.
func (s *BaseCListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production logicalAndExpression is exited.
func (s *BaseCListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseCListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseCListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *BaseCListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *BaseCListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseCListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseCListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseCListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseCListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseCListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseCListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *BaseCListener) EnterPostfixExpression(ctx *PostfixExpressionContext) {}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *BaseCListener) ExitPostfixExpression(ctx *PostfixExpressionContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseCListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseCListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseCListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseCListener) ExitConstant(ctx *ConstantContext) {}

// EnterArgumentExpressionList is called when production argumentExpressionList is entered.
func (s *BaseCListener) EnterArgumentExpressionList(ctx *ArgumentExpressionListContext) {}

// ExitArgumentExpressionList is called when production argumentExpressionList is exited.
func (s *BaseCListener) ExitArgumentExpressionList(ctx *ArgumentExpressionListContext) {}

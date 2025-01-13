package ast

// Node represents a node in an Abstract Syntax Tree.
type Node interface {
	// PositionInSource returns the position of the node in source code.
	PositionInSource() int
}

// Program is the root node of an Abstract Syntax Tree.
type Program struct {
	Functions []*FunctionDefinition
}

func (prog *Program) PositionInSource() int {
	return 0 // Stub
}

// FunctionDefinition represents a function definition.
type FunctionDefinition struct {
	Name       string
	ReturnType string
	Parameters []*Parameter
	Body       *BlockStatement
}

func (funDef *FunctionDefinition) PositionInSource() int {
	return 0 // Stub
}

// Parameter represents a function parameter.
type Parameter struct {
	Name     string
	DataType string
}

func (param *Parameter) PositionInSource() int {
	return 0 // Stub
}

// Declaration represents a declaration of a list of variables.
type Declaration struct {
	Type         string
	Names        []string
	Initializers []Expression // nil if no initialization expression is provided
}

func (d *Declaration) PositionInSource() int {
	return 0 // Stub
}

// Statement represents a program statement.
type Statement interface {
	Node
	// Other interface methods for Statements.
}

// BlockStatement represents a block statement (compound statement).
type BlockStatement struct {
	Declarations []*Declaration
	Statements   []Statement
}

func (bs *BlockStatement) PositionInSource() int {
	return 0 // Stub
}

// ExpressionStatement represents an expression that is a statement.
type ExpressionStatement struct {
	ContainedExpression Expression
}

func (es *ExpressionStatement) PositionInSource() int {
	return 0 // Stub
}

// IfStatement represents an if statement with an optional else clause.
type IfStatement struct {
	Condition Expression
	Then      Statement
	Else      Statement // nil if the else clause is omitted
}

func (ifs *IfStatement) PositionInSource() int {
	return 0 // Stub
}

// WhileStatement represents a while loop statement.
type WhileStatement struct {
	Condition Expression
	Body      Statement
}

func (ws *WhileStatement) PositionInSource() int {
	return 0 // Stub
}

// ReturnStatement represents a return jump statement.
type ReturnStatement struct {
	// ReturnValue is an expression whose result is the return
	// value of the function. If the return type `void` were permitted
	// then it should be set to nil.
	ReturnValue Expression
}

func (rs *ReturnStatement) PositionInSource() int {
	return 0 // Stub
}

// PrintfStatement represents a special, non-standard statement for writing
// to the standard output.
type PrintfStatement struct {
	Format    string
	Arguments []Expression
}

func (ps *PrintfStatement) PositionInSource() int {
	return 0 // Stub
}

// Expression represents an expression.
type Expression interface {
	Node
	// Other interface methods for Expressions.
}

// BinaryExpression represents an expression with two operands in the in-fix notation.
type BinaryExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (be *BinaryExpression) PositionInSource() int {
	return 0 // Stub
}

// UnaryExpression represents an expression with one operand.
type UnaryExpression struct {
	Operator string
	Operand  Expression
}

func (ue *UnaryExpression) PositionInSource() int {
	return 0 // Stub
}

// Literal represents a constant expression.
type Literal struct {
	Value interface{}
}

func (l *Literal) PositionInSource() int {
	return 0 // Stub
}

// VariableExpression represents a variable that is treated as
// an expression.
type VariableExpression struct {
	Name string
}

func (ve *VariableExpression) PositionInSource() int {
	return 0 // Stub
}

// CallExpression represents an expression using the function call operator.
type CallExpression struct {
	// Callee is considered an expression because the name of the function
	// being called is treated as a variable holding a pointer to the function.
	Callee    Expression
	Arguments []Expression
}

func (ce *CallExpression) PositionInSource() int {
	return 0 // Stub
}

package semantic

import (
    "fmt"

    "github.com/richardzmija/fusion-compiler/internal/ast"
)

// Analyzer holds state while traversing the AST.
type Analyzer struct {
    currentScope *Scope
    errors       []error
}

// NewAnalyzer creates a new Analyzer with a global scope.
func NewAnalyzer() *Analyzer {
    return &Analyzer{
        currentScope: NewScope(nil), // global scope
        errors:       nil,
    }
}

// pushScope creates a new child scope and makes it the current scope.
func (a *Analyzer) pushScope() {
    a.currentScope = NewScope(a.currentScope)
}

// popScope returns to the parent scope.
func (a *Analyzer) popScope() {
    a.currentScope = a.currentScope.parent
}

// addError records a semantic error.
func (a *Analyzer) addError(pos int, msg string) {
    a.errors = append(a.errors, fmt.Errorf("pos %d: %s", pos, msg))
}

// Analyze is the entry point for semantic analysis.
// It returns a slice of errors found, or an empty slice if none.
func (a *Analyzer) Analyze(prog *ast.Program) []error {
    // First pass: collect function definitions in the global scope
    for _, fun := range prog.Functions {
        if err := a.defineFunction(fun); err != nil {
            a.addError(fun.PositionInSource(), err.Error())
        }
    }

    // Second pass: analyze function bodies
    for _, fun := range prog.Functions {
        a.analyzeFunction(fun)
    }

    return a.errors
}

// defineFunction creates a Symbol for the function and adds it to the global scope.
func (a *Analyzer) defineFunction(fun *ast.FunctionDefinition) error {
    // We assume all functions return int in this subset of the C language.
    sym := &Symbol{
        Name:       fun.Name,
        Kind:       SymbolFunc,
        DataType:   fun.ReturnType,
        ParamTypes: make([]ast.DataType, len(fun.Parameters)),
    }
    for i, p := range fun.Parameters {
        sym.ParamTypes[i] = p.BaseType
    }
    return a.currentScope.Define(sym)
}

// analyzeFunction opens a new scope for parameters, then visits the body.
func (a *Analyzer) analyzeFunction(fun *ast.FunctionDefinition) {
    // Push a scope for function parameters
    a.pushScope()
    defer a.popScope()

    // Define parameters in the function scope
    for _, param := range fun.Parameters {
        sym := &Symbol{
            Name:     param.Name,
            Kind:     SymbolVar,
            DataType: param.BaseType,
        }
        if err := a.currentScope.Define(sym); err != nil {
            a.addError(param.PositionInSource(), err.Error())
        }
    }

    // Now analyze the body (which itself pushes/pops scopes for inner blocks)
    a.analyzeBlockStatement(fun.Body)
}

// analyzeBlockStatement pushes a new scope for the block if it’s not the function’s top-level body.
// For function bodies, we already pushed in `analyzeFunction`.
func (a *Analyzer) analyzeBlockStatement(bs *ast.BlockStatement) {
    if bs == nil {
        return
    }

    a.pushScope()
    defer a.popScope()

    // Declarations
    for _, decl := range bs.Declarations {
        a.analyzeDeclaration(decl)
    }

    // Statements
    for _, stmt := range bs.Statements {
        a.analyzeStatement(stmt)
    }
}

// analyzeDeclaration defines each declared variable in the current scope and checks its initializer (if any).
func (a *Analyzer) analyzeDeclaration(decl *ast.Declaration) {
    for i, name := range decl.Names {
        sym := &Symbol{
            Name:     name,
            Kind:     SymbolVar,
            DataType: decl.Type, // In this subset, it's always int
        }
        if err := a.currentScope.Define(sym); err != nil {
            a.addError(decl.PositionInSource(), err.Error())
        }

        // Check initializer if present
        initExpr := decl.Initializers[i]
        if initExpr != nil {
            exprType := a.checkExpression(initExpr)
            if exprType != ast.IntType {
                msg := "cannot initialize int with expression of non-int type"
                a.addError(initExpr.PositionInSource(), msg)
            }
        }
    }
}

// analyzeStatement dispatches by statement type.
func (a *Analyzer) analyzeStatement(stmt ast.Statement) {
    switch s := stmt.(type) {
    case *ast.BlockStatement:
        a.analyzeBlockStatement(s)

    case *ast.ExpressionStatement:
        if s.ContainedExpression != nil {
            a.checkExpression(s.ContainedExpression)
        }

    case *ast.IfStatement:
        condType := a.checkExpression(s.Condition)
        if condType != ast.IntType {
            a.addError(s.Condition.PositionInSource(), "if condition must be int")
        }
        a.analyzeStatement(s.Then)
        if s.Else != nil {
            a.analyzeStatement(s.Else)
        }

    case *ast.WhileStatement:
        condType := a.checkExpression(s.Condition)
        if condType != ast.IntType {
            a.addError(s.Condition.PositionInSource(), "while condition must be int")
        }
        a.analyzeStatement(s.Body)

    case *ast.ReturnStatement:
        if s.ReturnValue != nil {
            retType := a.checkExpression(s.ReturnValue)
            if retType != ast.IntType {
                a.addError(s.PositionInSource(), "return expression must be int")
            }
        }

    case *ast.PrintfStatement:
        // Format is a string by definition.
        // But we do check that all arguments are int.
        for _, arg := range s.Arguments {
            argType := a.checkExpression(arg)
            if argType != ast.IntType {
                a.addError(arg.PositionInSource(), "printf argument must be int (in this subset)")
            }
        }

    default:
        a.addError(stmt.PositionInSource(), "unknown statement type")
    }
}

// checkExpression returns the computed type of the expression (always int or error in this subset).
func (a *Analyzer) checkExpression(expr ast.Expression) ast.DataType {
    switch e := expr.(type) {
    case *ast.Literal:
        // int literal or string literal
        switch e.Type {
        case ast.IntLiteral:
            return ast.IntType
        case ast.StringLiteral:
            // In standard C, a string literal has type char*, but in this subset,
            // we only allow string literal in printf, not in general expressions.
            // Return something non-int to help catch usage errors.
            return -1
        }

    case *ast.VariableExpression:
        sym, err := a.currentScope.Resolve(e.Name)
        if err != nil {
            a.addError(e.PositionInSource(), err.Error())
            return -1
        }
        // Must be a variable or function name. In standard C, function name decays to a pointer, etc.
        // But here, we only have int variables or function. If it's a function symbol, that's also invalid
        // in an expression context unless used in a call. We'll flag that as an error.
        if sym.Kind == SymbolFunc {
            a.addError(e.PositionInSource(), fmt.Sprintf("'%s' is a function, not a variable", e.Name))
            return -1
        }
        return sym.DataType

    case *ast.BinaryExpression:
        leftType := a.checkExpression(e.Left)
        rightType := a.checkExpression(e.Right)
        // In real C, operators can produce different result types, but our subset is simpler, because
        // we only allow int. Check both sides are int.
        if leftType != ast.IntType || rightType != ast.IntType {
            a.addError(e.PositionInSource(), "binary operator operands must be int")
            return -1
        }
        return ast.IntType

    case *ast.UnaryExpression:
        operandType := a.checkExpression(e.Operand)
        if operandType != ast.IntType {
            a.addError(e.PositionInSource(), "unary operator operand must be int")
            return -1
        }
        return ast.IntType

    case *ast.CallExpression:
        // Check that Callee is a variable (function name)
        calleeType := a.checkExpression(e.Callee)
        if calleeType != -1 {
            // If e.Callee is a variable expression, check if it is a function
            var varExpr *ast.VariableExpression
            if ve, ok := e.Callee.(*ast.VariableExpression); ok {
                varExpr = ve
            } else {
                a.addError(e.PositionInSource(), "function call must have a function name (in this subset)")
                return -1
            }

            sym, err := a.currentScope.Resolve(varExpr.Name)
            if err != nil {
                a.addError(varExpr.PositionInSource(), err.Error())
                return -1
            }
            if sym.Kind != SymbolFunc {
                a.addError(varExpr.PositionInSource(), fmt.Sprintf("'%s' is not a function", sym.Name))
                return -1
            }

            // Check argument list length
            if len(e.Arguments) != len(sym.ParamTypes) {
                msg := fmt.Sprintf("function '%s' called with %d args, but defined with %d",
                    sym.Name, len(e.Arguments), len(sym.ParamTypes))
                a.addError(e.PositionInSource(), msg)
                return -1
            }

            // Check argument types
            for i, arg := range e.Arguments {
                argType := a.checkExpression(arg)
                if argType != sym.ParamTypes[i] {
                    msg := fmt.Sprintf("argument %d of '%s' must be int", i, sym.Name)
                    a.addError(arg.PositionInSource(), msg)
                }
            }
            // Return type is the function’s return type (always int in this subset).
            return sym.DataType
        }
        return -1

    default:
        a.addError(expr.PositionInSource(), "unknown expression type")
        return -1
    }

    return -1
}

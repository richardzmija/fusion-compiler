package ast

import (
    "fmt"
    "strings"
)

// StringifyProgram takes the top-level *Program node (root of the AST)
// and returns a string representation of the entire tree.
func StringifyProgram(prog *Program) string {
    return stringifyNode(prog, 0)
}

// stringifyNode is a helper that performs a recursive traversal of the AST.
// It dispatches based on the concrete node type and uses indentation
// to reflect the hierarchy.
func stringifyNode(node Node, indentLevel int) string {
    if node == nil {
        return indent(indentLevel) + "<nil node>\n"
    }

    switch n := node.(type) {

    case *Program:
        sb := &strings.Builder{}
        sb.WriteString(indent(indentLevel))
        sb.WriteString("Program:\n")
        for _, fun := range n.Functions {
            sb.WriteString(stringifyNode(fun, indentLevel+1))
        }
        return sb.String()

    case *FunctionDefinition:
        sb := &strings.Builder{}
        sb.WriteString(indent(indentLevel))
        sb.WriteString(fmt.Sprintf("FunctionDefinition: %s\n", n.Name))

        sb.WriteString(indent(indentLevel + 1))
        sb.WriteString(fmt.Sprintf("ReturnType: %v\n", n.ReturnType))

        sb.WriteString(indent(indentLevel + 1))
        sb.WriteString("Parameters:\n")
        for _, param := range n.Parameters {
            sb.WriteString(stringifyNode(param, indentLevel+2))
        }

        sb.WriteString(indent(indentLevel + 1))
        sb.WriteString("Body:\n")
        if n.Body != nil {
            sb.WriteString(stringifyNode(n.Body, indentLevel+2))
        } else {
            sb.WriteString(indent(indentLevel+2) + "<nil body>\n")
        }
        return sb.String()

    case *Parameter:
        return indent(indentLevel) +
            fmt.Sprintf("Parameter: %s (Type: %v)\n", n.Name, n.BaseType)

    case *BlockStatement:
        sb := &strings.Builder{}
        sb.WriteString(indent(indentLevel))
        sb.WriteString("BlockStatement:\n")

        if len(n.Declarations) > 0 {
            sb.WriteString(indent(indentLevel + 1))
            sb.WriteString("Declarations:\n")
            for _, decl := range n.Declarations {
                sb.WriteString(stringifyNode(decl, indentLevel+2))
            }
        }

        if len(n.Statements) > 0 {
            sb.WriteString(indent(indentLevel + 1))
            sb.WriteString("Statements:\n")
            for _, stmt := range n.Statements {
                sb.WriteString(stringifyNode(stmt, indentLevel+2))
            }
        }
        return sb.String()

    case *Declaration:
        sb := &strings.Builder{}
        sb.WriteString(indent(indentLevel))
        sb.WriteString(fmt.Sprintf("Declaration (Type: %v):\n", n.Type))
        for i, name := range n.Names {
            sb.WriteString(indent(indentLevel + 1))
            if n.Initializers[i] != nil {
                sb.WriteString(fmt.Sprintf("%s = \n", name))
                sb.WriteString(stringifyNode(n.Initializers[i], indentLevel+2))
            } else {
                sb.WriteString(fmt.Sprintf("%s\n", name))
            }
        }
        return sb.String()

    case *ExpressionStatement:
        sb := &strings.Builder{}
        sb.WriteString(indent(indentLevel))
        sb.WriteString("ExpressionStatement:\n")
        if n.ContainedExpression != nil {
            sb.WriteString(stringifyNode(n.ContainedExpression, indentLevel+1))
        } else {
            sb.WriteString(indent(indentLevel+1) + "<empty>\n")
        }
        return sb.String()

    case *IfStatement:
        sb := &strings.Builder{}
        sb.WriteString(indent(indentLevel))
        sb.WriteString("IfStatement:\n")

        sb.WriteString(indent(indentLevel + 1))
        sb.WriteString("Condition:\n")
        sb.WriteString(stringifyNode(n.Condition, indentLevel+2))

        sb.WriteString(indent(indentLevel + 1))
        sb.WriteString("Then:\n")
        sb.WriteString(stringifyNode(n.Then, indentLevel+2))

        if n.Else != nil {
            sb.WriteString(indent(indentLevel + 1))
            sb.WriteString("Else:\n")
            sb.WriteString(stringifyNode(n.Else, indentLevel+2))
        }
        return sb.String()

    case *WhileStatement:
        sb := &strings.Builder{}
        sb.WriteString(indent(indentLevel))
        sb.WriteString("WhileStatement:\n")

        sb.WriteString(indent(indentLevel + 1))
        sb.WriteString("Condition:\n")
        sb.WriteString(stringifyNode(n.Condition, indentLevel+2))

        sb.WriteString(indent(indentLevel + 1))
        sb.WriteString("Body:\n")
        sb.WriteString(stringifyNode(n.Body, indentLevel+2))
        return sb.String()

    case *ReturnStatement:
        sb := &strings.Builder{}
        sb.WriteString(indent(indentLevel))
        sb.WriteString("ReturnStatement:\n")
        if n.ReturnValue != nil {
            sb.WriteString(stringifyNode(n.ReturnValue, indentLevel+1))
        } else {
            sb.WriteString(indent(indentLevel+1) + "<no return value>\n")
        }
        return sb.String()

    case *PrintfStatement:
        sb := &strings.Builder{}
        sb.WriteString(indent(indentLevel))
        sb.WriteString("PrintfStatement:\n")
        sb.WriteString(indent(indentLevel + 1))
        sb.WriteString(fmt.Sprintf("Format: \"%s\"\n", n.Format))
        if len(n.Arguments) > 0 {
            sb.WriteString(indent(indentLevel + 1))
            sb.WriteString("Arguments:\n")
            for _, arg := range n.Arguments {
                sb.WriteString(stringifyNode(arg, indentLevel+2))
            }
        }
        return sb.String()

    case *BinaryExpression:
        sb := &strings.Builder{}
        sb.WriteString(indent(indentLevel))
        sb.WriteString(fmt.Sprintf("BinaryExpression: %s\n", n.Operator))

        sb.WriteString(indent(indentLevel + 1))
        sb.WriteString("Left:\n")
        sb.WriteString(stringifyNode(n.Left, indentLevel+2))

        sb.WriteString(indent(indentLevel + 1))
        sb.WriteString("Right:\n")
        sb.WriteString(stringifyNode(n.Right, indentLevel+2))
        return sb.String()

    case *UnaryExpression:
        sb := &strings.Builder{}
        sb.WriteString(indent(indentLevel))
        sb.WriteString(fmt.Sprintf("UnaryExpression: %s\n", n.Operator))
        sb.WriteString(indent(indentLevel + 1))
        sb.WriteString("Operand:\n")
        sb.WriteString(stringifyNode(n.Operand, indentLevel+2))
        return sb.String()

    case *Literal:
        sb := &strings.Builder{}
        sb.WriteString(indent(indentLevel))
        switch n.Type {
        case IntLiteral:
            sb.WriteString(fmt.Sprintf("Literal (Int): %s\n", n.Value))
        case StringLiteral:
            sb.WriteString(fmt.Sprintf("Literal (String): \"%s\"\n", n.Value))
        default:
            sb.WriteString(fmt.Sprintf("Literal (Unknown): %s\n", n.Value))
        }
        return sb.String()

    case *VariableExpression:
        return indent(indentLevel) + fmt.Sprintf("VariableExpression: %s\n", n.Name)

    case *CallExpression:
        sb := &strings.Builder{}
        sb.WriteString(indent(indentLevel))
        sb.WriteString("CallExpression:\n")

        sb.WriteString(indent(indentLevel + 1))
        sb.WriteString("Callee:\n")
        sb.WriteString(stringifyNode(n.Callee, indentLevel+2))

        if len(n.Arguments) > 0 {
            sb.WriteString(indent(indentLevel + 1))
            sb.WriteString("Arguments:\n")
            for _, arg := range n.Arguments {
                sb.WriteString(stringifyNode(arg, indentLevel+2))
            }
        }
        return sb.String()

    default:
        return indent(indentLevel) + fmt.Sprintf("<unknown node type %T>\n", node)
    }
}

// indent produces a string of spaces for hierarchical indentation.
func indent(level int) string {
    return strings.Repeat("  ", level)
}

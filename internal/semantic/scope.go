package semantic

import (
    "fmt"

    "github.com/richardzmija/fusion-compiler/internal/ast"
)

// SymbolKind indicates whether a symbol is a variable or a function.
type SymbolKind int

const (
    SymbolVar SymbolKind = iota
    SymbolFunc
)

// Symbol holds information about a declared identifier.
type Symbol struct {
    Name       string
    Kind       SymbolKind
    DataType   ast.DataType   // For variables or return type for functions
    ParamTypes []ast.DataType // For function parameters; empty for variables
}

// Scope holds symbols for a given lexical scope and a pointer to the parent scope.
type Scope struct {
    symbols map[string]*Symbol
    parent  *Scope
}

// NewScope creates a new empty scope with a given parent. If parent is nil, this is a global scope.
func NewScope(parent *Scope) *Scope {
    return &Scope{
        symbols: make(map[string]*Symbol),
        parent:  parent,
    }
}

// Define attempts to insert a new symbol in the current scope.
// Returns an error if the symbol is already defined in the *current* scope.
func (s *Scope) Define(sym *Symbol) error {
    if _, exists := s.symbols[sym.Name]; exists {
        return fmt.Errorf("symbol '%s' already defined in this scope", sym.Name)
    }
    s.symbols[sym.Name] = sym
    return nil
}

// Resolve searches the current scope and all parent scopes for a symbol with the given name.
// Returns the symbol and nil if found, or nil and an error if not found.
func (s *Scope) Resolve(name string) (*Symbol, error) {
    for current := s; current != nil; current = current.parent {
        if sym, ok := current.symbols[name]; ok {
            return sym, nil
        }
    }
    return nil, fmt.Errorf("symbol '%s' not found in any enclosing scope", name)
}

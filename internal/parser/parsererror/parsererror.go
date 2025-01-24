package parsererror

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

// ErrorListener implements the ANTLR error listener interface
// to gather syntax errors during the parsing phase.
type ErrorListener struct {
	*antlr.DefaultErrorListener
	errors []string
}

// NewErrorListener constructs a new ErrorListener.
func NewErrorListener() *ErrorListener {
	return &ErrorListener{
		DefaultErrorListener: &antlr.DefaultErrorListener{},
		errors:               make([]string, 0),
	}
}

// SyntaxError is called by the parser when a syntax error occurs.
func (el *ErrorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	e antlr.RecognitionException,
) {
	errorMsg := fmt.Sprintf("line %d:%d: %s", line, column, msg)
	el.errors = append(el.errors, errorMsg)
}

// HasErrors indicates whether any syntax errors were recorded.
func (el *ErrorListener) HasErrors() bool {
	return len(el.errors) > 0
}

// GetErrors returns all recorded syntax errors.
func (el *ErrorListener) GetErrors() []string {
	return el.errors
}

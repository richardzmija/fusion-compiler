package integration_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/richardzmija/fusion-compiler/internal/parser"
)

// TestParserIntegration reads multiple C source files, parses them,
// and checks whether we expect any syntax errors.
func TestParserIntegration(t *testing.T) {
	testDataDir := "../testdata"

	testCases := []struct {
		name        string
		fileName    string
		shouldError bool
	}{
		{
			name:        "Valid Basic Function",
			fileName:    "valid_basic_function.c",
			shouldError: false,
		},
		{
			name:        "Valid No Return",
			fileName:    "valid_no_return.c",
			shouldError: false,
		},
		{
			name:        "Valid If Else",
			fileName:    "valid_if_else.c",
			shouldError: false,
		},
		{
			name:        "Valid While Loop",
			fileName:    "valid_while_loop.c",
			shouldError: false,
		},
		{
			name:        "Valid Function With Params",
			fileName:    "valid_function_with_params.c",
			shouldError: false,
		},
		{
			name:        "Invalid Missing Semicolon",
			fileName:    "invalid_missing_semicolon.c",
			shouldError: true,
		},
		{
			name:        "Invalid Identifier",
			fileName:    "invalid_identifier.c",
			shouldError: true,
		},
		{
			name:        "Invalid No Closing Brace",
			fileName:    "invalid_no_closing_brace.c",
			shouldError: true,
		},
		{
			name:        "Valid Multiple Statements",
			fileName:    "valid_multiple_statements.c",
			shouldError: false,
		},
		{
			name:        "Valid Nested Expressions",
			fileName:    "valid_nested_expressions.c",
			shouldError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			filePath := filepath.Join(testDataDir, tc.fileName)

			inputBytes, err := os.ReadFile(filePath)
			if err != nil {
				t.Fatalf("Failed to read file %s: %v", filePath, err)
			}

			// ANTLR requires a string as the input.
			antlrInput := antlr.NewInputStream(string(inputBytes))

			lex := parser.NewCLexer(antlrInput)
			stream := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
			p := parser.NewCParser(stream)

			// Remove default error listeners and add a custom listener.
			p.RemoveErrorListeners()
			errorListener := &ErrorListener{t: t}
			p.AddErrorListener(errorListener)

			// Parse the input from the program rule (root non-terminal).
			_ = p.Program()

			if tc.shouldError && !errorListener.hasError {
				t.Errorf("Expected a parsing error, but got none.")
			}
			if !tc.shouldError && errorListener.hasError {
				t.Errorf("Did not expect a parsing error, but one was reported.")
			}
		})
	}
}

// ErrorListener is a custom listener which records syntax errors.
type ErrorListener struct {
	*antlr.DefaultErrorListener
	t        *testing.T
	hasError bool
}

// SyntaxError is called by ANTLR when a syntax error is encountered.
func (l *ErrorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	e antlr.RecognitionException,
) {
	l.hasError = true
	l.t.Logf("Syntax error at %d:%d - %s", line, column, msg)
}

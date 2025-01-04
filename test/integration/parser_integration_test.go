package integration_test

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/richardzmija/fusion-compiler/internal/parser"
)

// TestParserBasic tests the basic parsing functionality of the generated parser.
func TestParserBasic(t *testing.T) {
	const enableLogging = false

	// Sample C code to parse.
	input := `
		int main() {
			int x = 10;
			printf("Value of x is %d\n", x);
			return 0;
		}
	`

	antlrInput := antlr.NewInputStream(input)
	lex := parser.NewCLexer(antlrInput)
	stream := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewCParser(stream)

	// Remove the default listeners and add a custom one.
	p.RemoveErrorListeners()
	errorListener := &ErrorListener{t: t}
	p.AddErrorListener(errorListener)

	// Parse the input string starting from the 'program' rule.
	tree := p.Program()

	if enableLogging {
		t.Log(tree.ToStringTree(nil, p))
	}

	if errorListener.hasError {
		t.Errorf("Parser encountered syntax errors")
	}
}

// Custom error listener that records syntax errors.
type ErrorListener struct {
	*antlr.DefaultErrorListener
	t        *testing.T
	hasError bool
}

// Called by ANTLR when a syntax error is encountered.
func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {

	l.hasError = true
	l.t.Errorf("Syntax error at %d:%d - %s", line, column, msg)
}

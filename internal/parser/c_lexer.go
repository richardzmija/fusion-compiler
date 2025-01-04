// Code generated from C.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var CLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func clexerLexerInit() {
	staticData := &CLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'int'", "'if'", "'else'", "'while'", "'return'", "'printf'", "'+'",
		"'-'", "'*'", "'/'", "'='", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='",
		"'&&'", "'||'", "';'", "','", "'('", "')'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "IF", "ELSE", "WHILE", "RETURN", "PRINTF", "PLUS", "MINUS",
		"MULT", "DIV", "ASSIGN", "EQ", "NEQ", "LT", "GT", "LE", "GE", "AND",
		"OR", "SEMI", "COMMA", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "ID",
		"NUM", "STR", "WS", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"INT", "IF", "ELSE", "WHILE", "RETURN", "PRINTF", "PLUS", "MINUS", "MULT",
		"DIV", "ASSIGN", "EQ", "NEQ", "LT", "GT", "LE", "GE", "AND", "OR", "SEMI",
		"COMMA", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "ID", "NUM", "STR",
		"WS", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 31, 199, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 5, 25,
		142, 8, 25, 10, 25, 12, 25, 145, 9, 25, 1, 26, 1, 26, 1, 26, 5, 26, 150,
		8, 26, 10, 26, 12, 26, 153, 9, 26, 3, 26, 155, 8, 26, 1, 27, 1, 27, 1,
		27, 1, 27, 5, 27, 161, 8, 27, 10, 27, 12, 27, 164, 9, 27, 1, 27, 1, 27,
		1, 28, 4, 28, 169, 8, 28, 11, 28, 12, 28, 170, 1, 28, 1, 28, 1, 29, 1,
		29, 1, 29, 1, 29, 5, 29, 179, 8, 29, 10, 29, 12, 29, 182, 9, 29, 1, 29,
		1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 190, 8, 30, 10, 30, 12, 30, 193,
		9, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 191, 0, 31, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43,
		22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61,
		31, 1, 0, 7, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95,
		97, 122, 1, 0, 49, 57, 1, 0, 48, 57, 2, 0, 34, 34, 92, 92, 3, 0, 9, 10,
		13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 206, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27,
		1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0,
		35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0,
		0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0,
		0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0,
		0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 1, 63, 1, 0, 0, 0, 3, 67, 1,
		0, 0, 0, 5, 70, 1, 0, 0, 0, 7, 75, 1, 0, 0, 0, 9, 81, 1, 0, 0, 0, 11, 88,
		1, 0, 0, 0, 13, 95, 1, 0, 0, 0, 15, 97, 1, 0, 0, 0, 17, 99, 1, 0, 0, 0,
		19, 101, 1, 0, 0, 0, 21, 103, 1, 0, 0, 0, 23, 105, 1, 0, 0, 0, 25, 108,
		1, 0, 0, 0, 27, 111, 1, 0, 0, 0, 29, 113, 1, 0, 0, 0, 31, 115, 1, 0, 0,
		0, 33, 118, 1, 0, 0, 0, 35, 121, 1, 0, 0, 0, 37, 124, 1, 0, 0, 0, 39, 127,
		1, 0, 0, 0, 41, 129, 1, 0, 0, 0, 43, 131, 1, 0, 0, 0, 45, 133, 1, 0, 0,
		0, 47, 135, 1, 0, 0, 0, 49, 137, 1, 0, 0, 0, 51, 139, 1, 0, 0, 0, 53, 154,
		1, 0, 0, 0, 55, 156, 1, 0, 0, 0, 57, 168, 1, 0, 0, 0, 59, 174, 1, 0, 0,
		0, 61, 185, 1, 0, 0, 0, 63, 64, 5, 105, 0, 0, 64, 65, 5, 110, 0, 0, 65,
		66, 5, 116, 0, 0, 66, 2, 1, 0, 0, 0, 67, 68, 5, 105, 0, 0, 68, 69, 5, 102,
		0, 0, 69, 4, 1, 0, 0, 0, 70, 71, 5, 101, 0, 0, 71, 72, 5, 108, 0, 0, 72,
		73, 5, 115, 0, 0, 73, 74, 5, 101, 0, 0, 74, 6, 1, 0, 0, 0, 75, 76, 5, 119,
		0, 0, 76, 77, 5, 104, 0, 0, 77, 78, 5, 105, 0, 0, 78, 79, 5, 108, 0, 0,
		79, 80, 5, 101, 0, 0, 80, 8, 1, 0, 0, 0, 81, 82, 5, 114, 0, 0, 82, 83,
		5, 101, 0, 0, 83, 84, 5, 116, 0, 0, 84, 85, 5, 117, 0, 0, 85, 86, 5, 114,
		0, 0, 86, 87, 5, 110, 0, 0, 87, 10, 1, 0, 0, 0, 88, 89, 5, 112, 0, 0, 89,
		90, 5, 114, 0, 0, 90, 91, 5, 105, 0, 0, 91, 92, 5, 110, 0, 0, 92, 93, 5,
		116, 0, 0, 93, 94, 5, 102, 0, 0, 94, 12, 1, 0, 0, 0, 95, 96, 5, 43, 0,
		0, 96, 14, 1, 0, 0, 0, 97, 98, 5, 45, 0, 0, 98, 16, 1, 0, 0, 0, 99, 100,
		5, 42, 0, 0, 100, 18, 1, 0, 0, 0, 101, 102, 5, 47, 0, 0, 102, 20, 1, 0,
		0, 0, 103, 104, 5, 61, 0, 0, 104, 22, 1, 0, 0, 0, 105, 106, 5, 61, 0, 0,
		106, 107, 5, 61, 0, 0, 107, 24, 1, 0, 0, 0, 108, 109, 5, 33, 0, 0, 109,
		110, 5, 61, 0, 0, 110, 26, 1, 0, 0, 0, 111, 112, 5, 60, 0, 0, 112, 28,
		1, 0, 0, 0, 113, 114, 5, 62, 0, 0, 114, 30, 1, 0, 0, 0, 115, 116, 5, 60,
		0, 0, 116, 117, 5, 61, 0, 0, 117, 32, 1, 0, 0, 0, 118, 119, 5, 62, 0, 0,
		119, 120, 5, 61, 0, 0, 120, 34, 1, 0, 0, 0, 121, 122, 5, 38, 0, 0, 122,
		123, 5, 38, 0, 0, 123, 36, 1, 0, 0, 0, 124, 125, 5, 124, 0, 0, 125, 126,
		5, 124, 0, 0, 126, 38, 1, 0, 0, 0, 127, 128, 5, 59, 0, 0, 128, 40, 1, 0,
		0, 0, 129, 130, 5, 44, 0, 0, 130, 42, 1, 0, 0, 0, 131, 132, 5, 40, 0, 0,
		132, 44, 1, 0, 0, 0, 133, 134, 5, 41, 0, 0, 134, 46, 1, 0, 0, 0, 135, 136,
		5, 123, 0, 0, 136, 48, 1, 0, 0, 0, 137, 138, 5, 125, 0, 0, 138, 50, 1,
		0, 0, 0, 139, 143, 7, 0, 0, 0, 140, 142, 7, 1, 0, 0, 141, 140, 1, 0, 0,
		0, 142, 145, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144,
		52, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 146, 155, 5, 48, 0, 0, 147, 151,
		7, 2, 0, 0, 148, 150, 7, 3, 0, 0, 149, 148, 1, 0, 0, 0, 150, 153, 1, 0,
		0, 0, 151, 149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 155, 1, 0, 0, 0,
		153, 151, 1, 0, 0, 0, 154, 146, 1, 0, 0, 0, 154, 147, 1, 0, 0, 0, 155,
		54, 1, 0, 0, 0, 156, 162, 5, 34, 0, 0, 157, 158, 5, 92, 0, 0, 158, 161,
		9, 0, 0, 0, 159, 161, 8, 4, 0, 0, 160, 157, 1, 0, 0, 0, 160, 159, 1, 0,
		0, 0, 161, 164, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0,
		163, 165, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 165, 166, 5, 34, 0, 0, 166,
		56, 1, 0, 0, 0, 167, 169, 7, 5, 0, 0, 168, 167, 1, 0, 0, 0, 169, 170, 1,
		0, 0, 0, 170, 168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 172, 1, 0, 0,
		0, 172, 173, 6, 28, 0, 0, 173, 58, 1, 0, 0, 0, 174, 175, 5, 47, 0, 0, 175,
		176, 5, 47, 0, 0, 176, 180, 1, 0, 0, 0, 177, 179, 8, 6, 0, 0, 178, 177,
		1, 0, 0, 0, 179, 182, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0,
		0, 0, 181, 183, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 183, 184, 6, 29, 0, 0,
		184, 60, 1, 0, 0, 0, 185, 186, 5, 47, 0, 0, 186, 187, 5, 42, 0, 0, 187,
		191, 1, 0, 0, 0, 188, 190, 9, 0, 0, 0, 189, 188, 1, 0, 0, 0, 190, 193,
		1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 192, 194, 1, 0,
		0, 0, 193, 191, 1, 0, 0, 0, 194, 195, 5, 42, 0, 0, 195, 196, 5, 47, 0,
		0, 196, 197, 1, 0, 0, 0, 197, 198, 6, 30, 0, 0, 198, 62, 1, 0, 0, 0, 9,
		0, 143, 151, 154, 160, 162, 170, 180, 191, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CLexerInit initializes any static state used to implement CLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CLexerInit() {
	staticData := &CLexerLexerStaticData
	staticData.once.Do(clexerLexerInit)
}

// NewCLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCLexer(input antlr.CharStream) *CLexer {
	CLexerInit()
	l := new(CLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "C.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CLexer tokens.
const (
	CLexerINT           = 1
	CLexerIF            = 2
	CLexerELSE          = 3
	CLexerWHILE         = 4
	CLexerRETURN        = 5
	CLexerPRINTF        = 6
	CLexerPLUS          = 7
	CLexerMINUS         = 8
	CLexerMULT          = 9
	CLexerDIV           = 10
	CLexerASSIGN        = 11
	CLexerEQ            = 12
	CLexerNEQ           = 13
	CLexerLT            = 14
	CLexerGT            = 15
	CLexerLE            = 16
	CLexerGE            = 17
	CLexerAND           = 18
	CLexerOR            = 19
	CLexerSEMI          = 20
	CLexerCOMMA         = 21
	CLexerLPAREN        = 22
	CLexerRPAREN        = 23
	CLexerLBRACE        = 24
	CLexerRBRACE        = 25
	CLexerID            = 26
	CLexerNUM           = 27
	CLexerSTR           = 28
	CLexerWS            = 29
	CLexerLINE_COMMENT  = 30
	CLexerBLOCK_COMMENT = 31
)

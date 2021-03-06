// Code generated from D:/Projects/datatable/grammar/Datatable.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 19, 131,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 7, 11, 68, 10, 11,
	12, 11, 14, 11, 71, 11, 11, 3, 12, 3, 12, 7, 12, 75, 10, 12, 12, 12, 14,
	12, 78, 11, 12, 3, 12, 3, 12, 3, 13, 6, 13, 83, 10, 13, 13, 13, 14, 13,
	84, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14,
	96, 10, 14, 3, 15, 3, 15, 3, 16, 6, 16, 101, 10, 16, 13, 16, 14, 16, 102,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 111, 10, 17, 12, 17, 14,
	17, 114, 11, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18,
	3, 18, 7, 18, 125, 10, 18, 12, 18, 14, 18, 128, 11, 18, 3, 18, 3, 18, 3,
	112, 2, 19, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 3, 2, 8,
	4, 2, 67, 92, 99, 124, 5, 2, 50, 59, 67, 92, 99, 124, 3, 2, 36, 36, 3,
	2, 50, 59, 5, 2, 11, 12, 15, 15, 34, 34, 4, 2, 12, 12, 15, 15, 2, 137,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2,
	2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2,
	2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3,
	2, 2, 2, 2, 35, 3, 2, 2, 2, 3, 37, 3, 2, 2, 2, 5, 39, 3, 2, 2, 2, 7, 41,
	3, 2, 2, 2, 9, 43, 3, 2, 2, 2, 11, 45, 3, 2, 2, 2, 13, 47, 3, 2, 2, 2,
	15, 49, 3, 2, 2, 2, 17, 54, 3, 2, 2, 2, 19, 60, 3, 2, 2, 2, 21, 65, 3,
	2, 2, 2, 23, 72, 3, 2, 2, 2, 25, 82, 3, 2, 2, 2, 27, 95, 3, 2, 2, 2, 29,
	97, 3, 2, 2, 2, 31, 100, 3, 2, 2, 2, 33, 106, 3, 2, 2, 2, 35, 120, 3, 2,
	2, 2, 37, 38, 7, 66, 2, 2, 38, 4, 3, 2, 2, 2, 39, 40, 7, 47, 2, 2, 40,
	6, 3, 2, 2, 2, 41, 42, 7, 65, 2, 2, 42, 8, 3, 2, 2, 2, 43, 44, 7, 60, 2,
	2, 44, 10, 3, 2, 2, 2, 45, 46, 7, 125, 2, 2, 46, 12, 3, 2, 2, 2, 47, 48,
	7, 127, 2, 2, 48, 14, 3, 2, 2, 2, 49, 50, 7, 86, 2, 2, 50, 51, 7, 123,
	2, 2, 51, 52, 7, 114, 2, 2, 52, 53, 7, 103, 2, 2, 53, 16, 3, 2, 2, 2, 54,
	55, 7, 73, 2, 2, 55, 56, 7, 116, 2, 2, 56, 57, 7, 113, 2, 2, 57, 58, 7,
	119, 2, 2, 58, 59, 7, 114, 2, 2, 59, 18, 3, 2, 2, 2, 60, 61, 7, 79, 2,
	2, 61, 62, 7, 103, 2, 2, 62, 63, 7, 118, 2, 2, 63, 64, 7, 99, 2, 2, 64,
	20, 3, 2, 2, 2, 65, 69, 9, 2, 2, 2, 66, 68, 9, 3, 2, 2, 67, 66, 3, 2, 2,
	2, 68, 71, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 22,
	3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 72, 76, 7, 36, 2, 2, 73, 75, 10, 4, 2,
	2, 74, 73, 3, 2, 2, 2, 75, 78, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 77,
	3, 2, 2, 2, 77, 79, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 79, 80, 7, 36, 2, 2,
	80, 24, 3, 2, 2, 2, 81, 83, 9, 5, 2, 2, 82, 81, 3, 2, 2, 2, 83, 84, 3,
	2, 2, 2, 84, 82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 26, 3, 2, 2, 2, 86,
	87, 7, 118, 2, 2, 87, 88, 7, 116, 2, 2, 88, 89, 7, 119, 2, 2, 89, 96, 7,
	103, 2, 2, 90, 91, 7, 104, 2, 2, 91, 92, 7, 99, 2, 2, 92, 93, 7, 110, 2,
	2, 93, 94, 7, 117, 2, 2, 94, 96, 7, 103, 2, 2, 95, 86, 3, 2, 2, 2, 95,
	90, 3, 2, 2, 2, 96, 28, 3, 2, 2, 2, 97, 98, 9, 5, 2, 2, 98, 30, 3, 2, 2,
	2, 99, 101, 9, 6, 2, 2, 100, 99, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102,
	100, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 105,
	8, 16, 2, 2, 105, 32, 3, 2, 2, 2, 106, 107, 7, 49, 2, 2, 107, 108, 7, 44,
	2, 2, 108, 112, 3, 2, 2, 2, 109, 111, 11, 2, 2, 2, 110, 109, 3, 2, 2, 2,
	111, 114, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 113,
	115, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 115, 116, 7, 44, 2, 2, 116, 117,
	7, 49, 2, 2, 117, 118, 3, 2, 2, 2, 118, 119, 8, 17, 2, 2, 119, 34, 3, 2,
	2, 2, 120, 121, 7, 49, 2, 2, 121, 122, 7, 49, 2, 2, 122, 126, 3, 2, 2,
	2, 123, 125, 10, 7, 2, 2, 124, 123, 3, 2, 2, 2, 125, 128, 3, 2, 2, 2, 126,
	124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 129, 3, 2, 2, 2, 128, 126,
	3, 2, 2, 2, 129, 130, 8, 18, 2, 2, 130, 36, 3, 2, 2, 2, 10, 2, 69, 76,
	84, 95, 102, 112, 126, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'@'", "'-'", "'?'", "':'", "'{'", "'}'", "'Type'", "'Group'", "'Meta'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "TYPE", "GROUP", "META", "NAME", "STRING",
	"INTEGER", "BOOLEAN", "Digit", "WS", "BlockComment", "LineComment",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "TYPE", "GROUP", "META",
	"NAME", "STRING", "INTEGER", "BOOLEAN", "Digit", "WS", "BlockComment",
	"LineComment",
}

type DatatableLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewDatatableLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *DatatableLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewDatatableLexer(input antlr.CharStream) *DatatableLexer {
	l := new(DatatableLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Datatable.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DatatableLexer tokens.
const (
	DatatableLexerT__0         = 1
	DatatableLexerT__1         = 2
	DatatableLexerT__2         = 3
	DatatableLexerT__3         = 4
	DatatableLexerT__4         = 5
	DatatableLexerT__5         = 6
	DatatableLexerTYPE         = 7
	DatatableLexerGROUP        = 8
	DatatableLexerMETA         = 9
	DatatableLexerNAME         = 10
	DatatableLexerSTRING       = 11
	DatatableLexerINTEGER      = 12
	DatatableLexerBOOLEAN      = 13
	DatatableLexerDigit        = 14
	DatatableLexerWS           = 15
	DatatableLexerBlockComment = 16
	DatatableLexerLineComment  = 17
)

// Code generated from Cal.g4 by ANTLR 4.9. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 9, 89, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 3, 2, 6, 2, 19, 10, 2, 13, 2, 14, 2, 20, 3, 3, 6, 3, 24,
	10, 3, 13, 3, 14, 3, 25, 3, 4, 3, 4, 3, 4, 5, 4, 31, 10, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 5, 5, 37, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 44, 10,
	6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5,
	7, 81, 10, 7, 3, 8, 6, 8, 84, 10, 8, 13, 8, 14, 8, 85, 3, 8, 3, 8, 2, 2,
	9, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 3, 2, 5, 3, 2, 50, 59,
	3, 2, 67, 92, 5, 2, 11, 12, 15, 15, 34, 34, 2, 97, 2, 3, 3, 2, 2, 2, 2,
	5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2,
	13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 3, 18, 3, 2, 2, 2, 5, 23, 3, 2, 2, 2,
	7, 30, 3, 2, 2, 2, 9, 36, 3, 2, 2, 2, 11, 43, 3, 2, 2, 2, 13, 80, 3, 2,
	2, 2, 15, 83, 3, 2, 2, 2, 17, 19, 9, 2, 2, 2, 18, 17, 3, 2, 2, 2, 19, 20,
	3, 2, 2, 2, 20, 18, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 4, 3, 2, 2, 2,
	22, 24, 9, 3, 2, 2, 23, 22, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 23, 3,
	2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 6, 3, 2, 2, 2, 27, 28, 7, 81, 2, 2, 28,
	31, 7, 84, 2, 2, 29, 31, 7, 126, 2, 2, 30, 27, 3, 2, 2, 2, 30, 29, 3, 2,
	2, 2, 31, 8, 3, 2, 2, 2, 32, 33, 7, 67, 2, 2, 33, 34, 7, 80, 2, 2, 34,
	37, 7, 70, 2, 2, 35, 37, 7, 40, 2, 2, 36, 32, 3, 2, 2, 2, 36, 35, 3, 2,
	2, 2, 37, 10, 3, 2, 2, 2, 38, 44, 4, 62, 64, 2, 39, 40, 7, 64, 2, 2, 40,
	44, 7, 63, 2, 2, 41, 42, 7, 62, 2, 2, 42, 44, 7, 63, 2, 2, 43, 38, 3, 2,
	2, 2, 43, 39, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 44, 12, 3, 2, 2, 2, 45, 46,
	7, 69, 2, 2, 46, 47, 7, 81, 2, 2, 47, 48, 7, 80, 2, 2, 48, 49, 7, 86, 2,
	2, 49, 50, 7, 67, 2, 2, 50, 51, 7, 75, 2, 2, 51, 52, 7, 80, 2, 2, 52, 81,
	7, 85, 2, 2, 53, 54, 7, 85, 2, 2, 54, 55, 7, 86, 2, 2, 55, 56, 7, 84, 2,
	2, 56, 57, 7, 75, 2, 2, 57, 58, 7, 80, 2, 2, 58, 59, 7, 73, 2, 2, 59, 60,
	7, 71, 2, 2, 60, 61, 7, 83, 2, 2, 61, 62, 7, 87, 2, 2, 62, 63, 7, 67, 2,
	2, 63, 81, 7, 78, 2, 2, 64, 65, 7, 85, 2, 2, 65, 66, 7, 86, 2, 2, 66, 67,
	7, 84, 2, 2, 67, 68, 7, 75, 2, 2, 68, 69, 7, 80, 2, 2, 69, 70, 7, 73, 2,
	2, 70, 71, 7, 85, 2, 2, 71, 72, 7, 86, 2, 2, 72, 73, 7, 67, 2, 2, 73, 74,
	7, 84, 2, 2, 74, 75, 7, 86, 2, 2, 75, 76, 7, 85, 2, 2, 76, 77, 7, 89, 2,
	2, 77, 78, 7, 75, 2, 2, 78, 79, 7, 86, 2, 2, 79, 81, 7, 74, 2, 2, 80, 45,
	3, 2, 2, 2, 80, 53, 3, 2, 2, 2, 80, 64, 3, 2, 2, 2, 81, 14, 3, 2, 2, 2,
	82, 84, 9, 4, 2, 2, 83, 82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 83, 3,
	2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 88, 8, 8, 2, 2, 88,
	16, 3, 2, 2, 2, 10, 2, 20, 25, 30, 36, 43, 80, 85, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames []string

var lexerSymbolicNames = []string{
	"", "MYINT", "NAME", "OR", "AND", "NUMCOMP", "STRINGCOMP", "WS",
}

var lexerRuleNames = []string{
	"MYINT", "NAME", "OR", "AND", "NUMCOMP", "STRINGCOMP", "WS",
}

type CalLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewCalLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *CalLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewCalLexer(input antlr.CharStream) *CalLexer {
	l := new(CalLexer)
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
	l.GrammarFileName = "Cal.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalLexer tokens.
const (
	CalLexerMYINT      = 1
	CalLexerNAME       = 2
	CalLexerOR         = 3
	CalLexerAND        = 4
	CalLexerNUMCOMP    = 5
	CalLexerSTRINGCOMP = 6
	CalLexerWS         = 7
)

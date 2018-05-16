// Code generated from Test.g4 by ANTLR 4.7.1. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 14, 56, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 6, 9, 43, 10, 9, 13, 9, 14, 9, 44, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 2, 2, 14,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23,
	13, 25, 14, 3, 2, 3, 5, 2, 11, 12, 15, 15, 34, 34, 2, 56, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2,
	2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 3, 27, 3,
	2, 2, 2, 5, 29, 3, 2, 2, 2, 7, 31, 3, 2, 2, 2, 9, 33, 3, 2, 2, 2, 11, 35,
	3, 2, 2, 2, 13, 37, 3, 2, 2, 2, 15, 39, 3, 2, 2, 2, 17, 42, 3, 2, 2, 2,
	19, 48, 3, 2, 2, 2, 21, 50, 3, 2, 2, 2, 23, 52, 3, 2, 2, 2, 25, 54, 3,
	2, 2, 2, 27, 28, 7, 61, 2, 2, 28, 4, 3, 2, 2, 2, 29, 30, 7, 67, 2, 2, 30,
	6, 3, 2, 2, 2, 31, 32, 7, 68, 2, 2, 32, 8, 3, 2, 2, 2, 33, 34, 7, 69, 2,
	2, 34, 10, 3, 2, 2, 2, 35, 36, 7, 70, 2, 2, 36, 12, 3, 2, 2, 2, 37, 38,
	7, 71, 2, 2, 38, 14, 3, 2, 2, 2, 39, 40, 7, 72, 2, 2, 40, 16, 3, 2, 2,
	2, 41, 43, 9, 2, 2, 2, 42, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 42,
	3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 47, 8, 9, 2, 2,
	47, 18, 3, 2, 2, 2, 48, 49, 7, 42, 2, 2, 49, 20, 3, 2, 2, 2, 50, 51, 7,
	43, 2, 2, 51, 22, 3, 2, 2, 2, 52, 53, 7, 45, 2, 2, 53, 24, 3, 2, 2, 2,
	54, 55, 7, 47, 2, 2, 55, 26, 3, 2, 2, 2, 4, 2, 44, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "';'", "'A'", "'B'", "'C'", "'D'", "'E'", "'F'", "", "'('", "')'",
	"'+'", "'-'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "WHITESPACE", "LB", "RB", "PLUS", "MINUS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "WHITESPACE", "LB",
	"RB", "PLUS", "MINUS",
}

type TestLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewTestLexer(input antlr.CharStream) *TestLexer {

	l := new(TestLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Test.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TestLexer tokens.
const (
	TestLexerT__0       = 1
	TestLexerT__1       = 2
	TestLexerT__2       = 3
	TestLexerT__3       = 4
	TestLexerT__4       = 5
	TestLexerT__5       = 6
	TestLexerT__6       = 7
	TestLexerWHITESPACE = 8
	TestLexerLB         = 9
	TestLexerRB         = 10
	TestLexerPLUS       = 11
	TestLexerMINUS      = 12
)

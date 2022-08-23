// Code generated from Language.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type LanguageLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var languagelexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func languagelexerLexerInit() {
	staticData := &languagelexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'{'", "'}'", "','", "';'",
		"'case'", "'of'", "'->'",
	}
	staticData.symbolicNames = []string{
		"", "MUL", "DIV", "ADD", "SUB", "PAREN_LEFT", "PAREN_RIGHT", "BRACE_LEFT",
		"BRACE_RIGHT", "COMMA", "SEMICOLON", "CASE", "OF", "ARROW", "INT", "CHAR",
		"STRING", "VARID", "CONID", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"MUL", "DIV", "ADD", "SUB", "PAREN_LEFT", "PAREN_RIGHT", "BRACE_LEFT",
		"BRACE_RIGHT", "COMMA", "SEMICOLON", "CASE", "OF", "ARROW", "INT", "CHAR",
		"STRING", "VARID", "CONID", "WHITESPACE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 19, 104, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 4, 13, 72, 8, 13, 11, 13, 12, 13, 73,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 5,
		16, 86, 8, 16, 10, 16, 12, 16, 89, 9, 16, 1, 17, 1, 17, 5, 17, 93, 8, 17,
		10, 17, 12, 17, 96, 9, 17, 1, 18, 4, 18, 99, 8, 18, 11, 18, 12, 18, 100,
		1, 18, 1, 18, 0, 0, 19, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 1, 0, 5, 1, 0, 48, 57, 1, 0, 97, 122, 4, 0, 48, 57, 65,
		90, 95, 95, 97, 122, 1, 0, 65, 90, 3, 0, 9, 10, 13, 13, 32, 32, 107, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0,
		0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0,
		0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 1, 39, 1,
		0, 0, 0, 3, 41, 1, 0, 0, 0, 5, 43, 1, 0, 0, 0, 7, 45, 1, 0, 0, 0, 9, 47,
		1, 0, 0, 0, 11, 49, 1, 0, 0, 0, 13, 51, 1, 0, 0, 0, 15, 53, 1, 0, 0, 0,
		17, 55, 1, 0, 0, 0, 19, 57, 1, 0, 0, 0, 21, 59, 1, 0, 0, 0, 23, 64, 1,
		0, 0, 0, 25, 67, 1, 0, 0, 0, 27, 71, 1, 0, 0, 0, 29, 75, 1, 0, 0, 0, 31,
		79, 1, 0, 0, 0, 33, 83, 1, 0, 0, 0, 35, 90, 1, 0, 0, 0, 37, 98, 1, 0, 0,
		0, 39, 40, 5, 42, 0, 0, 40, 2, 1, 0, 0, 0, 41, 42, 5, 47, 0, 0, 42, 4,
		1, 0, 0, 0, 43, 44, 5, 43, 0, 0, 44, 6, 1, 0, 0, 0, 45, 46, 5, 45, 0, 0,
		46, 8, 1, 0, 0, 0, 47, 48, 5, 40, 0, 0, 48, 10, 1, 0, 0, 0, 49, 50, 5,
		41, 0, 0, 50, 12, 1, 0, 0, 0, 51, 52, 5, 123, 0, 0, 52, 14, 1, 0, 0, 0,
		53, 54, 5, 125, 0, 0, 54, 16, 1, 0, 0, 0, 55, 56, 5, 44, 0, 0, 56, 18,
		1, 0, 0, 0, 57, 58, 5, 59, 0, 0, 58, 20, 1, 0, 0, 0, 59, 60, 5, 99, 0,
		0, 60, 61, 5, 97, 0, 0, 61, 62, 5, 115, 0, 0, 62, 63, 5, 101, 0, 0, 63,
		22, 1, 0, 0, 0, 64, 65, 5, 111, 0, 0, 65, 66, 5, 102, 0, 0, 66, 24, 1,
		0, 0, 0, 67, 68, 5, 45, 0, 0, 68, 69, 5, 62, 0, 0, 69, 26, 1, 0, 0, 0,
		70, 72, 7, 0, 0, 0, 71, 70, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 71, 1,
		0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 28, 1, 0, 0, 0, 75, 76, 5, 39, 0, 0, 76,
		77, 9, 0, 0, 0, 77, 78, 5, 39, 0, 0, 78, 30, 1, 0, 0, 0, 79, 80, 5, 34,
		0, 0, 80, 81, 9, 0, 0, 0, 81, 82, 5, 34, 0, 0, 82, 32, 1, 0, 0, 0, 83,
		87, 7, 1, 0, 0, 84, 86, 7, 2, 0, 0, 85, 84, 1, 0, 0, 0, 86, 89, 1, 0, 0,
		0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 34, 1, 0, 0, 0, 89, 87,
		1, 0, 0, 0, 90, 94, 7, 3, 0, 0, 91, 93, 7, 2, 0, 0, 92, 91, 1, 0, 0, 0,
		93, 96, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 36, 1,
		0, 0, 0, 96, 94, 1, 0, 0, 0, 97, 99, 7, 4, 0, 0, 98, 97, 1, 0, 0, 0, 99,
		100, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 102, 1,
		0, 0, 0, 102, 103, 6, 18, 0, 0, 103, 38, 1, 0, 0, 0, 5, 0, 73, 87, 94,
		100, 1, 6, 0, 0,
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

// LanguageLexerInit initializes any static state used to implement LanguageLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLanguageLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func LanguageLexerInit() {
	staticData := &languagelexerLexerStaticData
	staticData.once.Do(languagelexerLexerInit)
}

// NewLanguageLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLanguageLexer(input antlr.CharStream) *LanguageLexer {
	LanguageLexerInit()
	l := new(LanguageLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &languagelexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Language.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LanguageLexer tokens.
const (
	LanguageLexerMUL         = 1
	LanguageLexerDIV         = 2
	LanguageLexerADD         = 3
	LanguageLexerSUB         = 4
	LanguageLexerPAREN_LEFT  = 5
	LanguageLexerPAREN_RIGHT = 6
	LanguageLexerBRACE_LEFT  = 7
	LanguageLexerBRACE_RIGHT = 8
	LanguageLexerCOMMA       = 9
	LanguageLexerSEMICOLON   = 10
	LanguageLexerCASE        = 11
	LanguageLexerOF          = 12
	LanguageLexerARROW       = 13
	LanguageLexerINT         = 14
	LanguageLexerCHAR        = 15
	LanguageLexerSTRING      = 16
	LanguageLexerVARID       = 17
	LanguageLexerCONID       = 18
	LanguageLexerWHITESPACE  = 19
)

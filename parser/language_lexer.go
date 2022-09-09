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
		"", "'*'", "'/'", "'+'", "'-'", "'='", "'('", "')'", "'{'", "'}'", "','",
		"';'", "'|'", "'::'", "'data'", "'case'", "'of'", "'->'",
	}
	staticData.symbolicNames = []string{
		"", "MUL", "DIV", "ADD", "SUB", "ASSIGN", "PAREN_LEFT", "PAREN_RIGHT",
		"BRACE_LEFT", "BRACE_RIGHT", "COMMA", "SEMICOLON", "VERTICAL_BAR", "DOUBLE_COLON",
		"DATA", "CASE", "OF", "ARROW", "INT", "CHAR", "STRING", "VARID", "CONID",
		"WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"MUL", "DIV", "ADD", "SUB", "ASSIGN", "PAREN_LEFT", "PAREN_RIGHT", "BRACE_LEFT",
		"BRACE_RIGHT", "COMMA", "SEMICOLON", "VERTICAL_BAR", "DOUBLE_COLON",
		"DATA", "CASE", "OF", "ARROW", "INT", "CHAR", "STRING", "VARID", "CONID",
		"WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 25, 158, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 4, 17, 96,
		8, 17, 11, 17, 12, 17, 97, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 5,
		19, 106, 8, 19, 10, 19, 12, 19, 109, 9, 19, 1, 19, 1, 19, 1, 20, 1, 20,
		5, 20, 115, 8, 20, 10, 20, 12, 20, 118, 9, 20, 1, 21, 1, 21, 5, 21, 122,
		8, 21, 10, 21, 12, 21, 125, 9, 21, 1, 22, 4, 22, 128, 8, 22, 11, 22, 12,
		22, 129, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 138, 8, 23, 10,
		23, 12, 23, 141, 9, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24,
		1, 24, 1, 24, 5, 24, 152, 8, 24, 10, 24, 12, 24, 155, 9, 24, 1, 24, 1,
		24, 2, 107, 139, 0, 25, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 1, 0, 6,
		1, 0, 48, 57, 1, 0, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1,
		0, 65, 90, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 164, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0,
		0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1,
		0, 0, 0, 0, 49, 1, 0, 0, 0, 1, 51, 1, 0, 0, 0, 3, 53, 1, 0, 0, 0, 5, 55,
		1, 0, 0, 0, 7, 57, 1, 0, 0, 0, 9, 59, 1, 0, 0, 0, 11, 61, 1, 0, 0, 0, 13,
		63, 1, 0, 0, 0, 15, 65, 1, 0, 0, 0, 17, 67, 1, 0, 0, 0, 19, 69, 1, 0, 0,
		0, 21, 71, 1, 0, 0, 0, 23, 73, 1, 0, 0, 0, 25, 75, 1, 0, 0, 0, 27, 78,
		1, 0, 0, 0, 29, 83, 1, 0, 0, 0, 31, 88, 1, 0, 0, 0, 33, 91, 1, 0, 0, 0,
		35, 95, 1, 0, 0, 0, 37, 99, 1, 0, 0, 0, 39, 103, 1, 0, 0, 0, 41, 112, 1,
		0, 0, 0, 43, 119, 1, 0, 0, 0, 45, 127, 1, 0, 0, 0, 47, 133, 1, 0, 0, 0,
		49, 147, 1, 0, 0, 0, 51, 52, 5, 42, 0, 0, 52, 2, 1, 0, 0, 0, 53, 54, 5,
		47, 0, 0, 54, 4, 1, 0, 0, 0, 55, 56, 5, 43, 0, 0, 56, 6, 1, 0, 0, 0, 57,
		58, 5, 45, 0, 0, 58, 8, 1, 0, 0, 0, 59, 60, 5, 61, 0, 0, 60, 10, 1, 0,
		0, 0, 61, 62, 5, 40, 0, 0, 62, 12, 1, 0, 0, 0, 63, 64, 5, 41, 0, 0, 64,
		14, 1, 0, 0, 0, 65, 66, 5, 123, 0, 0, 66, 16, 1, 0, 0, 0, 67, 68, 5, 125,
		0, 0, 68, 18, 1, 0, 0, 0, 69, 70, 5, 44, 0, 0, 70, 20, 1, 0, 0, 0, 71,
		72, 5, 59, 0, 0, 72, 22, 1, 0, 0, 0, 73, 74, 5, 124, 0, 0, 74, 24, 1, 0,
		0, 0, 75, 76, 5, 58, 0, 0, 76, 77, 5, 58, 0, 0, 77, 26, 1, 0, 0, 0, 78,
		79, 5, 100, 0, 0, 79, 80, 5, 97, 0, 0, 80, 81, 5, 116, 0, 0, 81, 82, 5,
		97, 0, 0, 82, 28, 1, 0, 0, 0, 83, 84, 5, 99, 0, 0, 84, 85, 5, 97, 0, 0,
		85, 86, 5, 115, 0, 0, 86, 87, 5, 101, 0, 0, 87, 30, 1, 0, 0, 0, 88, 89,
		5, 111, 0, 0, 89, 90, 5, 102, 0, 0, 90, 32, 1, 0, 0, 0, 91, 92, 5, 45,
		0, 0, 92, 93, 5, 62, 0, 0, 93, 34, 1, 0, 0, 0, 94, 96, 7, 0, 0, 0, 95,
		94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0,
		0, 98, 36, 1, 0, 0, 0, 99, 100, 5, 39, 0, 0, 100, 101, 9, 0, 0, 0, 101,
		102, 5, 39, 0, 0, 102, 38, 1, 0, 0, 0, 103, 107, 5, 34, 0, 0, 104, 106,
		9, 0, 0, 0, 105, 104, 1, 0, 0, 0, 106, 109, 1, 0, 0, 0, 107, 108, 1, 0,
		0, 0, 107, 105, 1, 0, 0, 0, 108, 110, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0,
		110, 111, 5, 34, 0, 0, 111, 40, 1, 0, 0, 0, 112, 116, 7, 1, 0, 0, 113,
		115, 7, 2, 0, 0, 114, 113, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 114,
		1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 42, 1, 0, 0, 0, 118, 116, 1, 0,
		0, 0, 119, 123, 7, 3, 0, 0, 120, 122, 7, 2, 0, 0, 121, 120, 1, 0, 0, 0,
		122, 125, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124,
		44, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 126, 128, 7, 4, 0, 0, 127, 126, 1,
		0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 129, 130, 1, 0, 0,
		0, 130, 131, 1, 0, 0, 0, 131, 132, 6, 22, 0, 0, 132, 46, 1, 0, 0, 0, 133,
		134, 5, 123, 0, 0, 134, 135, 5, 45, 0, 0, 135, 139, 1, 0, 0, 0, 136, 138,
		9, 0, 0, 0, 137, 136, 1, 0, 0, 0, 138, 141, 1, 0, 0, 0, 139, 140, 1, 0,
		0, 0, 139, 137, 1, 0, 0, 0, 140, 142, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0,
		142, 143, 5, 45, 0, 0, 143, 144, 5, 125, 0, 0, 144, 145, 1, 0, 0, 0, 145,
		146, 6, 23, 1, 0, 146, 48, 1, 0, 0, 0, 147, 148, 5, 45, 0, 0, 148, 149,
		5, 45, 0, 0, 149, 153, 1, 0, 0, 0, 150, 152, 8, 5, 0, 0, 151, 150, 1, 0,
		0, 0, 152, 155, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0,
		154, 156, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 156, 157, 6, 24, 1, 0, 157,
		50, 1, 0, 0, 0, 8, 0, 97, 107, 116, 123, 129, 139, 153, 2, 6, 0, 0, 0,
		1, 0,
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
	LanguageLexerMUL          = 1
	LanguageLexerDIV          = 2
	LanguageLexerADD          = 3
	LanguageLexerSUB          = 4
	LanguageLexerASSIGN       = 5
	LanguageLexerPAREN_LEFT   = 6
	LanguageLexerPAREN_RIGHT  = 7
	LanguageLexerBRACE_LEFT   = 8
	LanguageLexerBRACE_RIGHT  = 9
	LanguageLexerCOMMA        = 10
	LanguageLexerSEMICOLON    = 11
	LanguageLexerVERTICAL_BAR = 12
	LanguageLexerDOUBLE_COLON = 13
	LanguageLexerDATA         = 14
	LanguageLexerCASE         = 15
	LanguageLexerOF           = 16
	LanguageLexerARROW        = 17
	LanguageLexerINT          = 18
	LanguageLexerCHAR         = 19
	LanguageLexerSTRING       = 20
	LanguageLexerVARID        = 21
	LanguageLexerCONID        = 22
	LanguageLexerWHITESPACE   = 23
	LanguageLexerCOMMENT      = 24
	LanguageLexerLINE_COMMENT = 25
)

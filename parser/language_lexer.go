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
		"", "'*'", "'/'", "'+'", "'-'", "'='", "'<'", "'>'", "'<='", "'>='",
		"'=='", "'/='", "'&&'", "'||'", "'('", "')'", "'{'", "'}'", "','", "';'",
		"'|'", "'::'", "'import'", "'data'", "'case'", "'of'", "'->'",
	}
	staticData.symbolicNames = []string{
		"", "MUL", "DIV", "ADD", "SUB", "ASSIGN", "LT", "GT", "LTE", "GTE",
		"EQ", "NEQ", "AND", "OR", "PAREN_LEFT", "PAREN_RIGHT", "BRACE_LEFT",
		"BRACE_RIGHT", "COMMA", "SEMICOLON", "VERTICAL_BAR", "DOUBLE_COLON",
		"IMPORT", "DATA", "CASE", "OF", "ARROW", "INT", "CHAR", "STRING", "VARID",
		"CONID", "WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"MUL", "DIV", "ADD", "SUB", "ASSIGN", "LT", "GT", "LTE", "GTE", "EQ",
		"NEQ", "AND", "OR", "PAREN_LEFT", "PAREN_RIGHT", "BRACE_LEFT", "BRACE_RIGHT",
		"COMMA", "SEMICOLON", "VERTICAL_BAR", "DOUBLE_COLON", "IMPORT", "DATA",
		"CASE", "OF", "ARROW", "INT", "CHAR", "STRING", "VARID", "CONID", "WHITESPACE",
		"COMMENT", "LINE_COMMENT",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 34, 205, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1,
		25, 1, 25, 1, 26, 4, 26, 143, 8, 26, 11, 26, 12, 26, 144, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 28, 1, 28, 5, 28, 153, 8, 28, 10, 28, 12, 28, 156, 9,
		28, 1, 28, 1, 28, 1, 29, 1, 29, 5, 29, 162, 8, 29, 10, 29, 12, 29, 165,
		9, 29, 1, 30, 1, 30, 5, 30, 169, 8, 30, 10, 30, 12, 30, 172, 9, 30, 1,
		31, 4, 31, 175, 8, 31, 11, 31, 12, 31, 176, 1, 31, 1, 31, 1, 32, 1, 32,
		1, 32, 1, 32, 5, 32, 185, 8, 32, 10, 32, 12, 32, 188, 9, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 5, 33, 199, 8, 33,
		10, 33, 12, 33, 202, 9, 33, 1, 33, 1, 33, 2, 154, 186, 0, 34, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12,
		25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21,
		43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30,
		61, 31, 63, 32, 65, 33, 67, 34, 1, 0, 6, 1, 0, 48, 57, 1, 0, 97, 122, 4,
		0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 65, 90, 3, 0, 9, 10, 13, 13,
		32, 32, 2, 0, 10, 10, 13, 13, 211, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0,
		0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0,
		0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0,
		0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0,
		0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1,
		0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43,
		1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0,
		51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0,
		0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0,
		0, 0, 67, 1, 0, 0, 0, 1, 69, 1, 0, 0, 0, 3, 71, 1, 0, 0, 0, 5, 73, 1, 0,
		0, 0, 7, 75, 1, 0, 0, 0, 9, 77, 1, 0, 0, 0, 11, 79, 1, 0, 0, 0, 13, 81,
		1, 0, 0, 0, 15, 83, 1, 0, 0, 0, 17, 86, 1, 0, 0, 0, 19, 89, 1, 0, 0, 0,
		21, 92, 1, 0, 0, 0, 23, 95, 1, 0, 0, 0, 25, 98, 1, 0, 0, 0, 27, 101, 1,
		0, 0, 0, 29, 103, 1, 0, 0, 0, 31, 105, 1, 0, 0, 0, 33, 107, 1, 0, 0, 0,
		35, 109, 1, 0, 0, 0, 37, 111, 1, 0, 0, 0, 39, 113, 1, 0, 0, 0, 41, 115,
		1, 0, 0, 0, 43, 118, 1, 0, 0, 0, 45, 125, 1, 0, 0, 0, 47, 130, 1, 0, 0,
		0, 49, 135, 1, 0, 0, 0, 51, 138, 1, 0, 0, 0, 53, 142, 1, 0, 0, 0, 55, 146,
		1, 0, 0, 0, 57, 150, 1, 0, 0, 0, 59, 159, 1, 0, 0, 0, 61, 166, 1, 0, 0,
		0, 63, 174, 1, 0, 0, 0, 65, 180, 1, 0, 0, 0, 67, 194, 1, 0, 0, 0, 69, 70,
		5, 42, 0, 0, 70, 2, 1, 0, 0, 0, 71, 72, 5, 47, 0, 0, 72, 4, 1, 0, 0, 0,
		73, 74, 5, 43, 0, 0, 74, 6, 1, 0, 0, 0, 75, 76, 5, 45, 0, 0, 76, 8, 1,
		0, 0, 0, 77, 78, 5, 61, 0, 0, 78, 10, 1, 0, 0, 0, 79, 80, 5, 60, 0, 0,
		80, 12, 1, 0, 0, 0, 81, 82, 5, 62, 0, 0, 82, 14, 1, 0, 0, 0, 83, 84, 5,
		60, 0, 0, 84, 85, 5, 61, 0, 0, 85, 16, 1, 0, 0, 0, 86, 87, 5, 62, 0, 0,
		87, 88, 5, 61, 0, 0, 88, 18, 1, 0, 0, 0, 89, 90, 5, 61, 0, 0, 90, 91, 5,
		61, 0, 0, 91, 20, 1, 0, 0, 0, 92, 93, 5, 47, 0, 0, 93, 94, 5, 61, 0, 0,
		94, 22, 1, 0, 0, 0, 95, 96, 5, 38, 0, 0, 96, 97, 5, 38, 0, 0, 97, 24, 1,
		0, 0, 0, 98, 99, 5, 124, 0, 0, 99, 100, 5, 124, 0, 0, 100, 26, 1, 0, 0,
		0, 101, 102, 5, 40, 0, 0, 102, 28, 1, 0, 0, 0, 103, 104, 5, 41, 0, 0, 104,
		30, 1, 0, 0, 0, 105, 106, 5, 123, 0, 0, 106, 32, 1, 0, 0, 0, 107, 108,
		5, 125, 0, 0, 108, 34, 1, 0, 0, 0, 109, 110, 5, 44, 0, 0, 110, 36, 1, 0,
		0, 0, 111, 112, 5, 59, 0, 0, 112, 38, 1, 0, 0, 0, 113, 114, 5, 124, 0,
		0, 114, 40, 1, 0, 0, 0, 115, 116, 5, 58, 0, 0, 116, 117, 5, 58, 0, 0, 117,
		42, 1, 0, 0, 0, 118, 119, 5, 105, 0, 0, 119, 120, 5, 109, 0, 0, 120, 121,
		5, 112, 0, 0, 121, 122, 5, 111, 0, 0, 122, 123, 5, 114, 0, 0, 123, 124,
		5, 116, 0, 0, 124, 44, 1, 0, 0, 0, 125, 126, 5, 100, 0, 0, 126, 127, 5,
		97, 0, 0, 127, 128, 5, 116, 0, 0, 128, 129, 5, 97, 0, 0, 129, 46, 1, 0,
		0, 0, 130, 131, 5, 99, 0, 0, 131, 132, 5, 97, 0, 0, 132, 133, 5, 115, 0,
		0, 133, 134, 5, 101, 0, 0, 134, 48, 1, 0, 0, 0, 135, 136, 5, 111, 0, 0,
		136, 137, 5, 102, 0, 0, 137, 50, 1, 0, 0, 0, 138, 139, 5, 45, 0, 0, 139,
		140, 5, 62, 0, 0, 140, 52, 1, 0, 0, 0, 141, 143, 7, 0, 0, 0, 142, 141,
		1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 144, 145, 1, 0,
		0, 0, 145, 54, 1, 0, 0, 0, 146, 147, 5, 39, 0, 0, 147, 148, 9, 0, 0, 0,
		148, 149, 5, 39, 0, 0, 149, 56, 1, 0, 0, 0, 150, 154, 5, 34, 0, 0, 151,
		153, 9, 0, 0, 0, 152, 151, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154, 155,
		1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 155, 157, 1, 0, 0, 0, 156, 154, 1, 0,
		0, 0, 157, 158, 5, 34, 0, 0, 158, 58, 1, 0, 0, 0, 159, 163, 7, 1, 0, 0,
		160, 162, 7, 2, 0, 0, 161, 160, 1, 0, 0, 0, 162, 165, 1, 0, 0, 0, 163,
		161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 60, 1, 0, 0, 0, 165, 163, 1,
		0, 0, 0, 166, 170, 7, 3, 0, 0, 167, 169, 7, 2, 0, 0, 168, 167, 1, 0, 0,
		0, 169, 172, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171,
		62, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 173, 175, 7, 4, 0, 0, 174, 173, 1,
		0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0,
		0, 177, 178, 1, 0, 0, 0, 178, 179, 6, 31, 0, 0, 179, 64, 1, 0, 0, 0, 180,
		181, 5, 123, 0, 0, 181, 182, 5, 45, 0, 0, 182, 186, 1, 0, 0, 0, 183, 185,
		9, 0, 0, 0, 184, 183, 1, 0, 0, 0, 185, 188, 1, 0, 0, 0, 186, 187, 1, 0,
		0, 0, 186, 184, 1, 0, 0, 0, 187, 189, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0,
		189, 190, 5, 45, 0, 0, 190, 191, 5, 125, 0, 0, 191, 192, 1, 0, 0, 0, 192,
		193, 6, 32, 1, 0, 193, 66, 1, 0, 0, 0, 194, 195, 5, 45, 0, 0, 195, 196,
		5, 45, 0, 0, 196, 200, 1, 0, 0, 0, 197, 199, 8, 5, 0, 0, 198, 197, 1, 0,
		0, 0, 199, 202, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0,
		201, 203, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 203, 204, 6, 33, 1, 0, 204,
		68, 1, 0, 0, 0, 8, 0, 144, 154, 163, 170, 176, 186, 200, 2, 6, 0, 0, 0,
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
	LanguageLexerLT           = 6
	LanguageLexerGT           = 7
	LanguageLexerLTE          = 8
	LanguageLexerGTE          = 9
	LanguageLexerEQ           = 10
	LanguageLexerNEQ          = 11
	LanguageLexerAND          = 12
	LanguageLexerOR           = 13
	LanguageLexerPAREN_LEFT   = 14
	LanguageLexerPAREN_RIGHT  = 15
	LanguageLexerBRACE_LEFT   = 16
	LanguageLexerBRACE_RIGHT  = 17
	LanguageLexerCOMMA        = 18
	LanguageLexerSEMICOLON    = 19
	LanguageLexerVERTICAL_BAR = 20
	LanguageLexerDOUBLE_COLON = 21
	LanguageLexerIMPORT       = 22
	LanguageLexerDATA         = 23
	LanguageLexerCASE         = 24
	LanguageLexerOF           = 25
	LanguageLexerARROW        = 26
	LanguageLexerINT          = 27
	LanguageLexerCHAR         = 28
	LanguageLexerSTRING       = 29
	LanguageLexerVARID        = 30
	LanguageLexerCONID        = 31
	LanguageLexerWHITESPACE   = 32
	LanguageLexerCOMMENT      = 33
	LanguageLexerLINE_COMMENT = 34
)

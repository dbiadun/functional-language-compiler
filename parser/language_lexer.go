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
		"'=='", "'/='", "'&&'", "'||'", "'()'", "'('", "')'", "'{'", "'}'",
		"','", "';'", "'|'", "'::'", "'import'", "'data'", "'do'", "'case'",
		"'of'", "'->'", "'<-'",
	}
	staticData.symbolicNames = []string{
		"", "MUL", "DIV", "ADD", "SUB", "ASSIGN", "LT", "GT", "LTE", "GTE",
		"EQ", "NEQ", "AND", "OR", "PARENS", "PAREN_LEFT", "PAREN_RIGHT", "BRACE_LEFT",
		"BRACE_RIGHT", "COMMA", "SEMICOLON", "VERTICAL_BAR", "DOUBLE_COLON",
		"IMPORT", "DATA", "DO", "CASE", "OF", "ARROW_RIGHT", "ARROW_LEFT", "INT",
		"CHAR", "STRING", "VARID", "CONID", "WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"MUL", "DIV", "ADD", "SUB", "ASSIGN", "LT", "GT", "LTE", "GTE", "EQ",
		"NEQ", "AND", "OR", "PARENS", "PAREN_LEFT", "PAREN_RIGHT", "BRACE_LEFT",
		"BRACE_RIGHT", "COMMA", "SEMICOLON", "VERTICAL_BAR", "DOUBLE_COLON",
		"IMPORT", "DATA", "DO", "CASE", "OF", "ARROW_RIGHT", "ARROW_LEFT", "INT",
		"CHAR", "STRING", "VARID", "CONID", "WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 37, 220, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1,
		27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 29, 4, 29, 158, 8, 29, 11, 29,
		12, 29, 159, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 5, 31, 168, 8, 31,
		10, 31, 12, 31, 171, 9, 31, 1, 31, 1, 31, 1, 32, 1, 32, 5, 32, 177, 8,
		32, 10, 32, 12, 32, 180, 9, 32, 1, 33, 1, 33, 5, 33, 184, 8, 33, 10, 33,
		12, 33, 187, 9, 33, 1, 34, 4, 34, 190, 8, 34, 11, 34, 12, 34, 191, 1, 34,
		1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 5, 35, 200, 8, 35, 10, 35, 12, 35, 203,
		9, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 5,
		36, 214, 8, 36, 10, 36, 12, 36, 217, 9, 36, 1, 36, 1, 36, 2, 169, 201,
		0, 37, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28,
		57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37,
		1, 0, 6, 1, 0, 48, 57, 1, 0, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97,
		122, 1, 0, 65, 90, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 226,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0,
		47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0,
		0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0,
		0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0,
		0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 1, 75, 1, 0, 0, 0, 3, 77, 1,
		0, 0, 0, 5, 79, 1, 0, 0, 0, 7, 81, 1, 0, 0, 0, 9, 83, 1, 0, 0, 0, 11, 85,
		1, 0, 0, 0, 13, 87, 1, 0, 0, 0, 15, 89, 1, 0, 0, 0, 17, 92, 1, 0, 0, 0,
		19, 95, 1, 0, 0, 0, 21, 98, 1, 0, 0, 0, 23, 101, 1, 0, 0, 0, 25, 104, 1,
		0, 0, 0, 27, 107, 1, 0, 0, 0, 29, 110, 1, 0, 0, 0, 31, 112, 1, 0, 0, 0,
		33, 114, 1, 0, 0, 0, 35, 116, 1, 0, 0, 0, 37, 118, 1, 0, 0, 0, 39, 120,
		1, 0, 0, 0, 41, 122, 1, 0, 0, 0, 43, 124, 1, 0, 0, 0, 45, 127, 1, 0, 0,
		0, 47, 134, 1, 0, 0, 0, 49, 139, 1, 0, 0, 0, 51, 142, 1, 0, 0, 0, 53, 147,
		1, 0, 0, 0, 55, 150, 1, 0, 0, 0, 57, 153, 1, 0, 0, 0, 59, 157, 1, 0, 0,
		0, 61, 161, 1, 0, 0, 0, 63, 165, 1, 0, 0, 0, 65, 174, 1, 0, 0, 0, 67, 181,
		1, 0, 0, 0, 69, 189, 1, 0, 0, 0, 71, 195, 1, 0, 0, 0, 73, 209, 1, 0, 0,
		0, 75, 76, 5, 42, 0, 0, 76, 2, 1, 0, 0, 0, 77, 78, 5, 47, 0, 0, 78, 4,
		1, 0, 0, 0, 79, 80, 5, 43, 0, 0, 80, 6, 1, 0, 0, 0, 81, 82, 5, 45, 0, 0,
		82, 8, 1, 0, 0, 0, 83, 84, 5, 61, 0, 0, 84, 10, 1, 0, 0, 0, 85, 86, 5,
		60, 0, 0, 86, 12, 1, 0, 0, 0, 87, 88, 5, 62, 0, 0, 88, 14, 1, 0, 0, 0,
		89, 90, 5, 60, 0, 0, 90, 91, 5, 61, 0, 0, 91, 16, 1, 0, 0, 0, 92, 93, 5,
		62, 0, 0, 93, 94, 5, 61, 0, 0, 94, 18, 1, 0, 0, 0, 95, 96, 5, 61, 0, 0,
		96, 97, 5, 61, 0, 0, 97, 20, 1, 0, 0, 0, 98, 99, 5, 47, 0, 0, 99, 100,
		5, 61, 0, 0, 100, 22, 1, 0, 0, 0, 101, 102, 5, 38, 0, 0, 102, 103, 5, 38,
		0, 0, 103, 24, 1, 0, 0, 0, 104, 105, 5, 124, 0, 0, 105, 106, 5, 124, 0,
		0, 106, 26, 1, 0, 0, 0, 107, 108, 5, 40, 0, 0, 108, 109, 5, 41, 0, 0, 109,
		28, 1, 0, 0, 0, 110, 111, 5, 40, 0, 0, 111, 30, 1, 0, 0, 0, 112, 113, 5,
		41, 0, 0, 113, 32, 1, 0, 0, 0, 114, 115, 5, 123, 0, 0, 115, 34, 1, 0, 0,
		0, 116, 117, 5, 125, 0, 0, 117, 36, 1, 0, 0, 0, 118, 119, 5, 44, 0, 0,
		119, 38, 1, 0, 0, 0, 120, 121, 5, 59, 0, 0, 121, 40, 1, 0, 0, 0, 122, 123,
		5, 124, 0, 0, 123, 42, 1, 0, 0, 0, 124, 125, 5, 58, 0, 0, 125, 126, 5,
		58, 0, 0, 126, 44, 1, 0, 0, 0, 127, 128, 5, 105, 0, 0, 128, 129, 5, 109,
		0, 0, 129, 130, 5, 112, 0, 0, 130, 131, 5, 111, 0, 0, 131, 132, 5, 114,
		0, 0, 132, 133, 5, 116, 0, 0, 133, 46, 1, 0, 0, 0, 134, 135, 5, 100, 0,
		0, 135, 136, 5, 97, 0, 0, 136, 137, 5, 116, 0, 0, 137, 138, 5, 97, 0, 0,
		138, 48, 1, 0, 0, 0, 139, 140, 5, 100, 0, 0, 140, 141, 5, 111, 0, 0, 141,
		50, 1, 0, 0, 0, 142, 143, 5, 99, 0, 0, 143, 144, 5, 97, 0, 0, 144, 145,
		5, 115, 0, 0, 145, 146, 5, 101, 0, 0, 146, 52, 1, 0, 0, 0, 147, 148, 5,
		111, 0, 0, 148, 149, 5, 102, 0, 0, 149, 54, 1, 0, 0, 0, 150, 151, 5, 45,
		0, 0, 151, 152, 5, 62, 0, 0, 152, 56, 1, 0, 0, 0, 153, 154, 5, 60, 0, 0,
		154, 155, 5, 45, 0, 0, 155, 58, 1, 0, 0, 0, 156, 158, 7, 0, 0, 0, 157,
		156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159, 160,
		1, 0, 0, 0, 160, 60, 1, 0, 0, 0, 161, 162, 5, 39, 0, 0, 162, 163, 9, 0,
		0, 0, 163, 164, 5, 39, 0, 0, 164, 62, 1, 0, 0, 0, 165, 169, 5, 34, 0, 0,
		166, 168, 9, 0, 0, 0, 167, 166, 1, 0, 0, 0, 168, 171, 1, 0, 0, 0, 169,
		170, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 170, 172, 1, 0, 0, 0, 171, 169,
		1, 0, 0, 0, 172, 173, 5, 34, 0, 0, 173, 64, 1, 0, 0, 0, 174, 178, 7, 1,
		0, 0, 175, 177, 7, 2, 0, 0, 176, 175, 1, 0, 0, 0, 177, 180, 1, 0, 0, 0,
		178, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 66, 1, 0, 0, 0, 180, 178,
		1, 0, 0, 0, 181, 185, 7, 3, 0, 0, 182, 184, 7, 2, 0, 0, 183, 182, 1, 0,
		0, 0, 184, 187, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0,
		186, 68, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 188, 190, 7, 4, 0, 0, 189, 188,
		1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 191, 192, 1, 0,
		0, 0, 192, 193, 1, 0, 0, 0, 193, 194, 6, 34, 0, 0, 194, 70, 1, 0, 0, 0,
		195, 196, 5, 123, 0, 0, 196, 197, 5, 45, 0, 0, 197, 201, 1, 0, 0, 0, 198,
		200, 9, 0, 0, 0, 199, 198, 1, 0, 0, 0, 200, 203, 1, 0, 0, 0, 201, 202,
		1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 202, 204, 1, 0, 0, 0, 203, 201, 1, 0,
		0, 0, 204, 205, 5, 45, 0, 0, 205, 206, 5, 125, 0, 0, 206, 207, 1, 0, 0,
		0, 207, 208, 6, 35, 1, 0, 208, 72, 1, 0, 0, 0, 209, 210, 5, 45, 0, 0, 210,
		211, 5, 45, 0, 0, 211, 215, 1, 0, 0, 0, 212, 214, 8, 5, 0, 0, 213, 212,
		1, 0, 0, 0, 214, 217, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 215, 216, 1, 0,
		0, 0, 216, 218, 1, 0, 0, 0, 217, 215, 1, 0, 0, 0, 218, 219, 6, 36, 1, 0,
		219, 74, 1, 0, 0, 0, 8, 0, 159, 169, 178, 185, 191, 201, 215, 2, 6, 0,
		0, 0, 1, 0,
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
	LanguageLexerPARENS       = 14
	LanguageLexerPAREN_LEFT   = 15
	LanguageLexerPAREN_RIGHT  = 16
	LanguageLexerBRACE_LEFT   = 17
	LanguageLexerBRACE_RIGHT  = 18
	LanguageLexerCOMMA        = 19
	LanguageLexerSEMICOLON    = 20
	LanguageLexerVERTICAL_BAR = 21
	LanguageLexerDOUBLE_COLON = 22
	LanguageLexerIMPORT       = 23
	LanguageLexerDATA         = 24
	LanguageLexerDO           = 25
	LanguageLexerCASE         = 26
	LanguageLexerOF           = 27
	LanguageLexerARROW_RIGHT  = 28
	LanguageLexerARROW_LEFT   = 29
	LanguageLexerINT          = 30
	LanguageLexerCHAR         = 31
	LanguageLexerSTRING       = 32
	LanguageLexerVARID        = 33
	LanguageLexerCONID        = 34
	LanguageLexerWHITESPACE   = 35
	LanguageLexerCOMMENT      = 36
	LanguageLexerLINE_COMMENT = 37
)

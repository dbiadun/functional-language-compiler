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
		"", "'*'", "'/'", "'+'", "'-'", "'&'", "'='", "'<'", "'>'", "'<='",
		"'>='", "'=='", "'/='", "'&&'", "'||'", "'()'", "'('", "')'", "'{'",
		"'}'", "','", "';'", "'|'", "'::'", "'import'", "'data'", "'do'", "'case'",
		"'of'", "'->'", "'<-'",
	}
	staticData.symbolicNames = []string{
		"", "MUL", "DIV", "ADD", "SUB", "BIT_AND", "ASSIGN", "LT", "GT", "LTE",
		"GTE", "EQ", "NEQ", "AND", "OR", "PARENS", "PAREN_LEFT", "PAREN_RIGHT",
		"BRACE_LEFT", "BRACE_RIGHT", "COMMA", "SEMICOLON", "VERTICAL_BAR", "DOUBLE_COLON",
		"IMPORT", "DATA", "DO", "CASE", "OF", "ARROW_RIGHT", "ARROW_LEFT", "DEC",
		"HEX", "CHAR", "STRING", "VARID", "CONID", "WHITESPACE", "COMMENT",
		"LINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"MUL", "DIV", "ADD", "SUB", "BIT_AND", "ASSIGN", "LT", "GT", "LTE",
		"GTE", "EQ", "NEQ", "AND", "OR", "PARENS", "PAREN_LEFT", "PAREN_RIGHT",
		"BRACE_LEFT", "BRACE_RIGHT", "COMMA", "SEMICOLON", "VERTICAL_BAR", "DOUBLE_COLON",
		"IMPORT", "DATA", "DO", "CASE", "OF", "ARROW_RIGHT", "ARROW_LEFT", "DEC",
		"HEX", "CHAR", "STRING", "VARID", "CONID", "WHITESPACE", "COMMENT",
		"LINE_COMMENT",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 39, 234, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21,
		1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1,
		29, 1, 29, 1, 30, 4, 30, 164, 8, 30, 11, 30, 12, 30, 165, 1, 31, 1, 31,
		1, 31, 1, 31, 4, 31, 172, 8, 31, 11, 31, 12, 31, 173, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 33, 1, 33, 5, 33, 182, 8, 33, 10, 33, 12, 33, 185, 9, 33,
		1, 33, 1, 33, 1, 34, 1, 34, 5, 34, 191, 8, 34, 10, 34, 12, 34, 194, 9,
		34, 1, 35, 1, 35, 5, 35, 198, 8, 35, 10, 35, 12, 35, 201, 9, 35, 1, 36,
		4, 36, 204, 8, 36, 11, 36, 12, 36, 205, 1, 36, 1, 36, 1, 37, 1, 37, 1,
		37, 1, 37, 5, 37, 214, 8, 37, 10, 37, 12, 37, 217, 9, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 5, 38, 228, 8, 38, 10,
		38, 12, 38, 231, 9, 38, 1, 38, 1, 38, 2, 183, 215, 0, 39, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43,
		22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61,
		31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 1,
		0, 7, 1, 0, 48, 57, 3, 0, 48, 57, 65, 70, 97, 102, 1, 0, 97, 122, 4, 0,
		48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 65, 90, 3, 0, 9, 10, 13, 13, 32,
		32, 2, 0, 10, 10, 13, 13, 241, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1,
		0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59,
		1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0,
		67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0,
		0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 1, 79, 1, 0, 0, 0, 3, 81, 1, 0, 0,
		0, 5, 83, 1, 0, 0, 0, 7, 85, 1, 0, 0, 0, 9, 87, 1, 0, 0, 0, 11, 89, 1,
		0, 0, 0, 13, 91, 1, 0, 0, 0, 15, 93, 1, 0, 0, 0, 17, 95, 1, 0, 0, 0, 19,
		98, 1, 0, 0, 0, 21, 101, 1, 0, 0, 0, 23, 104, 1, 0, 0, 0, 25, 107, 1, 0,
		0, 0, 27, 110, 1, 0, 0, 0, 29, 113, 1, 0, 0, 0, 31, 116, 1, 0, 0, 0, 33,
		118, 1, 0, 0, 0, 35, 120, 1, 0, 0, 0, 37, 122, 1, 0, 0, 0, 39, 124, 1,
		0, 0, 0, 41, 126, 1, 0, 0, 0, 43, 128, 1, 0, 0, 0, 45, 130, 1, 0, 0, 0,
		47, 133, 1, 0, 0, 0, 49, 140, 1, 0, 0, 0, 51, 145, 1, 0, 0, 0, 53, 148,
		1, 0, 0, 0, 55, 153, 1, 0, 0, 0, 57, 156, 1, 0, 0, 0, 59, 159, 1, 0, 0,
		0, 61, 163, 1, 0, 0, 0, 63, 167, 1, 0, 0, 0, 65, 175, 1, 0, 0, 0, 67, 179,
		1, 0, 0, 0, 69, 188, 1, 0, 0, 0, 71, 195, 1, 0, 0, 0, 73, 203, 1, 0, 0,
		0, 75, 209, 1, 0, 0, 0, 77, 223, 1, 0, 0, 0, 79, 80, 5, 42, 0, 0, 80, 2,
		1, 0, 0, 0, 81, 82, 5, 47, 0, 0, 82, 4, 1, 0, 0, 0, 83, 84, 5, 43, 0, 0,
		84, 6, 1, 0, 0, 0, 85, 86, 5, 45, 0, 0, 86, 8, 1, 0, 0, 0, 87, 88, 5, 38,
		0, 0, 88, 10, 1, 0, 0, 0, 89, 90, 5, 61, 0, 0, 90, 12, 1, 0, 0, 0, 91,
		92, 5, 60, 0, 0, 92, 14, 1, 0, 0, 0, 93, 94, 5, 62, 0, 0, 94, 16, 1, 0,
		0, 0, 95, 96, 5, 60, 0, 0, 96, 97, 5, 61, 0, 0, 97, 18, 1, 0, 0, 0, 98,
		99, 5, 62, 0, 0, 99, 100, 5, 61, 0, 0, 100, 20, 1, 0, 0, 0, 101, 102, 5,
		61, 0, 0, 102, 103, 5, 61, 0, 0, 103, 22, 1, 0, 0, 0, 104, 105, 5, 47,
		0, 0, 105, 106, 5, 61, 0, 0, 106, 24, 1, 0, 0, 0, 107, 108, 5, 38, 0, 0,
		108, 109, 5, 38, 0, 0, 109, 26, 1, 0, 0, 0, 110, 111, 5, 124, 0, 0, 111,
		112, 5, 124, 0, 0, 112, 28, 1, 0, 0, 0, 113, 114, 5, 40, 0, 0, 114, 115,
		5, 41, 0, 0, 115, 30, 1, 0, 0, 0, 116, 117, 5, 40, 0, 0, 117, 32, 1, 0,
		0, 0, 118, 119, 5, 41, 0, 0, 119, 34, 1, 0, 0, 0, 120, 121, 5, 123, 0,
		0, 121, 36, 1, 0, 0, 0, 122, 123, 5, 125, 0, 0, 123, 38, 1, 0, 0, 0, 124,
		125, 5, 44, 0, 0, 125, 40, 1, 0, 0, 0, 126, 127, 5, 59, 0, 0, 127, 42,
		1, 0, 0, 0, 128, 129, 5, 124, 0, 0, 129, 44, 1, 0, 0, 0, 130, 131, 5, 58,
		0, 0, 131, 132, 5, 58, 0, 0, 132, 46, 1, 0, 0, 0, 133, 134, 5, 105, 0,
		0, 134, 135, 5, 109, 0, 0, 135, 136, 5, 112, 0, 0, 136, 137, 5, 111, 0,
		0, 137, 138, 5, 114, 0, 0, 138, 139, 5, 116, 0, 0, 139, 48, 1, 0, 0, 0,
		140, 141, 5, 100, 0, 0, 141, 142, 5, 97, 0, 0, 142, 143, 5, 116, 0, 0,
		143, 144, 5, 97, 0, 0, 144, 50, 1, 0, 0, 0, 145, 146, 5, 100, 0, 0, 146,
		147, 5, 111, 0, 0, 147, 52, 1, 0, 0, 0, 148, 149, 5, 99, 0, 0, 149, 150,
		5, 97, 0, 0, 150, 151, 5, 115, 0, 0, 151, 152, 5, 101, 0, 0, 152, 54, 1,
		0, 0, 0, 153, 154, 5, 111, 0, 0, 154, 155, 5, 102, 0, 0, 155, 56, 1, 0,
		0, 0, 156, 157, 5, 45, 0, 0, 157, 158, 5, 62, 0, 0, 158, 58, 1, 0, 0, 0,
		159, 160, 5, 60, 0, 0, 160, 161, 5, 45, 0, 0, 161, 60, 1, 0, 0, 0, 162,
		164, 7, 0, 0, 0, 163, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 163,
		1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 62, 1, 0, 0, 0, 167, 168, 5, 48,
		0, 0, 168, 169, 5, 120, 0, 0, 169, 171, 1, 0, 0, 0, 170, 172, 7, 1, 0,
		0, 171, 170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173,
		174, 1, 0, 0, 0, 174, 64, 1, 0, 0, 0, 175, 176, 5, 39, 0, 0, 176, 177,
		9, 0, 0, 0, 177, 178, 5, 39, 0, 0, 178, 66, 1, 0, 0, 0, 179, 183, 5, 34,
		0, 0, 180, 182, 9, 0, 0, 0, 181, 180, 1, 0, 0, 0, 182, 185, 1, 0, 0, 0,
		183, 184, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 184, 186, 1, 0, 0, 0, 185,
		183, 1, 0, 0, 0, 186, 187, 5, 34, 0, 0, 187, 68, 1, 0, 0, 0, 188, 192,
		7, 2, 0, 0, 189, 191, 7, 3, 0, 0, 190, 189, 1, 0, 0, 0, 191, 194, 1, 0,
		0, 0, 192, 190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 70, 1, 0, 0, 0,
		194, 192, 1, 0, 0, 0, 195, 199, 7, 4, 0, 0, 196, 198, 7, 3, 0, 0, 197,
		196, 1, 0, 0, 0, 198, 201, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 199, 200,
		1, 0, 0, 0, 200, 72, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 202, 204, 7, 5,
		0, 0, 203, 202, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0,
		205, 206, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 208, 6, 36, 0, 0, 208,
		74, 1, 0, 0, 0, 209, 210, 5, 123, 0, 0, 210, 211, 5, 45, 0, 0, 211, 215,
		1, 0, 0, 0, 212, 214, 9, 0, 0, 0, 213, 212, 1, 0, 0, 0, 214, 217, 1, 0,
		0, 0, 215, 216, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 216, 218, 1, 0, 0, 0,
		217, 215, 1, 0, 0, 0, 218, 219, 5, 45, 0, 0, 219, 220, 5, 125, 0, 0, 220,
		221, 1, 0, 0, 0, 221, 222, 6, 37, 1, 0, 222, 76, 1, 0, 0, 0, 223, 224,
		5, 45, 0, 0, 224, 225, 5, 45, 0, 0, 225, 229, 1, 0, 0, 0, 226, 228, 8,
		6, 0, 0, 227, 226, 1, 0, 0, 0, 228, 231, 1, 0, 0, 0, 229, 227, 1, 0, 0,
		0, 229, 230, 1, 0, 0, 0, 230, 232, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 232,
		233, 6, 38, 1, 0, 233, 78, 1, 0, 0, 0, 9, 0, 165, 173, 183, 192, 199, 205,
		215, 229, 2, 6, 0, 0, 0, 1, 0,
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
	LanguageLexerBIT_AND      = 5
	LanguageLexerASSIGN       = 6
	LanguageLexerLT           = 7
	LanguageLexerGT           = 8
	LanguageLexerLTE          = 9
	LanguageLexerGTE          = 10
	LanguageLexerEQ           = 11
	LanguageLexerNEQ          = 12
	LanguageLexerAND          = 13
	LanguageLexerOR           = 14
	LanguageLexerPARENS       = 15
	LanguageLexerPAREN_LEFT   = 16
	LanguageLexerPAREN_RIGHT  = 17
	LanguageLexerBRACE_LEFT   = 18
	LanguageLexerBRACE_RIGHT  = 19
	LanguageLexerCOMMA        = 20
	LanguageLexerSEMICOLON    = 21
	LanguageLexerVERTICAL_BAR = 22
	LanguageLexerDOUBLE_COLON = 23
	LanguageLexerIMPORT       = 24
	LanguageLexerDATA         = 25
	LanguageLexerDO           = 26
	LanguageLexerCASE         = 27
	LanguageLexerOF           = 28
	LanguageLexerARROW_RIGHT  = 29
	LanguageLexerARROW_LEFT   = 30
	LanguageLexerDEC          = 31
	LanguageLexerHEX          = 32
	LanguageLexerCHAR         = 33
	LanguageLexerSTRING       = 34
	LanguageLexerVARID        = 35
	LanguageLexerCONID        = 36
	LanguageLexerWHITESPACE   = 37
	LanguageLexerCOMMENT      = 38
	LanguageLexerLINE_COMMENT = 39
)

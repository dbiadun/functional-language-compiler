// Code generated from Language.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Language

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type LanguageParser struct {
	*antlr.BaseParser
}

var languageParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func languageParserInit() {
	staticData := &languageParserStaticData
	staticData.literalNames = []string{
		"", "'*'", "'/'", "'+'", "'-'", "'='", "'('", "')'", "'{'", "'}'", "','",
		"';'", "'|'", "'::'", "'data'", "'case'", "'of'", "'->'",
	}
	staticData.symbolicNames = []string{
		"", "MUL", "DIV", "ADD", "SUB", "ASSIGN", "PAREN_LEFT", "PAREN_RIGHT",
		"BRACE_LEFT", "BRACE_RIGHT", "COMMA", "SEMICOLON", "VERTICAL_BAR", "DOUBLE_COLON",
		"DATA", "CASE", "OF", "ARROW", "INT", "CHAR", "STRING", "VARID", "CONID",
		"WHITESPACE",
	}
	staticData.ruleNames = []string{
		"start", "topdecls", "topdecl", "simpletype", "constrs", "constrdef",
		"decl", "gendecl", "vars", "type", "btype", "atype", "funlhs", "rhs",
		"exp", "fexp", "aexp", "alts", "alt", "pat", "apat", "literal",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 23, 237, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 1, 0, 1, 0, 1, 0, 1, 1, 3, 1, 49, 8, 1, 1, 1, 1, 1, 5, 1, 53,
		8, 1, 10, 1, 12, 1, 56, 9, 1, 1, 1, 3, 1, 59, 8, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 3, 2, 67, 8, 2, 1, 3, 1, 3, 5, 3, 71, 8, 3, 10, 3, 12, 3,
		74, 9, 3, 1, 4, 1, 4, 1, 4, 5, 4, 79, 8, 4, 10, 4, 12, 4, 82, 9, 4, 1,
		5, 1, 5, 5, 5, 86, 8, 5, 10, 5, 12, 5, 89, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 3, 6, 98, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 5, 8, 107, 8, 8, 10, 8, 12, 8, 110, 9, 8, 1, 9, 1, 9, 1, 9, 5, 9, 115,
		8, 9, 10, 9, 12, 9, 118, 9, 9, 1, 10, 4, 10, 121, 8, 10, 11, 10, 12, 10,
		122, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 4, 11, 131, 8, 11, 11, 11,
		12, 11, 132, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 141, 8, 11,
		1, 12, 1, 12, 4, 12, 145, 8, 12, 11, 12, 12, 12, 146, 1, 13, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14,
		161, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 169, 8, 14,
		10, 14, 12, 14, 172, 9, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 179,
		8, 15, 10, 15, 12, 15, 182, 9, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 4, 16, 195, 8, 16, 11, 16, 12, 16,
		196, 1, 16, 1, 16, 3, 16, 201, 8, 16, 1, 17, 1, 17, 1, 17, 5, 17, 206,
		8, 17, 10, 17, 12, 17, 209, 9, 17, 1, 17, 3, 17, 212, 8, 17, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 4, 19, 221, 8, 19, 11, 19, 12, 19,
		222, 3, 19, 225, 8, 19, 1, 20, 1, 20, 1, 20, 3, 20, 230, 8, 20, 1, 21,
		1, 21, 1, 21, 3, 21, 235, 8, 21, 1, 21, 0, 2, 28, 30, 22, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 0,
		2, 1, 0, 1, 2, 1, 0, 3, 4, 248, 0, 44, 1, 0, 0, 0, 2, 48, 1, 0, 0, 0, 4,
		66, 1, 0, 0, 0, 6, 68, 1, 0, 0, 0, 8, 75, 1, 0, 0, 0, 10, 83, 1, 0, 0,
		0, 12, 97, 1, 0, 0, 0, 14, 99, 1, 0, 0, 0, 16, 103, 1, 0, 0, 0, 18, 111,
		1, 0, 0, 0, 20, 120, 1, 0, 0, 0, 22, 140, 1, 0, 0, 0, 24, 142, 1, 0, 0,
		0, 26, 148, 1, 0, 0, 0, 28, 160, 1, 0, 0, 0, 30, 173, 1, 0, 0, 0, 32, 200,
		1, 0, 0, 0, 34, 202, 1, 0, 0, 0, 36, 213, 1, 0, 0, 0, 38, 224, 1, 0, 0,
		0, 40, 229, 1, 0, 0, 0, 42, 234, 1, 0, 0, 0, 44, 45, 3, 2, 1, 0, 45, 46,
		5, 0, 0, 1, 46, 1, 1, 0, 0, 0, 47, 49, 3, 4, 2, 0, 48, 47, 1, 0, 0, 0,
		48, 49, 1, 0, 0, 0, 49, 54, 1, 0, 0, 0, 50, 51, 5, 11, 0, 0, 51, 53, 3,
		4, 2, 0, 52, 50, 1, 0, 0, 0, 53, 56, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 54,
		55, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 57, 59, 5, 11,
		0, 0, 58, 57, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 3, 1, 0, 0, 0, 60, 61,
		5, 14, 0, 0, 61, 62, 3, 6, 3, 0, 62, 63, 5, 5, 0, 0, 63, 64, 3, 8, 4, 0,
		64, 67, 1, 0, 0, 0, 65, 67, 3, 12, 6, 0, 66, 60, 1, 0, 0, 0, 66, 65, 1,
		0, 0, 0, 67, 5, 1, 0, 0, 0, 68, 72, 5, 22, 0, 0, 69, 71, 5, 21, 0, 0, 70,
		69, 1, 0, 0, 0, 71, 74, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 72, 73, 1, 0, 0,
		0, 73, 7, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 75, 80, 3, 10, 5, 0, 76, 77,
		5, 12, 0, 0, 77, 79, 3, 10, 5, 0, 78, 76, 1, 0, 0, 0, 79, 82, 1, 0, 0,
		0, 80, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 9, 1, 0, 0, 0, 82, 80, 1,
		0, 0, 0, 83, 87, 5, 22, 0, 0, 84, 86, 3, 22, 11, 0, 85, 84, 1, 0, 0, 0,
		86, 89, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 11, 1,
		0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 98, 3, 14, 7, 0, 91, 92, 3, 24, 12, 0,
		92, 93, 3, 26, 13, 0, 93, 98, 1, 0, 0, 0, 94, 95, 3, 38, 19, 0, 95, 96,
		3, 26, 13, 0, 96, 98, 1, 0, 0, 0, 97, 90, 1, 0, 0, 0, 97, 91, 1, 0, 0,
		0, 97, 94, 1, 0, 0, 0, 98, 13, 1, 0, 0, 0, 99, 100, 3, 16, 8, 0, 100, 101,
		5, 13, 0, 0, 101, 102, 3, 18, 9, 0, 102, 15, 1, 0, 0, 0, 103, 108, 5, 21,
		0, 0, 104, 105, 5, 10, 0, 0, 105, 107, 5, 21, 0, 0, 106, 104, 1, 0, 0,
		0, 107, 110, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109,
		17, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 111, 116, 3, 20, 10, 0, 112, 113,
		5, 17, 0, 0, 113, 115, 3, 20, 10, 0, 114, 112, 1, 0, 0, 0, 115, 118, 1,
		0, 0, 0, 116, 114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 19, 1, 0, 0,
		0, 118, 116, 1, 0, 0, 0, 119, 121, 3, 22, 11, 0, 120, 119, 1, 0, 0, 0,
		121, 122, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123,
		21, 1, 0, 0, 0, 124, 141, 5, 22, 0, 0, 125, 141, 5, 21, 0, 0, 126, 127,
		5, 6, 0, 0, 127, 130, 3, 18, 9, 0, 128, 129, 5, 10, 0, 0, 129, 131, 3,
		18, 9, 0, 130, 128, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 130, 1, 0, 0,
		0, 132, 133, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 135, 5, 7, 0, 0, 135,
		141, 1, 0, 0, 0, 136, 137, 5, 6, 0, 0, 137, 138, 3, 18, 9, 0, 138, 139,
		5, 7, 0, 0, 139, 141, 1, 0, 0, 0, 140, 124, 1, 0, 0, 0, 140, 125, 1, 0,
		0, 0, 140, 126, 1, 0, 0, 0, 140, 136, 1, 0, 0, 0, 141, 23, 1, 0, 0, 0,
		142, 144, 5, 21, 0, 0, 143, 145, 3, 40, 20, 0, 144, 143, 1, 0, 0, 0, 145,
		146, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 25, 1,
		0, 0, 0, 148, 149, 5, 5, 0, 0, 149, 150, 3, 28, 14, 0, 150, 27, 1, 0, 0,
		0, 151, 152, 6, 14, -1, 0, 152, 161, 3, 30, 15, 0, 153, 154, 5, 15, 0,
		0, 154, 155, 3, 28, 14, 0, 155, 156, 5, 16, 0, 0, 156, 157, 5, 8, 0, 0,
		157, 158, 3, 34, 17, 0, 158, 159, 5, 9, 0, 0, 159, 161, 1, 0, 0, 0, 160,
		151, 1, 0, 0, 0, 160, 153, 1, 0, 0, 0, 161, 170, 1, 0, 0, 0, 162, 163,
		10, 2, 0, 0, 163, 164, 7, 0, 0, 0, 164, 169, 3, 28, 14, 3, 165, 166, 10,
		1, 0, 0, 166, 167, 7, 1, 0, 0, 167, 169, 3, 28, 14, 2, 168, 162, 1, 0,
		0, 0, 168, 165, 1, 0, 0, 0, 169, 172, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0,
		170, 171, 1, 0, 0, 0, 171, 29, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 173, 174,
		6, 15, -1, 0, 174, 175, 3, 32, 16, 0, 175, 180, 1, 0, 0, 0, 176, 177, 10,
		2, 0, 0, 177, 179, 3, 32, 16, 0, 178, 176, 1, 0, 0, 0, 179, 182, 1, 0,
		0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 31, 1, 0, 0, 0,
		182, 180, 1, 0, 0, 0, 183, 201, 5, 21, 0, 0, 184, 201, 5, 22, 0, 0, 185,
		201, 3, 42, 21, 0, 186, 187, 5, 6, 0, 0, 187, 188, 3, 28, 14, 0, 188, 189,
		5, 7, 0, 0, 189, 201, 1, 0, 0, 0, 190, 191, 5, 6, 0, 0, 191, 194, 3, 28,
		14, 0, 192, 193, 5, 10, 0, 0, 193, 195, 3, 28, 14, 0, 194, 192, 1, 0, 0,
		0, 195, 196, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197,
		198, 1, 0, 0, 0, 198, 199, 5, 7, 0, 0, 199, 201, 1, 0, 0, 0, 200, 183,
		1, 0, 0, 0, 200, 184, 1, 0, 0, 0, 200, 185, 1, 0, 0, 0, 200, 186, 1, 0,
		0, 0, 200, 190, 1, 0, 0, 0, 201, 33, 1, 0, 0, 0, 202, 207, 3, 36, 18, 0,
		203, 204, 5, 11, 0, 0, 204, 206, 3, 36, 18, 0, 205, 203, 1, 0, 0, 0, 206,
		209, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 211,
		1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 210, 212, 5, 11, 0, 0, 211, 210, 1, 0,
		0, 0, 211, 212, 1, 0, 0, 0, 212, 35, 1, 0, 0, 0, 213, 214, 3, 38, 19, 0,
		214, 215, 5, 17, 0, 0, 215, 216, 3, 28, 14, 0, 216, 37, 1, 0, 0, 0, 217,
		225, 3, 40, 20, 0, 218, 220, 5, 22, 0, 0, 219, 221, 3, 40, 20, 0, 220,
		219, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 222, 223,
		1, 0, 0, 0, 223, 225, 1, 0, 0, 0, 224, 217, 1, 0, 0, 0, 224, 218, 1, 0,
		0, 0, 225, 39, 1, 0, 0, 0, 226, 230, 5, 21, 0, 0, 227, 230, 5, 22, 0, 0,
		228, 230, 3, 42, 21, 0, 229, 226, 1, 0, 0, 0, 229, 227, 1, 0, 0, 0, 229,
		228, 1, 0, 0, 0, 230, 41, 1, 0, 0, 0, 231, 235, 5, 18, 0, 0, 232, 235,
		5, 19, 0, 0, 233, 235, 5, 20, 0, 0, 234, 231, 1, 0, 0, 0, 234, 232, 1,
		0, 0, 0, 234, 233, 1, 0, 0, 0, 235, 43, 1, 0, 0, 0, 26, 48, 54, 58, 66,
		72, 80, 87, 97, 108, 116, 122, 132, 140, 146, 160, 168, 170, 180, 196,
		200, 207, 211, 222, 224, 229, 234,
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

// LanguageParserInit initializes any static state used to implement LanguageParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLanguageParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LanguageParserInit() {
	staticData := &languageParserStaticData
	staticData.once.Do(languageParserInit)
}

// NewLanguageParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLanguageParser(input antlr.TokenStream) *LanguageParser {
	LanguageParserInit()
	this := new(LanguageParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &languageParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Language.g4"

	return this
}

// LanguageParser tokens.
const (
	LanguageParserEOF          = antlr.TokenEOF
	LanguageParserMUL          = 1
	LanguageParserDIV          = 2
	LanguageParserADD          = 3
	LanguageParserSUB          = 4
	LanguageParserASSIGN       = 5
	LanguageParserPAREN_LEFT   = 6
	LanguageParserPAREN_RIGHT  = 7
	LanguageParserBRACE_LEFT   = 8
	LanguageParserBRACE_RIGHT  = 9
	LanguageParserCOMMA        = 10
	LanguageParserSEMICOLON    = 11
	LanguageParserVERTICAL_BAR = 12
	LanguageParserDOUBLE_COLON = 13
	LanguageParserDATA         = 14
	LanguageParserCASE         = 15
	LanguageParserOF           = 16
	LanguageParserARROW        = 17
	LanguageParserINT          = 18
	LanguageParserCHAR         = 19
	LanguageParserSTRING       = 20
	LanguageParserVARID        = 21
	LanguageParserCONID        = 22
	LanguageParserWHITESPACE   = 23
)

// LanguageParser rules.
const (
	LanguageParserRULE_start      = 0
	LanguageParserRULE_topdecls   = 1
	LanguageParserRULE_topdecl    = 2
	LanguageParserRULE_simpletype = 3
	LanguageParserRULE_constrs    = 4
	LanguageParserRULE_constrdef  = 5
	LanguageParserRULE_decl       = 6
	LanguageParserRULE_gendecl    = 7
	LanguageParserRULE_vars       = 8
	LanguageParserRULE_type       = 9
	LanguageParserRULE_btype      = 10
	LanguageParserRULE_atype      = 11
	LanguageParserRULE_funlhs     = 12
	LanguageParserRULE_rhs        = 13
	LanguageParserRULE_exp        = 14
	LanguageParserRULE_fexp       = 15
	LanguageParserRULE_aexp       = 16
	LanguageParserRULE_alts       = 17
	LanguageParserRULE_alt        = 18
	LanguageParserRULE_pat        = 19
	LanguageParserRULE_apat       = 20
	LanguageParserRULE_literal    = 21
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Topdecls() ITopdeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopdeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITopdeclsContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(LanguageParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LanguageParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Topdecls()
	}
	{
		p.SetState(45)
		p.Match(LanguageParserEOF)
	}

	return localctx
}

// ITopdeclsContext is an interface to support dynamic dispatch.
type ITopdeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopdeclsContext differentiates from other interfaces.
	IsTopdeclsContext()
}

type TopdeclsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopdeclsContext() *TopdeclsContext {
	var p = new(TopdeclsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_topdecls
	return p
}

func (*TopdeclsContext) IsTopdeclsContext() {}

func NewTopdeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopdeclsContext {
	var p = new(TopdeclsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_topdecls

	return p
}

func (s *TopdeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *TopdeclsContext) CopyFrom(ctx *TopdeclsContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TopdeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopdeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TopDeclsListContext struct {
	*TopdeclsContext
}

func NewTopDeclsListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TopDeclsListContext {
	var p = new(TopDeclsListContext)

	p.TopdeclsContext = NewEmptyTopdeclsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TopdeclsContext))

	return p
}

func (s *TopDeclsListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopDeclsListContext) AllTopdecl() []ITopdeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITopdeclContext); ok {
			len++
		}
	}

	tst := make([]ITopdeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITopdeclContext); ok {
			tst[i] = t.(ITopdeclContext)
			i++
		}
	}

	return tst
}

func (s *TopDeclsListContext) Topdecl(i int) ITopdeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopdeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITopdeclContext)
}

func (s *TopDeclsListContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(LanguageParserSEMICOLON)
}

func (s *TopDeclsListContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(LanguageParserSEMICOLON, i)
}

func (s *TopDeclsListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterTopDeclsList(s)
	}
}

func (s *TopDeclsListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitTopDeclsList(s)
	}
}

func (s *TopDeclsListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitTopDeclsList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Topdecls() (localctx ITopdeclsContext) {
	this := p
	_ = this

	localctx = NewTopdeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LanguageParserRULE_topdecls)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	localctx = NewTopDeclsListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LanguageParserDATA)|(1<<LanguageParserINT)|(1<<LanguageParserCHAR)|(1<<LanguageParserSTRING)|(1<<LanguageParserVARID)|(1<<LanguageParserCONID))) != 0 {
		{
			p.SetState(47)
			p.Topdecl()
		}

	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(50)
				p.Match(LanguageParserSEMICOLON)
			}
			{
				p.SetState(51)
				p.Topdecl()
			}

		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LanguageParserSEMICOLON {
		{
			p.SetState(57)
			p.Match(LanguageParserSEMICOLON)
		}

	}

	return localctx
}

// ITopdeclContext is an interface to support dynamic dispatch.
type ITopdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopdeclContext differentiates from other interfaces.
	IsTopdeclContext()
}

type TopdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopdeclContext() *TopdeclContext {
	var p = new(TopdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_topdecl
	return p
}

func (*TopdeclContext) IsTopdeclContext() {}

func NewTopdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopdeclContext {
	var p = new(TopdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_topdecl

	return p
}

func (s *TopdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TopdeclContext) CopyFrom(ctx *TopdeclContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TopdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DataTopDeclContext struct {
	*TopdeclContext
}

func NewDataTopDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DataTopDeclContext {
	var p = new(DataTopDeclContext)

	p.TopdeclContext = NewEmptyTopdeclContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TopdeclContext))

	return p
}

func (s *DataTopDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataTopDeclContext) DATA() antlr.TerminalNode {
	return s.GetToken(LanguageParserDATA, 0)
}

func (s *DataTopDeclContext) Simpletype() ISimpletypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpletypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpletypeContext)
}

func (s *DataTopDeclContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(LanguageParserASSIGN, 0)
}

func (s *DataTopDeclContext) Constrs() IConstrsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstrsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstrsContext)
}

func (s *DataTopDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterDataTopDecl(s)
	}
}

func (s *DataTopDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitDataTopDecl(s)
	}
}

func (s *DataTopDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitDataTopDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunTopDeclContext struct {
	*TopdeclContext
}

func NewFunTopDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunTopDeclContext {
	var p = new(FunTopDeclContext)

	p.TopdeclContext = NewEmptyTopdeclContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TopdeclContext))

	return p
}

func (s *FunTopDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunTopDeclContext) Decl() IDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *FunTopDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterFunTopDecl(s)
	}
}

func (s *FunTopDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitFunTopDecl(s)
	}
}

func (s *FunTopDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitFunTopDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Topdecl() (localctx ITopdeclContext) {
	this := p
	_ = this

	localctx = NewTopdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LanguageParserRULE_topdecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(66)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LanguageParserDATA:
		localctx = NewDataTopDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			p.Match(LanguageParserDATA)
		}
		{
			p.SetState(61)
			p.Simpletype()
		}
		{
			p.SetState(62)
			p.Match(LanguageParserASSIGN)
		}
		{
			p.SetState(63)
			p.Constrs()
		}

	case LanguageParserINT, LanguageParserCHAR, LanguageParserSTRING, LanguageParserVARID, LanguageParserCONID:
		localctx = NewFunTopDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Decl()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISimpletypeContext is an interface to support dynamic dispatch.
type ISimpletypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimpletypeContext differentiates from other interfaces.
	IsSimpletypeContext()
}

type SimpletypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpletypeContext() *SimpletypeContext {
	var p = new(SimpletypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_simpletype
	return p
}

func (*SimpletypeContext) IsSimpletypeContext() {}

func NewSimpletypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpletypeContext {
	var p = new(SimpletypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_simpletype

	return p
}

func (s *SimpletypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpletypeContext) CopyFrom(ctx *SimpletypeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SimpletypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpletypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DataTypeContext struct {
	*SimpletypeContext
}

func NewDataTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DataTypeContext {
	var p = new(DataTypeContext)

	p.SimpletypeContext = NewEmptySimpletypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SimpletypeContext))

	return p
}

func (s *DataTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataTypeContext) CONID() antlr.TerminalNode {
	return s.GetToken(LanguageParserCONID, 0)
}

func (s *DataTypeContext) AllVARID() []antlr.TerminalNode {
	return s.GetTokens(LanguageParserVARID)
}

func (s *DataTypeContext) VARID(i int) antlr.TerminalNode {
	return s.GetToken(LanguageParserVARID, i)
}

func (s *DataTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterDataType(s)
	}
}

func (s *DataTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitDataType(s)
	}
}

func (s *DataTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitDataType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Simpletype() (localctx ISimpletypeContext) {
	this := p
	_ = this

	localctx = NewSimpletypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LanguageParserRULE_simpletype)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewDataTypeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(LanguageParserCONID)
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LanguageParserVARID {
		{
			p.SetState(69)
			p.Match(LanguageParserVARID)
		}

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IConstrsContext is an interface to support dynamic dispatch.
type IConstrsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstrsContext differentiates from other interfaces.
	IsConstrsContext()
}

type ConstrsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstrsContext() *ConstrsContext {
	var p = new(ConstrsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_constrs
	return p
}

func (*ConstrsContext) IsConstrsContext() {}

func NewConstrsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstrsContext {
	var p = new(ConstrsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_constrs

	return p
}

func (s *ConstrsContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstrsContext) CopyFrom(ctx *ConstrsContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ConstrsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstrsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ConstrListContext struct {
	*ConstrsContext
}

func NewConstrListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstrListContext {
	var p = new(ConstrListContext)

	p.ConstrsContext = NewEmptyConstrsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstrsContext))

	return p
}

func (s *ConstrListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstrListContext) AllConstrdef() []IConstrdefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstrdefContext); ok {
			len++
		}
	}

	tst := make([]IConstrdefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstrdefContext); ok {
			tst[i] = t.(IConstrdefContext)
			i++
		}
	}

	return tst
}

func (s *ConstrListContext) Constrdef(i int) IConstrdefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstrdefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstrdefContext)
}

func (s *ConstrListContext) AllVERTICAL_BAR() []antlr.TerminalNode {
	return s.GetTokens(LanguageParserVERTICAL_BAR)
}

func (s *ConstrListContext) VERTICAL_BAR(i int) antlr.TerminalNode {
	return s.GetToken(LanguageParserVERTICAL_BAR, i)
}

func (s *ConstrListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterConstrList(s)
	}
}

func (s *ConstrListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitConstrList(s)
	}
}

func (s *ConstrListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitConstrList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Constrs() (localctx IConstrsContext) {
	this := p
	_ = this

	localctx = NewConstrsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LanguageParserRULE_constrs)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewConstrListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Constrdef()
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LanguageParserVERTICAL_BAR {
		{
			p.SetState(76)
			p.Match(LanguageParserVERTICAL_BAR)
		}
		{
			p.SetState(77)
			p.Constrdef()
		}

		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IConstrdefContext is an interface to support dynamic dispatch.
type IConstrdefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstrdefContext differentiates from other interfaces.
	IsConstrdefContext()
}

type ConstrdefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstrdefContext() *ConstrdefContext {
	var p = new(ConstrdefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_constrdef
	return p
}

func (*ConstrdefContext) IsConstrdefContext() {}

func NewConstrdefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstrdefContext {
	var p = new(ConstrdefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_constrdef

	return p
}

func (s *ConstrdefContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstrdefContext) CopyFrom(ctx *ConstrdefContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ConstrdefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstrdefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ConstrTypeContext struct {
	*ConstrdefContext
}

func NewConstrTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstrTypeContext {
	var p = new(ConstrTypeContext)

	p.ConstrdefContext = NewEmptyConstrdefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstrdefContext))

	return p
}

func (s *ConstrTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstrTypeContext) CONID() antlr.TerminalNode {
	return s.GetToken(LanguageParserCONID, 0)
}

func (s *ConstrTypeContext) AllAtype() []IAtypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAtypeContext); ok {
			len++
		}
	}

	tst := make([]IAtypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAtypeContext); ok {
			tst[i] = t.(IAtypeContext)
			i++
		}
	}

	return tst
}

func (s *ConstrTypeContext) Atype(i int) IAtypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtypeContext)
}

func (s *ConstrTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterConstrType(s)
	}
}

func (s *ConstrTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitConstrType(s)
	}
}

func (s *ConstrTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitConstrType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Constrdef() (localctx IConstrdefContext) {
	this := p
	_ = this

	localctx = NewConstrdefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LanguageParserRULE_constrdef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewConstrTypeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Match(LanguageParserCONID)
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LanguageParserPAREN_LEFT)|(1<<LanguageParserVARID)|(1<<LanguageParserCONID))) != 0 {
		{
			p.SetState(84)
			p.Atype()
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDeclContext is an interface to support dynamic dispatch.
type IDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclContext differentiates from other interfaces.
	IsDeclContext()
}

type DeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclContext() *DeclContext {
	var p = new(DeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_decl
	return p
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclContext) CopyFrom(ctx *DeclContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunDeclContext struct {
	*DeclContext
}

func NewFunDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunDeclContext {
	var p = new(FunDeclContext)

	p.DeclContext = NewEmptyDeclContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclContext))

	return p
}

func (s *FunDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunDeclContext) Funlhs() IFunlhsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunlhsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunlhsContext)
}

func (s *FunDeclContext) Rhs() IRhsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRhsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRhsContext)
}

func (s *FunDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterFunDecl(s)
	}
}

func (s *FunDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitFunDecl(s)
	}
}

func (s *FunDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitFunDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarDeclContext struct {
	*DeclContext
}

func NewVarDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarDeclContext {
	var p = new(VarDeclContext)

	p.DeclContext = NewEmptyDeclContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclContext))

	return p
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) Pat() IPatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatContext)
}

func (s *VarDeclContext) Rhs() IRhsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRhsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRhsContext)
}

func (s *VarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterVarDecl(s)
	}
}

func (s *VarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitVarDecl(s)
	}
}

func (s *VarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunTypeDeclContext struct {
	*DeclContext
}

func NewFunTypeDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunTypeDeclContext {
	var p = new(FunTypeDeclContext)

	p.DeclContext = NewEmptyDeclContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclContext))

	return p
}

func (s *FunTypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunTypeDeclContext) Gendecl() IGendeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGendeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGendeclContext)
}

func (s *FunTypeDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterFunTypeDecl(s)
	}
}

func (s *FunTypeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitFunTypeDecl(s)
	}
}

func (s *FunTypeDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitFunTypeDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Decl() (localctx IDeclContext) {
	this := p
	_ = this

	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LanguageParserRULE_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunTypeDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)
			p.Gendecl()
		}

	case 2:
		localctx = NewFunDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
			p.Funlhs()
		}
		{
			p.SetState(92)
			p.Rhs()
		}

	case 3:
		localctx = NewVarDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(94)
			p.Pat()
		}
		{
			p.SetState(95)
			p.Rhs()
		}

	}

	return localctx
}

// IGendeclContext is an interface to support dynamic dispatch.
type IGendeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGendeclContext differentiates from other interfaces.
	IsGendeclContext()
}

type GendeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGendeclContext() *GendeclContext {
	var p = new(GendeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_gendecl
	return p
}

func (*GendeclContext) IsGendeclContext() {}

func NewGendeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GendeclContext {
	var p = new(GendeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_gendecl

	return p
}

func (s *GendeclContext) GetParser() antlr.Parser { return s.parser }

func (s *GendeclContext) CopyFrom(ctx *GendeclContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *GendeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GendeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeSignatureContext struct {
	*GendeclContext
}

func NewTypeSignatureContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeSignatureContext {
	var p = new(TypeSignatureContext)

	p.GendeclContext = NewEmptyGendeclContext()
	p.parser = parser
	p.CopyFrom(ctx.(*GendeclContext))

	return p
}

func (s *TypeSignatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSignatureContext) Vars() IVarsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarsContext)
}

func (s *TypeSignatureContext) DOUBLE_COLON() antlr.TerminalNode {
	return s.GetToken(LanguageParserDOUBLE_COLON, 0)
}

func (s *TypeSignatureContext) Type() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypeSignatureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterTypeSignature(s)
	}
}

func (s *TypeSignatureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitTypeSignature(s)
	}
}

func (s *TypeSignatureContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitTypeSignature(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Gendecl() (localctx IGendeclContext) {
	this := p
	_ = this

	localctx = NewGendeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LanguageParserRULE_gendecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewTypeSignatureContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Vars()
	}
	{
		p.SetState(100)
		p.Match(LanguageParserDOUBLE_COLON)
	}
	{
		p.SetState(101)
		p.Type()
	}

	return localctx
}

// IVarsContext is an interface to support dynamic dispatch.
type IVarsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarsContext differentiates from other interfaces.
	IsVarsContext()
}

type VarsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarsContext() *VarsContext {
	var p = new(VarsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_vars
	return p
}

func (*VarsContext) IsVarsContext() {}

func NewVarsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsContext {
	var p = new(VarsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_vars

	return p
}

func (s *VarsContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsContext) CopyFrom(ctx *VarsContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *VarsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VarListContext struct {
	*VarsContext
}

func NewVarListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarListContext {
	var p = new(VarListContext)

	p.VarsContext = NewEmptyVarsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*VarsContext))

	return p
}

func (s *VarListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarListContext) AllVARID() []antlr.TerminalNode {
	return s.GetTokens(LanguageParserVARID)
}

func (s *VarListContext) VARID(i int) antlr.TerminalNode {
	return s.GetToken(LanguageParserVARID, i)
}

func (s *VarListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LanguageParserCOMMA)
}

func (s *VarListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LanguageParserCOMMA, i)
}

func (s *VarListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterVarList(s)
	}
}

func (s *VarListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitVarList(s)
	}
}

func (s *VarListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitVarList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Vars() (localctx IVarsContext) {
	this := p
	_ = this

	localctx = NewVarsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LanguageParserRULE_vars)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewVarListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(LanguageParserVARID)
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LanguageParserCOMMA {
		{
			p.SetState(104)
			p.Match(LanguageParserCOMMA)
		}
		{
			p.SetState(105)
			p.Match(LanguageParserVARID)
		}

		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_type
	return p
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) CopyFrom(ctx *TypeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunTypeContext struct {
	*TypeContext
}

func NewFunTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunTypeContext {
	var p = new(FunTypeContext)

	p.TypeContext = NewEmptyTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeContext))

	return p
}

func (s *FunTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunTypeContext) AllBtype() []IBtypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBtypeContext); ok {
			len++
		}
	}

	tst := make([]IBtypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBtypeContext); ok {
			tst[i] = t.(IBtypeContext)
			i++
		}
	}

	return tst
}

func (s *FunTypeContext) Btype(i int) IBtypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBtypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBtypeContext)
}

func (s *FunTypeContext) AllARROW() []antlr.TerminalNode {
	return s.GetTokens(LanguageParserARROW)
}

func (s *FunTypeContext) ARROW(i int) antlr.TerminalNode {
	return s.GetToken(LanguageParserARROW, i)
}

func (s *FunTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterFunType(s)
	}
}

func (s *FunTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitFunType(s)
	}
}

func (s *FunTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitFunType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Type() (localctx ITypeContext) {
	this := p
	_ = this

	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LanguageParserRULE_type)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewFunTypeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Btype()
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LanguageParserARROW {
		{
			p.SetState(112)
			p.Match(LanguageParserARROW)
		}
		{
			p.SetState(113)
			p.Btype()
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBtypeContext is an interface to support dynamic dispatch.
type IBtypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBtypeContext differentiates from other interfaces.
	IsBtypeContext()
}

type BtypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBtypeContext() *BtypeContext {
	var p = new(BtypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_btype
	return p
}

func (*BtypeContext) IsBtypeContext() {}

func NewBtypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BtypeContext {
	var p = new(BtypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_btype

	return p
}

func (s *BtypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BtypeContext) CopyFrom(ctx *BtypeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BtypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeAppContext struct {
	*BtypeContext
}

func NewTypeAppContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeAppContext {
	var p = new(TypeAppContext)

	p.BtypeContext = NewEmptyBtypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BtypeContext))

	return p
}

func (s *TypeAppContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeAppContext) AllAtype() []IAtypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAtypeContext); ok {
			len++
		}
	}

	tst := make([]IAtypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAtypeContext); ok {
			tst[i] = t.(IAtypeContext)
			i++
		}
	}

	return tst
}

func (s *TypeAppContext) Atype(i int) IAtypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtypeContext)
}

func (s *TypeAppContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterTypeApp(s)
	}
}

func (s *TypeAppContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitTypeApp(s)
	}
}

func (s *TypeAppContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitTypeApp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Btype() (localctx IBtypeContext) {
	this := p
	_ = this

	localctx = NewBtypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LanguageParserRULE_btype)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewTypeAppContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LanguageParserPAREN_LEFT)|(1<<LanguageParserVARID)|(1<<LanguageParserCONID))) != 0) {
		{
			p.SetState(119)
			p.Atype()
		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAtypeContext is an interface to support dynamic dispatch.
type IAtypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtypeContext differentiates from other interfaces.
	IsAtypeContext()
}

type AtypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtypeContext() *AtypeContext {
	var p = new(AtypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_atype
	return p
}

func (*AtypeContext) IsAtypeContext() {}

func NewAtypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtypeContext {
	var p = new(AtypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_atype

	return p
}

func (s *AtypeContext) GetParser() antlr.Parser { return s.parser }

func (s *AtypeContext) CopyFrom(ctx *AtypeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TupleTypeContext struct {
	*AtypeContext
}

func NewTupleTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleTypeContext {
	var p = new(TupleTypeContext)

	p.AtypeContext = NewEmptyAtypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtypeContext))

	return p
}

func (s *TupleTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleTypeContext) PAREN_LEFT() antlr.TerminalNode {
	return s.GetToken(LanguageParserPAREN_LEFT, 0)
}

func (s *TupleTypeContext) AllType() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *TupleTypeContext) Type(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TupleTypeContext) PAREN_RIGHT() antlr.TerminalNode {
	return s.GetToken(LanguageParserPAREN_RIGHT, 0)
}

func (s *TupleTypeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LanguageParserCOMMA)
}

func (s *TupleTypeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LanguageParserCOMMA, i)
}

func (s *TupleTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterTupleType(s)
	}
}

func (s *TupleTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitTupleType(s)
	}
}

func (s *TupleTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitTupleType(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConTypeContext struct {
	*AtypeContext
}

func NewConTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConTypeContext {
	var p = new(ConTypeContext)

	p.AtypeContext = NewEmptyAtypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtypeContext))

	return p
}

func (s *ConTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConTypeContext) CONID() antlr.TerminalNode {
	return s.GetToken(LanguageParserCONID, 0)
}

func (s *ConTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterConType(s)
	}
}

func (s *ConTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitConType(s)
	}
}

func (s *ConTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitConType(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarTypeContext struct {
	*AtypeContext
}

func NewVarTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarTypeContext {
	var p = new(VarTypeContext)

	p.AtypeContext = NewEmptyAtypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtypeContext))

	return p
}

func (s *VarTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarTypeContext) VARID() antlr.TerminalNode {
	return s.GetToken(LanguageParserVARID, 0)
}

func (s *VarTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterVarType(s)
	}
}

func (s *VarTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitVarType(s)
	}
}

func (s *VarTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitVarType(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenTypeContext struct {
	*AtypeContext
}

func NewParenTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenTypeContext {
	var p = new(ParenTypeContext)

	p.AtypeContext = NewEmptyAtypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtypeContext))

	return p
}

func (s *ParenTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenTypeContext) PAREN_LEFT() antlr.TerminalNode {
	return s.GetToken(LanguageParserPAREN_LEFT, 0)
}

func (s *ParenTypeContext) Type() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ParenTypeContext) PAREN_RIGHT() antlr.TerminalNode {
	return s.GetToken(LanguageParserPAREN_RIGHT, 0)
}

func (s *ParenTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterParenType(s)
	}
}

func (s *ParenTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitParenType(s)
	}
}

func (s *ParenTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitParenType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Atype() (localctx IAtypeContext) {
	this := p
	_ = this

	localctx = NewAtypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LanguageParserRULE_atype)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewConTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(124)
			p.Match(LanguageParserCONID)
		}

	case 2:
		localctx = NewVarTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.Match(LanguageParserVARID)
		}

	case 3:
		localctx = NewTupleTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(126)
			p.Match(LanguageParserPAREN_LEFT)
		}
		{
			p.SetState(127)
			p.Type()
		}
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == LanguageParserCOMMA {
			{
				p.SetState(128)
				p.Match(LanguageParserCOMMA)
			}
			{
				p.SetState(129)
				p.Type()
			}

			p.SetState(132)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(134)
			p.Match(LanguageParserPAREN_RIGHT)
		}

	case 4:
		localctx = NewParenTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(136)
			p.Match(LanguageParserPAREN_LEFT)
		}
		{
			p.SetState(137)
			p.Type()
		}
		{
			p.SetState(138)
			p.Match(LanguageParserPAREN_RIGHT)
		}

	}

	return localctx
}

// IFunlhsContext is an interface to support dynamic dispatch.
type IFunlhsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunlhsContext differentiates from other interfaces.
	IsFunlhsContext()
}

type FunlhsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunlhsContext() *FunlhsContext {
	var p = new(FunlhsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_funlhs
	return p
}

func (*FunlhsContext) IsFunlhsContext() {}

func NewFunlhsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunlhsContext {
	var p = new(FunlhsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_funlhs

	return p
}

func (s *FunlhsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunlhsContext) CopyFrom(ctx *FunlhsContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FunlhsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunlhsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeclLhsContext struct {
	*FunlhsContext
}

func NewDeclLhsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclLhsContext {
	var p = new(DeclLhsContext)

	p.FunlhsContext = NewEmptyFunlhsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunlhsContext))

	return p
}

func (s *DeclLhsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclLhsContext) VARID() antlr.TerminalNode {
	return s.GetToken(LanguageParserVARID, 0)
}

func (s *DeclLhsContext) AllApat() []IApatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IApatContext); ok {
			len++
		}
	}

	tst := make([]IApatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IApatContext); ok {
			tst[i] = t.(IApatContext)
			i++
		}
	}

	return tst
}

func (s *DeclLhsContext) Apat(i int) IApatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IApatContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IApatContext)
}

func (s *DeclLhsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterDeclLhs(s)
	}
}

func (s *DeclLhsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitDeclLhs(s)
	}
}

func (s *DeclLhsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitDeclLhs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Funlhs() (localctx IFunlhsContext) {
	this := p
	_ = this

	localctx = NewFunlhsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LanguageParserRULE_funlhs)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewDeclLhsContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(LanguageParserVARID)
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LanguageParserINT)|(1<<LanguageParserCHAR)|(1<<LanguageParserSTRING)|(1<<LanguageParserVARID)|(1<<LanguageParserCONID))) != 0) {
		{
			p.SetState(143)
			p.Apat()
		}

		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRhsContext is an interface to support dynamic dispatch.
type IRhsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRhsContext differentiates from other interfaces.
	IsRhsContext()
}

type RhsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRhsContext() *RhsContext {
	var p = new(RhsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_rhs
	return p
}

func (*RhsContext) IsRhsContext() {}

func NewRhsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RhsContext {
	var p = new(RhsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_rhs

	return p
}

func (s *RhsContext) GetParser() antlr.Parser { return s.parser }

func (s *RhsContext) CopyFrom(ctx *RhsContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *RhsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RhsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeclExpContext struct {
	*RhsContext
}

func NewDeclExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclExpContext {
	var p = new(DeclExpContext)

	p.RhsContext = NewEmptyRhsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RhsContext))

	return p
}

func (s *DeclExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclExpContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(LanguageParserASSIGN, 0)
}

func (s *DeclExpContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *DeclExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterDeclExp(s)
	}
}

func (s *DeclExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitDeclExp(s)
	}
}

func (s *DeclExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitDeclExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Rhs() (localctx IRhsContext) {
	this := p
	_ = this

	localctx = NewRhsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LanguageParserRULE_rhs)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewDeclExpContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(LanguageParserASSIGN)
	}
	{
		p.SetState(149)
		p.exp(0)
	}

	return localctx
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) CopyFrom(ctx *ExpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EMulDivContext struct {
	*ExpContext
	e1 IExpContext
	op antlr.Token
	e2 IExpContext
}

func NewEMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EMulDivContext {
	var p = new(EMulDivContext)

	p.ExpContext = NewEmptyExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpContext))

	return p
}

func (s *EMulDivContext) GetOp() antlr.Token { return s.op }

func (s *EMulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *EMulDivContext) GetE1() IExpContext { return s.e1 }

func (s *EMulDivContext) GetE2() IExpContext { return s.e2 }

func (s *EMulDivContext) SetE1(v IExpContext) { s.e1 = v }

func (s *EMulDivContext) SetE2(v IExpContext) { s.e2 = v }

func (s *EMulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EMulDivContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *EMulDivContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *EMulDivContext) MUL() antlr.TerminalNode {
	return s.GetToken(LanguageParserMUL, 0)
}

func (s *EMulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(LanguageParserDIV, 0)
}

func (s *EMulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterEMulDiv(s)
	}
}

func (s *EMulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitEMulDiv(s)
	}
}

func (s *EMulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitEMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type ECaseContext struct {
	*ExpContext
}

func NewECaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ECaseContext {
	var p = new(ECaseContext)

	p.ExpContext = NewEmptyExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpContext))

	return p
}

func (s *ECaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ECaseContext) CASE() antlr.TerminalNode {
	return s.GetToken(LanguageParserCASE, 0)
}

func (s *ECaseContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ECaseContext) OF() antlr.TerminalNode {
	return s.GetToken(LanguageParserOF, 0)
}

func (s *ECaseContext) BRACE_LEFT() antlr.TerminalNode {
	return s.GetToken(LanguageParserBRACE_LEFT, 0)
}

func (s *ECaseContext) Alts() IAltsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAltsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAltsContext)
}

func (s *ECaseContext) BRACE_RIGHT() antlr.TerminalNode {
	return s.GetToken(LanguageParserBRACE_RIGHT, 0)
}

func (s *ECaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterECase(s)
	}
}

func (s *ECaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitECase(s)
	}
}

func (s *ECaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitECase(s)

	default:
		return t.VisitChildren(s)
	}
}

type EFunContext struct {
	*ExpContext
}

func NewEFunContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EFunContext {
	var p = new(EFunContext)

	p.ExpContext = NewEmptyExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpContext))

	return p
}

func (s *EFunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EFunContext) Fexp() IFexpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFexpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFexpContext)
}

func (s *EFunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterEFun(s)
	}
}

func (s *EFunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitEFun(s)
	}
}

func (s *EFunContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitEFun(s)

	default:
		return t.VisitChildren(s)
	}
}

type EAddSubContext struct {
	*ExpContext
	e1 IExpContext
	op antlr.Token
	e2 IExpContext
}

func NewEAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EAddSubContext {
	var p = new(EAddSubContext)

	p.ExpContext = NewEmptyExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpContext))

	return p
}

func (s *EAddSubContext) GetOp() antlr.Token { return s.op }

func (s *EAddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *EAddSubContext) GetE1() IExpContext { return s.e1 }

func (s *EAddSubContext) GetE2() IExpContext { return s.e2 }

func (s *EAddSubContext) SetE1(v IExpContext) { s.e1 = v }

func (s *EAddSubContext) SetE2(v IExpContext) { s.e2 = v }

func (s *EAddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EAddSubContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *EAddSubContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *EAddSubContext) ADD() antlr.TerminalNode {
	return s.GetToken(LanguageParserADD, 0)
}

func (s *EAddSubContext) SUB() antlr.TerminalNode {
	return s.GetToken(LanguageParserSUB, 0)
}

func (s *EAddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterEAddSub(s)
	}
}

func (s *EAddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitEAddSub(s)
	}
}

func (s *EAddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitEAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Exp() (localctx IExpContext) {
	return p.exp(0)
}

func (p *LanguageParser) exp(_p int) (localctx IExpContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 28
	p.EnterRecursionRule(localctx, 28, LanguageParserRULE_exp, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(160)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LanguageParserPAREN_LEFT, LanguageParserINT, LanguageParserCHAR, LanguageParserSTRING, LanguageParserVARID, LanguageParserCONID:
		localctx = NewEFunContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(152)
			p.fexp(0)
		}

	case LanguageParserCASE:
		localctx = NewECaseContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(153)
			p.Match(LanguageParserCASE)
		}
		{
			p.SetState(154)
			p.exp(0)
		}
		{
			p.SetState(155)
			p.Match(LanguageParserOF)
		}
		{
			p.SetState(156)
			p.Match(LanguageParserBRACE_LEFT)
		}
		{
			p.SetState(157)
			p.Alts()
		}
		{
			p.SetState(158)
			p.Match(LanguageParserBRACE_RIGHT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(168)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEMulDivContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*EMulDivContext).e1 = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LanguageParserRULE_exp)
				p.SetState(162)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(163)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EMulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == LanguageParserMUL || _la == LanguageParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EMulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(164)

					var _x = p.exp(3)

					localctx.(*EMulDivContext).e2 = _x
				}

			case 2:
				localctx = NewEAddSubContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*EAddSubContext).e1 = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LanguageParserRULE_exp)
				p.SetState(165)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(166)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EAddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == LanguageParserADD || _la == LanguageParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EAddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(167)

					var _x = p.exp(2)

					localctx.(*EAddSubContext).e2 = _x
				}

			}

		}
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}

	return localctx
}

// IFexpContext is an interface to support dynamic dispatch.
type IFexpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFexpContext differentiates from other interfaces.
	IsFexpContext()
}

type FexpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFexpContext() *FexpContext {
	var p = new(FexpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_fexp
	return p
}

func (*FexpContext) IsFexpContext() {}

func NewFexpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FexpContext {
	var p = new(FexpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_fexp

	return p
}

func (s *FexpContext) GetParser() antlr.Parser { return s.parser }

func (s *FexpContext) CopyFrom(ctx *FexpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FexpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FexpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FArgContext struct {
	*FexpContext
}

func NewFArgContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FArgContext {
	var p = new(FArgContext)

	p.FexpContext = NewEmptyFexpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FexpContext))

	return p
}

func (s *FArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FArgContext) Aexp() IAexpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAexpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAexpContext)
}

func (s *FArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterFArg(s)
	}
}

func (s *FArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitFArg(s)
	}
}

func (s *FArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitFArg(s)

	default:
		return t.VisitChildren(s)
	}
}

type FAppContext struct {
	*FexpContext
}

func NewFAppContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FAppContext {
	var p = new(FAppContext)

	p.FexpContext = NewEmptyFexpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FexpContext))

	return p
}

func (s *FAppContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FAppContext) Fexp() IFexpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFexpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFexpContext)
}

func (s *FAppContext) Aexp() IAexpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAexpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAexpContext)
}

func (s *FAppContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterFApp(s)
	}
}

func (s *FAppContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitFApp(s)
	}
}

func (s *FAppContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitFApp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Fexp() (localctx IFexpContext) {
	return p.fexp(0)
}

func (p *LanguageParser) fexp(_p int) (localctx IFexpContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewFexpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFexpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, LanguageParserRULE_fexp, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewFArgContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(174)
		p.Aexp()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFAppContext(p, NewFexpContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, LanguageParserRULE_fexp)
			p.SetState(176)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(177)
				p.Aexp()
			}

		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}

	return localctx
}

// IAexpContext is an interface to support dynamic dispatch.
type IAexpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAexpContext differentiates from other interfaces.
	IsAexpContext()
}

type AexpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAexpContext() *AexpContext {
	var p = new(AexpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_aexp
	return p
}

func (*AexpContext) IsAexpContext() {}

func NewAexpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AexpContext {
	var p = new(AexpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_aexp

	return p
}

func (s *AexpContext) GetParser() antlr.Parser { return s.parser }

func (s *AexpContext) CopyFrom(ctx *AexpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AexpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AexpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VarContext struct {
	*AexpContext
}

func NewVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarContext {
	var p = new(VarContext)

	p.AexpContext = NewEmptyAexpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AexpContext))

	return p
}

func (s *VarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarContext) VARID() antlr.TerminalNode {
	return s.GetToken(LanguageParserVARID, 0)
}

func (s *VarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterVar(s)
	}
}

func (s *VarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitVar(s)
	}
}

func (s *VarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitVar(s)

	default:
		return t.VisitChildren(s)
	}
}

type LitContext struct {
	*AexpContext
}

func NewLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitContext {
	var p = new(LitContext)

	p.AexpContext = NewEmptyAexpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AexpContext))

	return p
}

func (s *LitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterLit(s)
	}
}

func (s *LitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitLit(s)
	}
}

func (s *LitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitLit(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExpContext struct {
	*AexpContext
	e IExpContext
}

func NewParenExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpContext {
	var p = new(ParenExpContext)

	p.AexpContext = NewEmptyAexpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AexpContext))

	return p
}

func (s *ParenExpContext) GetE() IExpContext { return s.e }

func (s *ParenExpContext) SetE(v IExpContext) { s.e = v }

func (s *ParenExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpContext) PAREN_LEFT() antlr.TerminalNode {
	return s.GetToken(LanguageParserPAREN_LEFT, 0)
}

func (s *ParenExpContext) PAREN_RIGHT() antlr.TerminalNode {
	return s.GetToken(LanguageParserPAREN_RIGHT, 0)
}

func (s *ParenExpContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ParenExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterParenExp(s)
	}
}

func (s *ParenExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitParenExp(s)
	}
}

func (s *ParenExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitParenExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConstrContext struct {
	*AexpContext
}

func NewConstrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstrContext {
	var p = new(ConstrContext)

	p.AexpContext = NewEmptyAexpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AexpContext))

	return p
}

func (s *ConstrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstrContext) CONID() antlr.TerminalNode {
	return s.GetToken(LanguageParserCONID, 0)
}

func (s *ConstrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterConstr(s)
	}
}

func (s *ConstrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitConstr(s)
	}
}

func (s *ConstrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitConstr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TupleContext struct {
	*AexpContext
}

func NewTupleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleContext {
	var p = new(TupleContext)

	p.AexpContext = NewEmptyAexpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AexpContext))

	return p
}

func (s *TupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleContext) PAREN_LEFT() antlr.TerminalNode {
	return s.GetToken(LanguageParserPAREN_LEFT, 0)
}

func (s *TupleContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *TupleContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *TupleContext) PAREN_RIGHT() antlr.TerminalNode {
	return s.GetToken(LanguageParserPAREN_RIGHT, 0)
}

func (s *TupleContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LanguageParserCOMMA)
}

func (s *TupleContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LanguageParserCOMMA, i)
}

func (s *TupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterTuple(s)
	}
}

func (s *TupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitTuple(s)
	}
}

func (s *TupleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitTuple(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Aexp() (localctx IAexpContext) {
	this := p
	_ = this

	localctx = NewAexpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LanguageParserRULE_aexp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(183)
			p.Match(LanguageParserVARID)
		}

	case 2:
		localctx = NewConstrContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(184)
			p.Match(LanguageParserCONID)
		}

	case 3:
		localctx = NewLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(185)
			p.Literal()
		}

	case 4:
		localctx = NewParenExpContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(186)
			p.Match(LanguageParserPAREN_LEFT)
		}
		{
			p.SetState(187)

			var _x = p.exp(0)

			localctx.(*ParenExpContext).e = _x
		}
		{
			p.SetState(188)
			p.Match(LanguageParserPAREN_RIGHT)
		}

	case 5:
		localctx = NewTupleContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(190)
			p.Match(LanguageParserPAREN_LEFT)
		}
		{
			p.SetState(191)
			p.exp(0)
		}
		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == LanguageParserCOMMA {
			{
				p.SetState(192)
				p.Match(LanguageParserCOMMA)
			}
			{
				p.SetState(193)
				p.exp(0)
			}

			p.SetState(196)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(198)
			p.Match(LanguageParserPAREN_RIGHT)
		}

	}

	return localctx
}

// IAltsContext is an interface to support dynamic dispatch.
type IAltsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAltsContext differentiates from other interfaces.
	IsAltsContext()
}

type AltsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAltsContext() *AltsContext {
	var p = new(AltsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_alts
	return p
}

func (*AltsContext) IsAltsContext() {}

func NewAltsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AltsContext {
	var p = new(AltsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_alts

	return p
}

func (s *AltsContext) GetParser() antlr.Parser { return s.parser }

func (s *AltsContext) CopyFrom(ctx *AltsContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AltsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AltsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AlternativesContext struct {
	*AltsContext
}

func NewAlternativesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AlternativesContext {
	var p = new(AlternativesContext)

	p.AltsContext = NewEmptyAltsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AltsContext))

	return p
}

func (s *AlternativesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlternativesContext) AllAlt() []IAltContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAltContext); ok {
			len++
		}
	}

	tst := make([]IAltContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAltContext); ok {
			tst[i] = t.(IAltContext)
			i++
		}
	}

	return tst
}

func (s *AlternativesContext) Alt(i int) IAltContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAltContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAltContext)
}

func (s *AlternativesContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(LanguageParserSEMICOLON)
}

func (s *AlternativesContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(LanguageParserSEMICOLON, i)
}

func (s *AlternativesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterAlternatives(s)
	}
}

func (s *AlternativesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitAlternatives(s)
	}
}

func (s *AlternativesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitAlternatives(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Alts() (localctx IAltsContext) {
	this := p
	_ = this

	localctx = NewAltsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, LanguageParserRULE_alts)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	localctx = NewAlternativesContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Alt()
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(203)
				p.Match(LanguageParserSEMICOLON)
			}
			{
				p.SetState(204)
				p.Alt()
			}

		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LanguageParserSEMICOLON {
		{
			p.SetState(210)
			p.Match(LanguageParserSEMICOLON)
		}

	}

	return localctx
}

// IAltContext is an interface to support dynamic dispatch.
type IAltContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAltContext differentiates from other interfaces.
	IsAltContext()
}

type AltContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAltContext() *AltContext {
	var p = new(AltContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_alt
	return p
}

func (*AltContext) IsAltContext() {}

func NewAltContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AltContext {
	var p = new(AltContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_alt

	return p
}

func (s *AltContext) GetParser() antlr.Parser { return s.parser }

func (s *AltContext) CopyFrom(ctx *AltContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AltContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AltContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AlternativeContext struct {
	*AltContext
}

func NewAlternativeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AlternativeContext {
	var p = new(AlternativeContext)

	p.AltContext = NewEmptyAltContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AltContext))

	return p
}

func (s *AlternativeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlternativeContext) Pat() IPatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatContext)
}

func (s *AlternativeContext) ARROW() antlr.TerminalNode {
	return s.GetToken(LanguageParserARROW, 0)
}

func (s *AlternativeContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AlternativeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterAlternative(s)
	}
}

func (s *AlternativeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitAlternative(s)
	}
}

func (s *AlternativeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitAlternative(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Alt() (localctx IAltContext) {
	this := p
	_ = this

	localctx = NewAltContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, LanguageParserRULE_alt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewAlternativeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Pat()
	}
	{
		p.SetState(214)
		p.Match(LanguageParserARROW)
	}
	{
		p.SetState(215)
		p.exp(0)
	}

	return localctx
}

// IPatContext is an interface to support dynamic dispatch.
type IPatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPatContext differentiates from other interfaces.
	IsPatContext()
}

type PatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatContext() *PatContext {
	var p = new(PatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_pat
	return p
}

func (*PatContext) IsPatContext() {}

func NewPatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatContext {
	var p = new(PatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_pat

	return p
}

func (s *PatContext) GetParser() antlr.Parser { return s.parser }

func (s *PatContext) CopyFrom(ctx *PatContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PatArgContext struct {
	*PatContext
}

func NewPatArgContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PatArgContext {
	var p = new(PatArgContext)

	p.PatContext = NewEmptyPatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PatContext))

	return p
}

func (s *PatArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatArgContext) Apat() IApatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IApatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IApatContext)
}

func (s *PatArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterPatArg(s)
	}
}

func (s *PatArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitPatArg(s)
	}
}

func (s *PatArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitPatArg(s)

	default:
		return t.VisitChildren(s)
	}
}

type PatConContext struct {
	*PatContext
}

func NewPatConContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PatConContext {
	var p = new(PatConContext)

	p.PatContext = NewEmptyPatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PatContext))

	return p
}

func (s *PatConContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatConContext) CONID() antlr.TerminalNode {
	return s.GetToken(LanguageParserCONID, 0)
}

func (s *PatConContext) AllApat() []IApatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IApatContext); ok {
			len++
		}
	}

	tst := make([]IApatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IApatContext); ok {
			tst[i] = t.(IApatContext)
			i++
		}
	}

	return tst
}

func (s *PatConContext) Apat(i int) IApatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IApatContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IApatContext)
}

func (s *PatConContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterPatCon(s)
	}
}

func (s *PatConContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitPatCon(s)
	}
}

func (s *PatConContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitPatCon(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Pat() (localctx IPatContext) {
	this := p
	_ = this

	localctx = NewPatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, LanguageParserRULE_pat)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPatArgContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(217)
			p.Apat()
		}

	case 2:
		localctx = NewPatConContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(218)
			p.Match(LanguageParserCONID)
		}
		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LanguageParserINT)|(1<<LanguageParserCHAR)|(1<<LanguageParserSTRING)|(1<<LanguageParserVARID)|(1<<LanguageParserCONID))) != 0) {
			{
				p.SetState(219)
				p.Apat()
			}

			p.SetState(222)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IApatContext is an interface to support dynamic dispatch.
type IApatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsApatContext differentiates from other interfaces.
	IsApatContext()
}

type ApatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyApatContext() *ApatContext {
	var p = new(ApatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_apat
	return p
}

func (*ApatContext) IsApatContext() {}

func NewApatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApatContext {
	var p = new(ApatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_apat

	return p
}

func (s *ApatContext) GetParser() antlr.Parser { return s.parser }

func (s *ApatContext) CopyFrom(ctx *ApatContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ApatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ApatConContext struct {
	*ApatContext
}

func NewApatConContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ApatConContext {
	var p = new(ApatConContext)

	p.ApatContext = NewEmptyApatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ApatContext))

	return p
}

func (s *ApatConContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApatConContext) CONID() antlr.TerminalNode {
	return s.GetToken(LanguageParserCONID, 0)
}

func (s *ApatConContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterApatCon(s)
	}
}

func (s *ApatConContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitApatCon(s)
	}
}

func (s *ApatConContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitApatCon(s)

	default:
		return t.VisitChildren(s)
	}
}

type ApatVarContext struct {
	*ApatContext
}

func NewApatVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ApatVarContext {
	var p = new(ApatVarContext)

	p.ApatContext = NewEmptyApatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ApatContext))

	return p
}

func (s *ApatVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApatVarContext) VARID() antlr.TerminalNode {
	return s.GetToken(LanguageParserVARID, 0)
}

func (s *ApatVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterApatVar(s)
	}
}

func (s *ApatVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitApatVar(s)
	}
}

func (s *ApatVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitApatVar(s)

	default:
		return t.VisitChildren(s)
	}
}

type ApatLitContext struct {
	*ApatContext
}

func NewApatLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ApatLitContext {
	var p = new(ApatLitContext)

	p.ApatContext = NewEmptyApatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ApatContext))

	return p
}

func (s *ApatLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApatLitContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ApatLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterApatLit(s)
	}
}

func (s *ApatLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitApatLit(s)
	}
}

func (s *ApatLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitApatLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Apat() (localctx IApatContext) {
	this := p
	_ = this

	localctx = NewApatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, LanguageParserRULE_apat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(229)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LanguageParserVARID:
		localctx = NewApatVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(226)
			p.Match(LanguageParserVARID)
		}

	case LanguageParserCONID:
		localctx = NewApatConContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(227)
			p.Match(LanguageParserCONID)
		}

	case LanguageParserINT, LanguageParserCHAR, LanguageParserSTRING:
		localctx = NewApatLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(228)
			p.Literal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyFrom(ctx *LiteralContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CharContext struct {
	*LiteralContext
}

func NewCharContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CharContext {
	var p = new(CharContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *CharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharContext) CHAR() antlr.TerminalNode {
	return s.GetToken(LanguageParserCHAR, 0)
}

func (s *CharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterChar(s)
	}
}

func (s *CharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitChar(s)
	}
}

func (s *CharContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitChar(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringContext struct {
	*LiteralContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(LanguageParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntContext struct {
	*LiteralContext
}

func NewIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntContext {
	var p = new(IntContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) INT() antlr.TerminalNode {
	return s.GetToken(LanguageParserINT, 0)
}

func (s *IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterInt(s)
	}
}

func (s *IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitInt(s)
	}
}

func (s *IntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LanguageVisitor:
		return t.VisitInt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LanguageParser) Literal() (localctx ILiteralContext) {
	this := p
	_ = this

	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, LanguageParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(234)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LanguageParserINT:
		localctx = NewIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(231)
			p.Match(LanguageParserINT)
		}

	case LanguageParserCHAR:
		localctx = NewCharContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(232)
			p.Match(LanguageParserCHAR)
		}

	case LanguageParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(233)
			p.Match(LanguageParserSTRING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *LanguageParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 14:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	case 15:
		var t *FexpContext = nil
		if localctx != nil {
			t = localctx.(*FexpContext)
		}
		return p.Fexp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LanguageParser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *LanguageParser) Fexp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

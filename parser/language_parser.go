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
		"", "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'{'", "'}'", "','", "';'",
		"'case'", "'of'", "'->'",
	}
	staticData.symbolicNames = []string{
		"", "MUL", "DIV", "ADD", "SUB", "PAREN_LEFT", "PAREN_RIGHT", "BRACE_LEFT",
		"BRACE_RIGHT", "COMMA", "SEMICOLON", "CASE", "OF", "ARROW", "INT", "CHAR",
		"STRING", "VARID", "CONID", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"start", "exp", "fexp", "aexp", "alts", "alt", "pat", "apat", "literal",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 19, 104, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 31, 8, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 39, 8, 1, 10, 1, 12, 1, 42, 9, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 49, 8, 2, 10, 2, 12, 2, 52, 9, 2, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 4, 3, 65, 8,
		3, 11, 3, 12, 3, 66, 1, 3, 1, 3, 3, 3, 71, 8, 3, 1, 4, 1, 4, 1, 4, 5, 4,
		76, 8, 4, 10, 4, 12, 4, 79, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1,
		6, 4, 6, 88, 8, 6, 11, 6, 12, 6, 89, 3, 6, 92, 8, 6, 1, 7, 1, 7, 1, 7,
		3, 7, 97, 8, 7, 1, 8, 1, 8, 1, 8, 3, 8, 102, 8, 8, 1, 8, 0, 2, 2, 4, 9,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 2, 1, 0, 1, 2, 1, 0, 3, 4, 110, 0, 18,
		1, 0, 0, 0, 2, 30, 1, 0, 0, 0, 4, 43, 1, 0, 0, 0, 6, 70, 1, 0, 0, 0, 8,
		72, 1, 0, 0, 0, 10, 80, 1, 0, 0, 0, 12, 91, 1, 0, 0, 0, 14, 96, 1, 0, 0,
		0, 16, 101, 1, 0, 0, 0, 18, 19, 3, 2, 1, 0, 19, 20, 5, 0, 0, 1, 20, 1,
		1, 0, 0, 0, 21, 22, 6, 1, -1, 0, 22, 31, 3, 4, 2, 0, 23, 24, 5, 11, 0,
		0, 24, 25, 3, 2, 1, 0, 25, 26, 5, 12, 0, 0, 26, 27, 5, 7, 0, 0, 27, 28,
		3, 8, 4, 0, 28, 29, 5, 8, 0, 0, 29, 31, 1, 0, 0, 0, 30, 21, 1, 0, 0, 0,
		30, 23, 1, 0, 0, 0, 31, 40, 1, 0, 0, 0, 32, 33, 10, 2, 0, 0, 33, 34, 7,
		0, 0, 0, 34, 39, 3, 2, 1, 3, 35, 36, 10, 1, 0, 0, 36, 37, 7, 1, 0, 0, 37,
		39, 3, 2, 1, 2, 38, 32, 1, 0, 0, 0, 38, 35, 1, 0, 0, 0, 39, 42, 1, 0, 0,
		0, 40, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 3, 1, 0, 0, 0, 42, 40, 1,
		0, 0, 0, 43, 44, 6, 2, -1, 0, 44, 45, 3, 6, 3, 0, 45, 50, 1, 0, 0, 0, 46,
		47, 10, 2, 0, 0, 47, 49, 3, 6, 3, 0, 48, 46, 1, 0, 0, 0, 49, 52, 1, 0,
		0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 5, 1, 0, 0, 0, 52, 50,
		1, 0, 0, 0, 53, 71, 5, 17, 0, 0, 54, 71, 5, 18, 0, 0, 55, 71, 3, 16, 8,
		0, 56, 57, 5, 5, 0, 0, 57, 58, 3, 2, 1, 0, 58, 59, 5, 6, 0, 0, 59, 71,
		1, 0, 0, 0, 60, 61, 5, 5, 0, 0, 61, 64, 3, 2, 1, 0, 62, 63, 5, 9, 0, 0,
		63, 65, 3, 2, 1, 0, 64, 62, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 64, 1,
		0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 69, 5, 6, 0, 0, 69,
		71, 1, 0, 0, 0, 70, 53, 1, 0, 0, 0, 70, 54, 1, 0, 0, 0, 70, 55, 1, 0, 0,
		0, 70, 56, 1, 0, 0, 0, 70, 60, 1, 0, 0, 0, 71, 7, 1, 0, 0, 0, 72, 77, 3,
		10, 5, 0, 73, 74, 5, 10, 0, 0, 74, 76, 3, 10, 5, 0, 75, 73, 1, 0, 0, 0,
		76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 9, 1, 0,
		0, 0, 79, 77, 1, 0, 0, 0, 80, 81, 3, 12, 6, 0, 81, 82, 5, 13, 0, 0, 82,
		83, 3, 2, 1, 0, 83, 11, 1, 0, 0, 0, 84, 92, 3, 14, 7, 0, 85, 87, 5, 18,
		0, 0, 86, 88, 3, 14, 7, 0, 87, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89,
		87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 92, 1, 0, 0, 0, 91, 84, 1, 0, 0,
		0, 91, 85, 1, 0, 0, 0, 92, 13, 1, 0, 0, 0, 93, 97, 5, 17, 0, 0, 94, 97,
		5, 18, 0, 0, 95, 97, 3, 16, 8, 0, 96, 93, 1, 0, 0, 0, 96, 94, 1, 0, 0,
		0, 96, 95, 1, 0, 0, 0, 97, 15, 1, 0, 0, 0, 98, 102, 5, 14, 0, 0, 99, 102,
		5, 15, 0, 0, 100, 102, 5, 16, 0, 0, 101, 98, 1, 0, 0, 0, 101, 99, 1, 0,
		0, 0, 101, 100, 1, 0, 0, 0, 102, 17, 1, 0, 0, 0, 11, 30, 38, 40, 50, 66,
		70, 77, 89, 91, 96, 101,
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
	LanguageParserEOF         = antlr.TokenEOF
	LanguageParserMUL         = 1
	LanguageParserDIV         = 2
	LanguageParserADD         = 3
	LanguageParserSUB         = 4
	LanguageParserPAREN_LEFT  = 5
	LanguageParserPAREN_RIGHT = 6
	LanguageParserBRACE_LEFT  = 7
	LanguageParserBRACE_RIGHT = 8
	LanguageParserCOMMA       = 9
	LanguageParserSEMICOLON   = 10
	LanguageParserCASE        = 11
	LanguageParserOF          = 12
	LanguageParserARROW       = 13
	LanguageParserINT         = 14
	LanguageParserCHAR        = 15
	LanguageParserSTRING      = 16
	LanguageParserVARID       = 17
	LanguageParserCONID       = 18
	LanguageParserWHITESPACE  = 19
)

// LanguageParser rules.
const (
	LanguageParserRULE_start   = 0
	LanguageParserRULE_exp     = 1
	LanguageParserRULE_fexp    = 2
	LanguageParserRULE_aexp    = 3
	LanguageParserRULE_alts    = 4
	LanguageParserRULE_alt     = 5
	LanguageParserRULE_pat     = 6
	LanguageParserRULE_apat    = 7
	LanguageParserRULE_literal = 8
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

func (s *StartContext) Exp() IExpContext {
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
		p.SetState(18)
		p.exp(0)
	}
	{
		p.SetState(19)
		p.Match(LanguageParserEOF)
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
	_startState := 2
	p.EnterRecursionRule(localctx, 2, LanguageParserRULE_exp, _p)
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
	p.SetState(30)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LanguageParserPAREN_LEFT, LanguageParserINT, LanguageParserCHAR, LanguageParserSTRING, LanguageParserVARID, LanguageParserCONID:
		localctx = NewEFunContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(22)
			p.fexp(0)
		}

	case LanguageParserCASE:
		localctx = NewECaseContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(23)
			p.Match(LanguageParserCASE)
		}
		{
			p.SetState(24)
			p.exp(0)
		}
		{
			p.SetState(25)
			p.Match(LanguageParserOF)
		}
		{
			p.SetState(26)
			p.Match(LanguageParserBRACE_LEFT)
		}
		{
			p.SetState(27)
			p.Alts()
		}
		{
			p.SetState(28)
			p.Match(LanguageParserBRACE_RIGHT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(38)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEMulDivContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*EMulDivContext).e1 = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LanguageParserRULE_exp)
				p.SetState(32)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(33)

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
					p.SetState(34)

					var _x = p.exp(3)

					localctx.(*EMulDivContext).e2 = _x
				}

			case 2:
				localctx = NewEAddSubContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*EAddSubContext).e1 = _prevctx

				p.PushNewRecursionContext(localctx, _startState, LanguageParserRULE_exp)
				p.SetState(35)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(36)

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
					p.SetState(37)

					var _x = p.exp(2)

					localctx.(*EAddSubContext).e2 = _x
				}

			}

		}
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
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
	_startState := 4
	p.EnterRecursionRule(localctx, 4, LanguageParserRULE_fexp, _p)

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
		p.SetState(44)
		p.Aexp()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFAppContext(p, NewFexpContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, LanguageParserRULE_fexp)
			p.SetState(46)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(47)
				p.Aexp()
			}

		}
		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 6, LanguageParserRULE_aexp)
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

	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.Match(LanguageParserVARID)
		}

	case 2:
		localctx = NewConstrContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Match(LanguageParserCONID)
		}

	case 3:
		localctx = NewLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(55)
			p.Literal()
		}

	case 4:
		localctx = NewParenExpContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(56)
			p.Match(LanguageParserPAREN_LEFT)
		}
		{
			p.SetState(57)

			var _x = p.exp(0)

			localctx.(*ParenExpContext).e = _x
		}
		{
			p.SetState(58)
			p.Match(LanguageParserPAREN_RIGHT)
		}

	case 5:
		localctx = NewTupleContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(60)
			p.Match(LanguageParserPAREN_LEFT)
		}
		{
			p.SetState(61)
			p.exp(0)
		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == LanguageParserCOMMA {
			{
				p.SetState(62)
				p.Match(LanguageParserCOMMA)
			}
			{
				p.SetState(63)
				p.exp(0)
			}

			p.SetState(66)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(68)
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
	p.EnterRule(localctx, 8, LanguageParserRULE_alts)
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

	localctx = NewAlternativesContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Alt()
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LanguageParserSEMICOLON {
		{
			p.SetState(73)
			p.Match(LanguageParserSEMICOLON)
		}
		{
			p.SetState(74)
			p.Alt()
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 10, LanguageParserRULE_alt)

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
		p.SetState(80)
		p.Pat()
	}
	{
		p.SetState(81)
		p.Match(LanguageParserARROW)
	}
	{
		p.SetState(82)
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
	p.EnterRule(localctx, 12, LanguageParserRULE_pat)
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

	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPatArgContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.Apat()
		}

	case 2:
		localctx = NewPatConContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.Match(LanguageParserCONID)
		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LanguageParserINT)|(1<<LanguageParserCHAR)|(1<<LanguageParserSTRING)|(1<<LanguageParserVARID)|(1<<LanguageParserCONID))) != 0) {
			{
				p.SetState(86)
				p.Apat()
			}

			p.SetState(89)
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
	p.EnterRule(localctx, 14, LanguageParserRULE_apat)

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

	p.SetState(96)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LanguageParserVARID:
		localctx = NewApatVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.Match(LanguageParserVARID)
		}

	case LanguageParserCONID:
		localctx = NewApatConContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.Match(LanguageParserCONID)
		}

	case LanguageParserINT, LanguageParserCHAR, LanguageParserSTRING:
		localctx = NewApatLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(95)
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
	p.EnterRule(localctx, 16, LanguageParserRULE_literal)

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

	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LanguageParserINT:
		localctx = NewIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)
			p.Match(LanguageParserINT)
		}

	case LanguageParserCHAR:
		localctx = NewCharContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(99)
			p.Match(LanguageParserCHAR)
		}

	case LanguageParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(100)
			p.Match(LanguageParserSTRING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *LanguageParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	case 2:
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

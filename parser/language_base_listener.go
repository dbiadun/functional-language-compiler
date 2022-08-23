// Code generated from Language.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Language

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLanguageListener is a complete listener for a parse tree produced by LanguageParser.
type BaseLanguageListener struct{}

var _ LanguageListener = &BaseLanguageListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLanguageListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLanguageListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLanguageListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLanguageListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseLanguageListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseLanguageListener) ExitStart(ctx *StartContext) {}

// EnterEMulDiv is called when production EMulDiv is entered.
func (s *BaseLanguageListener) EnterEMulDiv(ctx *EMulDivContext) {}

// ExitEMulDiv is called when production EMulDiv is exited.
func (s *BaseLanguageListener) ExitEMulDiv(ctx *EMulDivContext) {}

// EnterECase is called when production ECase is entered.
func (s *BaseLanguageListener) EnterECase(ctx *ECaseContext) {}

// ExitECase is called when production ECase is exited.
func (s *BaseLanguageListener) ExitECase(ctx *ECaseContext) {}

// EnterEFun is called when production EFun is entered.
func (s *BaseLanguageListener) EnterEFun(ctx *EFunContext) {}

// ExitEFun is called when production EFun is exited.
func (s *BaseLanguageListener) ExitEFun(ctx *EFunContext) {}

// EnterEAddSub is called when production EAddSub is entered.
func (s *BaseLanguageListener) EnterEAddSub(ctx *EAddSubContext) {}

// ExitEAddSub is called when production EAddSub is exited.
func (s *BaseLanguageListener) ExitEAddSub(ctx *EAddSubContext) {}

// EnterFArg is called when production FArg is entered.
func (s *BaseLanguageListener) EnterFArg(ctx *FArgContext) {}

// ExitFArg is called when production FArg is exited.
func (s *BaseLanguageListener) ExitFArg(ctx *FArgContext) {}

// EnterFApp is called when production FApp is entered.
func (s *BaseLanguageListener) EnterFApp(ctx *FAppContext) {}

// ExitFApp is called when production FApp is exited.
func (s *BaseLanguageListener) ExitFApp(ctx *FAppContext) {}

// EnterVar is called when production Var is entered.
func (s *BaseLanguageListener) EnterVar(ctx *VarContext) {}

// ExitVar is called when production Var is exited.
func (s *BaseLanguageListener) ExitVar(ctx *VarContext) {}

// EnterConstr is called when production Constr is entered.
func (s *BaseLanguageListener) EnterConstr(ctx *ConstrContext) {}

// ExitConstr is called when production Constr is exited.
func (s *BaseLanguageListener) ExitConstr(ctx *ConstrContext) {}

// EnterLit is called when production Lit is entered.
func (s *BaseLanguageListener) EnterLit(ctx *LitContext) {}

// ExitLit is called when production Lit is exited.
func (s *BaseLanguageListener) ExitLit(ctx *LitContext) {}

// EnterParenExp is called when production ParenExp is entered.
func (s *BaseLanguageListener) EnterParenExp(ctx *ParenExpContext) {}

// ExitParenExp is called when production ParenExp is exited.
func (s *BaseLanguageListener) ExitParenExp(ctx *ParenExpContext) {}

// EnterTuple is called when production Tuple is entered.
func (s *BaseLanguageListener) EnterTuple(ctx *TupleContext) {}

// ExitTuple is called when production Tuple is exited.
func (s *BaseLanguageListener) ExitTuple(ctx *TupleContext) {}

// EnterAlternatives is called when production Alternatives is entered.
func (s *BaseLanguageListener) EnterAlternatives(ctx *AlternativesContext) {}

// ExitAlternatives is called when production Alternatives is exited.
func (s *BaseLanguageListener) ExitAlternatives(ctx *AlternativesContext) {}

// EnterAlternative is called when production Alternative is entered.
func (s *BaseLanguageListener) EnterAlternative(ctx *AlternativeContext) {}

// ExitAlternative is called when production Alternative is exited.
func (s *BaseLanguageListener) ExitAlternative(ctx *AlternativeContext) {}

// EnterPatArg is called when production PatArg is entered.
func (s *BaseLanguageListener) EnterPatArg(ctx *PatArgContext) {}

// ExitPatArg is called when production PatArg is exited.
func (s *BaseLanguageListener) ExitPatArg(ctx *PatArgContext) {}

// EnterPatCon is called when production PatCon is entered.
func (s *BaseLanguageListener) EnterPatCon(ctx *PatConContext) {}

// ExitPatCon is called when production PatCon is exited.
func (s *BaseLanguageListener) ExitPatCon(ctx *PatConContext) {}

// EnterApatVar is called when production ApatVar is entered.
func (s *BaseLanguageListener) EnterApatVar(ctx *ApatVarContext) {}

// ExitApatVar is called when production ApatVar is exited.
func (s *BaseLanguageListener) ExitApatVar(ctx *ApatVarContext) {}

// EnterApatCon is called when production ApatCon is entered.
func (s *BaseLanguageListener) EnterApatCon(ctx *ApatConContext) {}

// ExitApatCon is called when production ApatCon is exited.
func (s *BaseLanguageListener) ExitApatCon(ctx *ApatConContext) {}

// EnterApatLit is called when production ApatLit is entered.
func (s *BaseLanguageListener) EnterApatLit(ctx *ApatLitContext) {}

// ExitApatLit is called when production ApatLit is exited.
func (s *BaseLanguageListener) ExitApatLit(ctx *ApatLitContext) {}

// EnterInt is called when production Int is entered.
func (s *BaseLanguageListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production Int is exited.
func (s *BaseLanguageListener) ExitInt(ctx *IntContext) {}

// EnterChar is called when production Char is entered.
func (s *BaseLanguageListener) EnterChar(ctx *CharContext) {}

// ExitChar is called when production Char is exited.
func (s *BaseLanguageListener) ExitChar(ctx *CharContext) {}

// EnterString is called when production String is entered.
func (s *BaseLanguageListener) EnterString(ctx *StringContext) {}

// ExitString is called when production String is exited.
func (s *BaseLanguageListener) ExitString(ctx *StringContext) {}

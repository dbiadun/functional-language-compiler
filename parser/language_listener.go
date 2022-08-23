// Code generated from Language.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Language

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LanguageListener is a complete listener for a parse tree produced by LanguageParser.
type LanguageListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterEMulDiv is called when entering the EMulDiv production.
	EnterEMulDiv(c *EMulDivContext)

	// EnterECase is called when entering the ECase production.
	EnterECase(c *ECaseContext)

	// EnterEFun is called when entering the EFun production.
	EnterEFun(c *EFunContext)

	// EnterEAddSub is called when entering the EAddSub production.
	EnterEAddSub(c *EAddSubContext)

	// EnterFArg is called when entering the FArg production.
	EnterFArg(c *FArgContext)

	// EnterFApp is called when entering the FApp production.
	EnterFApp(c *FAppContext)

	// EnterVar is called when entering the Var production.
	EnterVar(c *VarContext)

	// EnterConstr is called when entering the Constr production.
	EnterConstr(c *ConstrContext)

	// EnterLit is called when entering the Lit production.
	EnterLit(c *LitContext)

	// EnterParenExp is called when entering the ParenExp production.
	EnterParenExp(c *ParenExpContext)

	// EnterTuple is called when entering the Tuple production.
	EnterTuple(c *TupleContext)

	// EnterAlternatives is called when entering the Alternatives production.
	EnterAlternatives(c *AlternativesContext)

	// EnterAlternative is called when entering the Alternative production.
	EnterAlternative(c *AlternativeContext)

	// EnterPatArg is called when entering the PatArg production.
	EnterPatArg(c *PatArgContext)

	// EnterPatCon is called when entering the PatCon production.
	EnterPatCon(c *PatConContext)

	// EnterApatVar is called when entering the ApatVar production.
	EnterApatVar(c *ApatVarContext)

	// EnterApatCon is called when entering the ApatCon production.
	EnterApatCon(c *ApatConContext)

	// EnterApatLit is called when entering the ApatLit production.
	EnterApatLit(c *ApatLitContext)

	// EnterInt is called when entering the Int production.
	EnterInt(c *IntContext)

	// EnterChar is called when entering the Char production.
	EnterChar(c *CharContext)

	// EnterString is called when entering the String production.
	EnterString(c *StringContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitEMulDiv is called when exiting the EMulDiv production.
	ExitEMulDiv(c *EMulDivContext)

	// ExitECase is called when exiting the ECase production.
	ExitECase(c *ECaseContext)

	// ExitEFun is called when exiting the EFun production.
	ExitEFun(c *EFunContext)

	// ExitEAddSub is called when exiting the EAddSub production.
	ExitEAddSub(c *EAddSubContext)

	// ExitFArg is called when exiting the FArg production.
	ExitFArg(c *FArgContext)

	// ExitFApp is called when exiting the FApp production.
	ExitFApp(c *FAppContext)

	// ExitVar is called when exiting the Var production.
	ExitVar(c *VarContext)

	// ExitConstr is called when exiting the Constr production.
	ExitConstr(c *ConstrContext)

	// ExitLit is called when exiting the Lit production.
	ExitLit(c *LitContext)

	// ExitParenExp is called when exiting the ParenExp production.
	ExitParenExp(c *ParenExpContext)

	// ExitTuple is called when exiting the Tuple production.
	ExitTuple(c *TupleContext)

	// ExitAlternatives is called when exiting the Alternatives production.
	ExitAlternatives(c *AlternativesContext)

	// ExitAlternative is called when exiting the Alternative production.
	ExitAlternative(c *AlternativeContext)

	// ExitPatArg is called when exiting the PatArg production.
	ExitPatArg(c *PatArgContext)

	// ExitPatCon is called when exiting the PatCon production.
	ExitPatCon(c *PatConContext)

	// ExitApatVar is called when exiting the ApatVar production.
	ExitApatVar(c *ApatVarContext)

	// ExitApatCon is called when exiting the ApatCon production.
	ExitApatCon(c *ApatConContext)

	// ExitApatLit is called when exiting the ApatLit production.
	ExitApatLit(c *ApatLitContext)

	// ExitInt is called when exiting the Int production.
	ExitInt(c *IntContext)

	// ExitChar is called when exiting the Char production.
	ExitChar(c *CharContext)

	// ExitString is called when exiting the String production.
	ExitString(c *StringContext)
}

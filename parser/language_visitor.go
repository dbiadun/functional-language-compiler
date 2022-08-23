// Code generated from Language.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Language

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by LanguageParser.
type LanguageVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LanguageParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by LanguageParser#EMulDiv.
	VisitEMulDiv(ctx *EMulDivContext) interface{}

	// Visit a parse tree produced by LanguageParser#ECase.
	VisitECase(ctx *ECaseContext) interface{}

	// Visit a parse tree produced by LanguageParser#EFun.
	VisitEFun(ctx *EFunContext) interface{}

	// Visit a parse tree produced by LanguageParser#EAddSub.
	VisitEAddSub(ctx *EAddSubContext) interface{}

	// Visit a parse tree produced by LanguageParser#FArg.
	VisitFArg(ctx *FArgContext) interface{}

	// Visit a parse tree produced by LanguageParser#FApp.
	VisitFApp(ctx *FAppContext) interface{}

	// Visit a parse tree produced by LanguageParser#Var.
	VisitVar(ctx *VarContext) interface{}

	// Visit a parse tree produced by LanguageParser#Constr.
	VisitConstr(ctx *ConstrContext) interface{}

	// Visit a parse tree produced by LanguageParser#Lit.
	VisitLit(ctx *LitContext) interface{}

	// Visit a parse tree produced by LanguageParser#ParenExp.
	VisitParenExp(ctx *ParenExpContext) interface{}

	// Visit a parse tree produced by LanguageParser#Tuple.
	VisitTuple(ctx *TupleContext) interface{}

	// Visit a parse tree produced by LanguageParser#Alternatives.
	VisitAlternatives(ctx *AlternativesContext) interface{}

	// Visit a parse tree produced by LanguageParser#Alternative.
	VisitAlternative(ctx *AlternativeContext) interface{}

	// Visit a parse tree produced by LanguageParser#PatArg.
	VisitPatArg(ctx *PatArgContext) interface{}

	// Visit a parse tree produced by LanguageParser#PatCon.
	VisitPatCon(ctx *PatConContext) interface{}

	// Visit a parse tree produced by LanguageParser#ApatVar.
	VisitApatVar(ctx *ApatVarContext) interface{}

	// Visit a parse tree produced by LanguageParser#ApatCon.
	VisitApatCon(ctx *ApatConContext) interface{}

	// Visit a parse tree produced by LanguageParser#ApatLit.
	VisitApatLit(ctx *ApatLitContext) interface{}

	// Visit a parse tree produced by LanguageParser#Int.
	VisitInt(ctx *IntContext) interface{}

	// Visit a parse tree produced by LanguageParser#Char.
	VisitChar(ctx *CharContext) interface{}

	// Visit a parse tree produced by LanguageParser#String.
	VisitString(ctx *StringContext) interface{}
}

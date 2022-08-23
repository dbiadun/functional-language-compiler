// Code generated from Language.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Language

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseLanguageVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLanguageVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitEMulDiv(ctx *EMulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitECase(ctx *ECaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitEFun(ctx *EFunContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitEAddSub(ctx *EAddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitFArg(ctx *FArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitFApp(ctx *FAppContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitVar(ctx *VarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitConstr(ctx *ConstrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitLit(ctx *LitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitParenExp(ctx *ParenExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitTuple(ctx *TupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitAlternatives(ctx *AlternativesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitAlternative(ctx *AlternativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitPatArg(ctx *PatArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitPatCon(ctx *PatConContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitApatVar(ctx *ApatVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitApatCon(ctx *ApatConContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitApatLit(ctx *ApatLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitInt(ctx *IntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitChar(ctx *CharContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

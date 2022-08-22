// Code generated from Language.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Language

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseLanguageVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLanguageVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitEMulDiv(ctx *EMulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitEAddSub(ctx *EAddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

// Code generated from Language.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Language

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by LanguageParser.
type LanguageVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LanguageParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by LanguageParser#Number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by LanguageParser#EMulDiv.
	VisitEMulDiv(ctx *EMulDivContext) interface{}

	// Visit a parse tree produced by LanguageParser#EAddSub.
	VisitEAddSub(ctx *EAddSubContext) interface{}
}

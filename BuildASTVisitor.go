package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
	"participle-test/parser"
	"strconv"
)

type BuildASTVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BuildASTVisitor) VisitStart(ctx *parser.StartContext) interface{} {
	return ctx.Expr().Accept(v)
}

func (v *BuildASTVisitor) VisitEMulDiv(ctx *parser.EMulDivContext) interface{} {
	subtree1 := ctx.GetE1()
	subtree2 := ctx.GetE2()
	op := ctx.GetOp().GetTokenType()
	line := ctx.GetOp().GetLine()
	col := ctx.GetOp().GetColumn()

	return v.VisitEBinary(subtree1, subtree2, op, line, col)
}
func (v *BuildASTVisitor) VisitEAddSub(ctx *parser.EAddSubContext) interface{} {
	subtree1 := ctx.GetE1()
	subtree2 := ctx.GetE2()
	op := ctx.GetOp().GetTokenType()
	line := ctx.GetOp().GetLine()
	col := ctx.GetOp().GetColumn()

	return v.VisitEBinary(subtree1, subtree2, op, line, col)
}
func (v *BuildASTVisitor) VisitNumber(ctx *parser.NumberContext) interface{} {
	e := new(Number)

	n, err := strconv.Atoi(ctx.GetText())
	line := ctx.GetStart().GetLine()
	col := ctx.GetStart().GetColumn()

	if err == nil {
		e.n = n
	}

	e.setPos(line, col)

	return e
}

func (v *BuildASTVisitor) VisitEBinary(subtree1 parser.IExprContext, subtree2 parser.IExprContext, op int, line int, col int) interface{} {
	var e Expr
	var e1, e2 Expr

	r1 := subtree1.Accept(v)
	r2 := subtree2.Accept(v)

	switch r1 := r1.(type) {
	case Expr:
		e1 = r1
	}
	switch r2 := r2.(type) {
	case Expr:
		e2 = r2
	}

	baseBinaryExpr := BaseEBinary{e1: e1, e2: e2}

	switch op {
	case parser.LanguageLexerMUL:
		e = &EMul{baseBinaryExpr}
	case parser.LanguageLexerDIV:
		e = &EDiv{baseBinaryExpr}
	case parser.LanguageLexerADD:
		e = &EAdd{baseBinaryExpr}
	case parser.LanguageLexerSUB:
		e = &ESub{baseBinaryExpr}
	default:
		log.Fatalln("Invalid operation type")
	}

	e.setPos(line, col)

	return e
}

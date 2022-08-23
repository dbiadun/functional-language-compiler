package main

import (
	"fmt"
	"functional-language-compiler/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
	"strconv"
)

type BuildASTVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BuildASTVisitor) VisitStart(ctx *parser.StartContext) interface{} {
	return ctx.Exp().Accept(v)
}

// exp

func (v *BuildASTVisitor) VisitEFun(ctx *parser.EFunContext) interface{} {
	fe := ctx.Fexp().Accept(v)
	f := fe.(FExp)
	return &EFun{fExp: f}
}

func (v *BuildASTVisitor) VisitECase(ctx *parser.ECaseContext) interface{} {
	return nil
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

func (v *BuildASTVisitor) VisitEBinary(subtree1 parser.IExpContext, subtree2 parser.IExpContext, op int, line int, col int) interface{} {
	var e Exp
	var e1, e2 Exp

	r1 := subtree1.Accept(v)
	r2 := subtree2.Accept(v)

	switch r1 := r1.(type) {
	case Exp:
		fmt.Println(r1)
		e1 = r1
	}
	switch r2 := r2.(type) {
	case Exp:
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

// fExp

func (v *BuildASTVisitor) VisitFApp(ctx *parser.FAppContext) interface{} {
	return nil
}

func (v *BuildASTVisitor) VisitFArg(ctx *parser.FArgContext) interface{} {
	arg := ctx.Aexp().Accept(v)
	a := arg.(AExp)
	return &FArg{arg: a}
}

// aexp

func (v *BuildASTVisitor) VisitVar(ctx *parser.VarContext) interface{} {
	return nil
}

func (v *BuildASTVisitor) VisitConstr(ctx *parser.ConstrContext) interface{} {
	return nil
}

func (v *BuildASTVisitor) VisitLit(ctx *parser.LitContext) interface{} {
	lit := ctx.Literal().Accept(v)
	l := lit.(Literal)
	return &Lit{lit: l}

}

func (v *BuildASTVisitor) VisitParenExp(ctx *parser.ParenExpContext) interface{} {
	return nil
}

func (v *BuildASTVisitor) VisitTuple(ctx *parser.TupleContext) interface{} {
	return nil
}

// alts

func (v *BuildASTVisitor) VisitAlternatives(ctx *parser.AlternativesContext) interface{} {
	return nil
}

func (v *BuildASTVisitor) VisitAlternative(ctx *parser.AlternativeContext) interface{} {
	return nil
}

// pat

func (v *BuildASTVisitor) VisitPatArg(ctx *parser.PatArgContext) interface{} {
	return nil
}

func (v *BuildASTVisitor) VisitPatCon(ctx *parser.PatConContext) interface{} {
	return nil
}

// apat

func (v *BuildASTVisitor) VisitApatVar(ctx *parser.ApatVarContext) interface{} {
	return nil
}

func (v *BuildASTVisitor) VisitApatCon(ctx *parser.ApatConContext) interface{} {
	return nil
}

func (v *BuildASTVisitor) VisitApatLit(ctx *parser.ApatLitContext) interface{} {
	return nil
}

// literal

func (v *BuildASTVisitor) VisitInt(ctx *parser.IntContext) interface{} {
	e := new(Int)

	n, err := strconv.Atoi(ctx.GetText())
	line := ctx.GetStart().GetLine()
	col := ctx.GetStart().GetColumn()

	if err == nil {
		e.n = n
	}

	e.setPos(line, col)

	return e
}

func (v *BuildASTVisitor) VisitChar(ctx *parser.CharContext) interface{} {
	return nil
}

func (v *BuildASTVisitor) VisitString(ctx *parser.StringContext) interface{} {
	return nil
}

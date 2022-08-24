package main

import (
	"functional-language-compiler/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
	"strconv"
	"strings"
)

type BuildASTVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BuildASTVisitor) VisitStart(ctx *parser.StartContext) interface{} {
	return ctx.Exp().Accept(v)
}

// exp

func (v *BuildASTVisitor) VisitEFun(ctx *parser.EFunContext) interface{} {
	r := new(EFun)
	r.setPosFromCtx(ctx)

	r.fExp = ctx.Fexp().Accept(v).(FExp)

	return r
}

func (v *BuildASTVisitor) VisitECase(ctx *parser.ECaseContext) interface{} {
	r := new(ECase)
	r.setPosFromCtx(ctx)

	r.e = ctx.Exp().Accept(v).(Exp)
	r.alts = ctx.Alts().Accept(v).([]*Alternative)

	return r
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
	var e EBinary

	switch op {
	case parser.LanguageLexerMUL:
		e = new(EMul)
	case parser.LanguageLexerDIV:
		e = new(EDiv)
	case parser.LanguageLexerADD:
		e = new(EAdd)
	case parser.LanguageLexerSUB:
		e = new(ESub)
	default:
		log.Fatalln("Invalid operation type")
	}

	e.setPos(line, col)

	e1 := subtree1.Accept(v).(Exp)
	e2 := subtree2.Accept(v).(Exp)

	e.setExps(e1, e2)

	return e
}

// fExp

func (v *BuildASTVisitor) VisitFApp(ctx *parser.FAppContext) interface{} {
	r := new(FApp)
	r.setPosFromCtx(ctx)

	r.fun = ctx.Fexp().Accept(v).(FExp)
	r.arg = ctx.Aexp().Accept(v).(AExp)

	return r
}

func (v *BuildASTVisitor) VisitFArg(ctx *parser.FArgContext) interface{} {
	r := new(FArg)
	r.setPosFromCtx(ctx)

	r.arg = ctx.Aexp().Accept(v).(AExp)

	return r
}

// aexp

func (v *BuildASTVisitor) VisitVar(ctx *parser.VarContext) interface{} {
	r := new(Var)
	r.setPosFromCtx(ctx)

	r.id = ctx.VARID().GetText()

	return r
}

func (v *BuildASTVisitor) VisitConstr(ctx *parser.ConstrContext) interface{} {
	r := new(Constr)
	r.setPosFromCtx(ctx)

	r.id = ctx.CONID().GetText()

	return r
}

func (v *BuildASTVisitor) VisitLit(ctx *parser.LitContext) interface{} {
	r := new(Lit)
	r.setPosFromCtx(ctx)

	r.lit = ctx.Literal().Accept(v).(Literal)

	return r
}

func (v *BuildASTVisitor) VisitParenExp(ctx *parser.ParenExpContext) interface{} {
	r := new(ParenExp)
	r.setPosFromCtx(ctx)

	r.e = ctx.Exp().Accept(v).(Exp)

	return r
}

func (v *BuildASTVisitor) VisitTuple(ctx *parser.TupleContext) interface{} {
	r := new(Tuple)
	r.setPosFromCtx(ctx)

	expTrees := ctx.AllExp()
	for _, tree := range expTrees {
		e := tree.Accept(v).(Exp)
		r.exps = append(r.exps, e)
	}

	return r
}

// alts

func (v *BuildASTVisitor) VisitAlternatives(ctx *parser.AlternativesContext) interface{} {
	r := []*Alternative{}

	altTrees := ctx.AllAlt()

	for _, tree := range altTrees {
		alt := tree.Accept(v).(*Alternative)
		r = append(r, alt)
	}

	return r
}

func (v *BuildASTVisitor) VisitAlternative(ctx *parser.AlternativeContext) interface{} {
	r := new(Alternative)
	r.setPosFromCtx(ctx)

	r.pat = ctx.Pat().Accept(v).(Pat)
	r.exp = ctx.Exp().Accept(v).(Exp)

	return r
}

// pat

func (v *BuildASTVisitor) VisitPatArg(ctx *parser.PatArgContext) interface{} {
	r := new(PatArg)
	r.setPosFromCtx(ctx)

	r.arg = ctx.Apat().Accept(v).(APat)

	return r
}

func (v *BuildASTVisitor) VisitPatCon(ctx *parser.PatConContext) interface{} {
	r := new(PatCon)
	r.setPosFromCtx(ctx)

	r.conId = ctx.CONID().GetText()

	argTrees := ctx.AllApat()
	for _, tree := range argTrees {
		arg := tree.Accept(v).(APat)
		r.args = append(r.args, arg)
	}

	return r
}

// apat

func (v *BuildASTVisitor) VisitApatVar(ctx *parser.ApatVarContext) interface{} {
	r := new(APatVar)
	r.setPosFromCtx(ctx)

	r.id = ctx.VARID().GetText()

	return r
}

func (v *BuildASTVisitor) VisitApatCon(ctx *parser.ApatConContext) interface{} {
	r := new(APatCon)
	r.setPosFromCtx(ctx)

	r.id = ctx.CONID().GetText()

	return r
}

func (v *BuildASTVisitor) VisitApatLit(ctx *parser.ApatLitContext) interface{} {
	r := new(APatLit)
	r.setPosFromCtx(ctx)

	r.lit = ctx.Literal().Accept(v).(Literal)

	return r
}

// literal

func (v *BuildASTVisitor) VisitInt(ctx *parser.IntContext) interface{} {
	e := new(Int)
	e.setPosFromCtx(ctx)

	n, err := strconv.Atoi(ctx.GetText())

	if err == nil {
		e.n = n
	}

	return e
}

func (v *BuildASTVisitor) VisitChar(ctx *parser.CharContext) interface{} {
	r := new(Char)
	r.setPosFromCtx(ctx)

	s := ctx.GetText()
	r.c = s[1]

	return r
}

func (v *BuildASTVisitor) VisitString(ctx *parser.StringContext) interface{} {
	r := new(String)
	r.setPosFromCtx(ctx)

	s := ctx.GetText()
	r.s = strings.TrimPrefix(strings.TrimSuffix(s, "\""), "\"")

	return r
}

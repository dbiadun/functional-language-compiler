package main

import (
	"fmt"
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
	return ctx.Topdecls().Accept(v)
}

// topdecls

func (v *BuildASTVisitor) VisitTopDeclsList(ctx *parser.TopDeclsListContext) interface{} {
	r := new(TopDeclsList)
	r.setPosFromCtx(ctx)
	r.decls = []TopDecl{}

	trees := ctx.AllTopdecl()

	for _, tree := range trees {
		decl := tree.Accept(v).(TopDecl)
		r.decls = append(r.decls, decl)
	}

	return r
}

// topdecl

func (v *BuildASTVisitor) VisitImportTopDecl(ctx *parser.ImportTopDeclContext) interface{} {
	r := new(ImportTopDecl)
	r.setPosFromCtx(ctx)

	s := ctx.STRING().GetText()
	r.file = strings.TrimPrefix(strings.TrimSuffix(s, "\""), "\"")

	return r
}

func (v *BuildASTVisitor) VisitDataTopDecl(ctx *parser.DataTopDeclContext) interface{} {
	r := new(DataTopDecl)
	r.setPosFromCtx(ctx)

	r.t = ctx.Simpletype().Accept(v).(SimpleType)
	r.constrs = ctx.Constrs().Accept(v).(Constrs)

	return r
}

func (v *BuildASTVisitor) VisitFunTopDecl(ctx *parser.FunTopDeclContext) interface{} {
	r := new(FunTopDecl)
	r.setPosFromCtx(ctx)

	r.decl = ctx.Decl().Accept(v).(Decl)

	return r
}

// simpletype

func (v *BuildASTVisitor) VisitDataType(ctx *parser.DataTypeContext) interface{} {
	r := new(DataType)
	r.setPosFromCtx(ctx)

	r.constr = ctx.CONID().GetText()
	r.vars = []string{}

	vars := ctx.AllVARID()

	for _, var_ := range vars {
		r.vars = append(r.vars, var_.GetText())
	}

	return r
}

// constrs

func (v *BuildASTVisitor) VisitConstrList(ctx *parser.ConstrListContext) interface{} {
	r := new(ConstrList)
	r.setPosFromCtx(ctx)

	trees := ctx.AllConstrdef()
	r.defs = []ConstrDef{}

	for _, tree := range trees {
		def := tree.Accept(v).(ConstrDef)
		r.defs = append(r.defs, def)
	}

	return r
}

// constrdef

func (v *BuildASTVisitor) VisitConstrType(ctx *parser.ConstrTypeContext) interface{} {
	r := new(ConstrType)
	r.setPosFromCtx(ctx)

	r.constr = ctx.CONID().GetText()
	r.args = []AType{}

	trees := ctx.AllAtype()

	for _, tree := range trees {
		arg := tree.Accept(v).(AType)
		r.args = append(r.args, arg)
	}

	return r
}

// decl

func (v *BuildASTVisitor) VisitFunTypeDecl(ctx *parser.FunTypeDeclContext) interface{} {
	r := new(FunTypeDecl)
	r.setPosFromCtx(ctx)

	r.d = ctx.Gendecl().Accept(v).(GenDecl)

	return r
}

func (v *BuildASTVisitor) VisitFunDecl(ctx *parser.FunDeclContext) interface{} {
	r := new(FunDecl)
	r.setPosFromCtx(ctx)

	r.lhs = ctx.Funlhs().Accept(v).(FunLhs)
	r.rhs = ctx.Rhs().Accept(v).(Rhs)

	return r
}

func (v *BuildASTVisitor) VisitVarDecl(ctx *parser.VarDeclContext) interface{} {
	r := new(VarDecl)
	r.setPosFromCtx(ctx)

	r.pat = ctx.Pat().Accept(v).(Pat)
	r.rhs = ctx.Rhs().Accept(v).(Rhs)

	return r
}

// gendecl

func (v *BuildASTVisitor) VisitTypeSignature(ctx *parser.TypeSignatureContext) interface{} {
	r := new(TypeSignature)
	r.setPosFromCtx(ctx)

	r.vars = ctx.Vars().Accept(v).([]string)
	r.t = ctx.Type().Accept(v).(Type)

	return r
}

// vars

func (v *BuildASTVisitor) VisitVarList(ctx *parser.VarListContext) interface{} {
	r := []string{}

	vars := ctx.AllVARID()

	for _, v := range vars {
		v := v.GetText()
		r = append(r, v)
	}

	return r
}

// type

func (v *BuildASTVisitor) VisitFunType(ctx *parser.FunTypeContext) interface{} {
	r := new(FunType)
	r.setPosFromCtx(ctx)

	r.types = []BType{}

	trees := ctx.AllBtype()

	for _, tree := range trees {
		bType := tree.Accept(v).(BType)
		r.types = append(r.types, bType)
	}

	return r
}

// btype

func (v *BuildASTVisitor) VisitTypeApp(ctx *parser.TypeAppContext) interface{} {
	r := new(TypeApp)
	r.setPosFromCtx(ctx)

	r.types = []AType{}

	trees := ctx.AllAtype()

	for _, tree := range trees {
		aType := tree.Accept(v).(AType)
		r.types = append(r.types, aType)
	}

	return r
}

// atype

func (v *BuildASTVisitor) VisitConType(ctx *parser.ConTypeContext) interface{} {
	r := new(ConType)
	r.setPosFromCtx(ctx)

	r.id = ctx.CONID().GetText()

	return r
}

func (v *BuildASTVisitor) VisitVarType(ctx *parser.VarTypeContext) interface{} {
	r := new(VarType)
	r.setPosFromCtx(ctx)

	r.id = ctx.VARID().GetText()

	return r
}

func (v *BuildASTVisitor) VisitTupleType(ctx *parser.TupleTypeContext) interface{} {
	r := new(TupleType)
	r.setPosFromCtx(ctx)

	r.types = []Type{}

	trees := ctx.AllType()

	for _, tree := range trees {
		t := tree.Accept(v).(Type)
		r.types = append(r.types, t)
	}

	return r
}

func (v *BuildASTVisitor) VisitParenType(ctx *parser.ParenTypeContext) interface{} {
	r := new(ParenType)
	r.setPosFromCtx(ctx)

	r.t = ctx.Type().Accept(v).(Type)

	return r
}

// funlhs

func (v *BuildASTVisitor) VisitDeclLhs(ctx *parser.DeclLhsContext) interface{} {
	r := new(DeclLhs)
	r.setPosFromCtx(ctx)

	r.fun = ctx.VARID().GetText()
	r.args = []APat{}

	trees := ctx.AllApat()

	for _, tree := range trees {
		arg := tree.Accept(v).(APat)
		r.args = append(r.args, arg)
	}

	return r
}

// rhs

func (v *BuildASTVisitor) VisitDeclExp(ctx *parser.DeclExpContext) interface{} {
	r := new(DeclExp)
	r.setPosFromCtx(ctx)

	r.e = ctx.Exp().Accept(v).(Exp)

	return r
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

func (v *BuildASTVisitor) VisitEComp(ctx *parser.ECompContext) interface{} {
	e := new(EComp)

	subtree1 := ctx.GetE1()
	subtree2 := ctx.GetE2()
	line := ctx.GetOp().GetLine()
	col := ctx.GetOp().GetColumn()

	e.setPos(line, col)

	e.e1 = subtree1.Accept(v).(Exp)
	e.e2 = subtree2.Accept(v).(Exp)
	e.op = ctx.GetOp().GetText()

	return e
}

func (v *BuildASTVisitor) VisitELogical(ctx *parser.ELogicalContext) interface{} {
	e := new(ELogical)

	subtree1 := ctx.GetE1()
	subtree2 := ctx.GetE2()
	line := ctx.GetOp().GetLine()
	col := ctx.GetOp().GetColumn()

	e.setPos(line, col)

	e.e1 = subtree1.Accept(v).(Exp)
	e.e2 = subtree2.Accept(v).(Exp)
	e.op = ctx.GetOp().GetText()

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
	} else {
		log.Fatalln(fmt.Sprintf("Error: '%s' is not a permitted integer", ctx.GetText()))
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

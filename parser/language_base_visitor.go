// Code generated from Language.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Language

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseLanguageVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLanguageVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitTopDeclsList(ctx *TopDeclsListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitImportTopDecl(ctx *ImportTopDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitDataTopDecl(ctx *DataTopDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitFunTopDecl(ctx *FunTopDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitDataType(ctx *DataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitConstrList(ctx *ConstrListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitConstrType(ctx *ConstrTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitFunTypeDecl(ctx *FunTypeDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitFunDecl(ctx *FunDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitVarDecl(ctx *VarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitTypeSignature(ctx *TypeSignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitVarList(ctx *VarListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitFunType(ctx *FunTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitTypeApp(ctx *TypeAppContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitConType(ctx *ConTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitVarType(ctx *VarTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitTupleType(ctx *TupleTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitParenType(ctx *ParenTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitUnitType(ctx *UnitTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitDeclLhs(ctx *DeclLhsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitDeclExp(ctx *DeclExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitEDo(ctx *EDoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitEMulDiv(ctx *EMulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitELogical(ctx *ELogicalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitEComp(ctx *ECompContext) interface{} {
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

func (v *BaseLanguageVisitor) VisitStmtsList(ctx *StmtsListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitSExp(ctx *SExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitSAssign(ctx *SAssignContext) interface{} {
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

func (v *BaseLanguageVisitor) VisitDec(ctx *DecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitHex(ctx *HexContext) interface{} {
	return v.VisitChildren(ctx)
}

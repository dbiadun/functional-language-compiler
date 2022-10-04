// Code generated from Language.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Language

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LanguageListener is a complete listener for a parse tree produced by LanguageParser.
type LanguageListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterTopDeclsList is called when entering the TopDeclsList production.
	EnterTopDeclsList(c *TopDeclsListContext)

	// EnterImportTopDecl is called when entering the ImportTopDecl production.
	EnterImportTopDecl(c *ImportTopDeclContext)

	// EnterDataTopDecl is called when entering the DataTopDecl production.
	EnterDataTopDecl(c *DataTopDeclContext)

	// EnterFunTopDecl is called when entering the FunTopDecl production.
	EnterFunTopDecl(c *FunTopDeclContext)

	// EnterDataType is called when entering the DataType production.
	EnterDataType(c *DataTypeContext)

	// EnterConstrList is called when entering the ConstrList production.
	EnterConstrList(c *ConstrListContext)

	// EnterConstrType is called when entering the ConstrType production.
	EnterConstrType(c *ConstrTypeContext)

	// EnterFunTypeDecl is called when entering the FunTypeDecl production.
	EnterFunTypeDecl(c *FunTypeDeclContext)

	// EnterFunDecl is called when entering the FunDecl production.
	EnterFunDecl(c *FunDeclContext)

	// EnterVarDecl is called when entering the VarDecl production.
	EnterVarDecl(c *VarDeclContext)

	// EnterTypeSignature is called when entering the TypeSignature production.
	EnterTypeSignature(c *TypeSignatureContext)

	// EnterVarList is called when entering the VarList production.
	EnterVarList(c *VarListContext)

	// EnterFunType is called when entering the FunType production.
	EnterFunType(c *FunTypeContext)

	// EnterTypeApp is called when entering the TypeApp production.
	EnterTypeApp(c *TypeAppContext)

	// EnterConType is called when entering the ConType production.
	EnterConType(c *ConTypeContext)

	// EnterVarType is called when entering the VarType production.
	EnterVarType(c *VarTypeContext)

	// EnterTupleType is called when entering the TupleType production.
	EnterTupleType(c *TupleTypeContext)

	// EnterParenType is called when entering the ParenType production.
	EnterParenType(c *ParenTypeContext)

	// EnterUnitType is called when entering the UnitType production.
	EnterUnitType(c *UnitTypeContext)

	// EnterDeclLhs is called when entering the DeclLhs production.
	EnterDeclLhs(c *DeclLhsContext)

	// EnterDeclExp is called when entering the DeclExp production.
	EnterDeclExp(c *DeclExpContext)

	// EnterEDo is called when entering the EDo production.
	EnterEDo(c *EDoContext)

	// EnterEMulDiv is called when entering the EMulDiv production.
	EnterEMulDiv(c *EMulDivContext)

	// EnterELogical is called when entering the ELogical production.
	EnterELogical(c *ELogicalContext)

	// EnterEComp is called when entering the EComp production.
	EnterEComp(c *ECompContext)

	// EnterECase is called when entering the ECase production.
	EnterECase(c *ECaseContext)

	// EnterEFun is called when entering the EFun production.
	EnterEFun(c *EFunContext)

	// EnterEBitOr is called when entering the EBitOr production.
	EnterEBitOr(c *EBitOrContext)

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

	// EnterStmtsList is called when entering the StmtsList production.
	EnterStmtsList(c *StmtsListContext)

	// EnterSExp is called when entering the SExp production.
	EnterSExp(c *SExpContext)

	// EnterSAssign is called when entering the SAssign production.
	EnterSAssign(c *SAssignContext)

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

	// EnterDec is called when entering the Dec production.
	EnterDec(c *DecContext)

	// EnterHex is called when entering the Hex production.
	EnterHex(c *HexContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitTopDeclsList is called when exiting the TopDeclsList production.
	ExitTopDeclsList(c *TopDeclsListContext)

	// ExitImportTopDecl is called when exiting the ImportTopDecl production.
	ExitImportTopDecl(c *ImportTopDeclContext)

	// ExitDataTopDecl is called when exiting the DataTopDecl production.
	ExitDataTopDecl(c *DataTopDeclContext)

	// ExitFunTopDecl is called when exiting the FunTopDecl production.
	ExitFunTopDecl(c *FunTopDeclContext)

	// ExitDataType is called when exiting the DataType production.
	ExitDataType(c *DataTypeContext)

	// ExitConstrList is called when exiting the ConstrList production.
	ExitConstrList(c *ConstrListContext)

	// ExitConstrType is called when exiting the ConstrType production.
	ExitConstrType(c *ConstrTypeContext)

	// ExitFunTypeDecl is called when exiting the FunTypeDecl production.
	ExitFunTypeDecl(c *FunTypeDeclContext)

	// ExitFunDecl is called when exiting the FunDecl production.
	ExitFunDecl(c *FunDeclContext)

	// ExitVarDecl is called when exiting the VarDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitTypeSignature is called when exiting the TypeSignature production.
	ExitTypeSignature(c *TypeSignatureContext)

	// ExitVarList is called when exiting the VarList production.
	ExitVarList(c *VarListContext)

	// ExitFunType is called when exiting the FunType production.
	ExitFunType(c *FunTypeContext)

	// ExitTypeApp is called when exiting the TypeApp production.
	ExitTypeApp(c *TypeAppContext)

	// ExitConType is called when exiting the ConType production.
	ExitConType(c *ConTypeContext)

	// ExitVarType is called when exiting the VarType production.
	ExitVarType(c *VarTypeContext)

	// ExitTupleType is called when exiting the TupleType production.
	ExitTupleType(c *TupleTypeContext)

	// ExitParenType is called when exiting the ParenType production.
	ExitParenType(c *ParenTypeContext)

	// ExitUnitType is called when exiting the UnitType production.
	ExitUnitType(c *UnitTypeContext)

	// ExitDeclLhs is called when exiting the DeclLhs production.
	ExitDeclLhs(c *DeclLhsContext)

	// ExitDeclExp is called when exiting the DeclExp production.
	ExitDeclExp(c *DeclExpContext)

	// ExitEDo is called when exiting the EDo production.
	ExitEDo(c *EDoContext)

	// ExitEMulDiv is called when exiting the EMulDiv production.
	ExitEMulDiv(c *EMulDivContext)

	// ExitELogical is called when exiting the ELogical production.
	ExitELogical(c *ELogicalContext)

	// ExitEComp is called when exiting the EComp production.
	ExitEComp(c *ECompContext)

	// ExitECase is called when exiting the ECase production.
	ExitECase(c *ECaseContext)

	// ExitEFun is called when exiting the EFun production.
	ExitEFun(c *EFunContext)

	// ExitEBitOr is called when exiting the EBitOr production.
	ExitEBitOr(c *EBitOrContext)

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

	// ExitStmtsList is called when exiting the StmtsList production.
	ExitStmtsList(c *StmtsListContext)

	// ExitSExp is called when exiting the SExp production.
	ExitSExp(c *SExpContext)

	// ExitSAssign is called when exiting the SAssign production.
	ExitSAssign(c *SAssignContext)

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

	// ExitDec is called when exiting the Dec production.
	ExitDec(c *DecContext)

	// ExitHex is called when exiting the Hex production.
	ExitHex(c *HexContext)
}

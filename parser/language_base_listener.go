// Code generated from Language.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Language

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLanguageListener is a complete listener for a parse tree produced by LanguageParser.
type BaseLanguageListener struct{}

var _ LanguageListener = &BaseLanguageListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLanguageListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLanguageListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLanguageListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLanguageListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseLanguageListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseLanguageListener) ExitStart(ctx *StartContext) {}

// EnterTopDeclsList is called when production TopDeclsList is entered.
func (s *BaseLanguageListener) EnterTopDeclsList(ctx *TopDeclsListContext) {}

// ExitTopDeclsList is called when production TopDeclsList is exited.
func (s *BaseLanguageListener) ExitTopDeclsList(ctx *TopDeclsListContext) {}

// EnterImportTopDecl is called when production ImportTopDecl is entered.
func (s *BaseLanguageListener) EnterImportTopDecl(ctx *ImportTopDeclContext) {}

// ExitImportTopDecl is called when production ImportTopDecl is exited.
func (s *BaseLanguageListener) ExitImportTopDecl(ctx *ImportTopDeclContext) {}

// EnterDataTopDecl is called when production DataTopDecl is entered.
func (s *BaseLanguageListener) EnterDataTopDecl(ctx *DataTopDeclContext) {}

// ExitDataTopDecl is called when production DataTopDecl is exited.
func (s *BaseLanguageListener) ExitDataTopDecl(ctx *DataTopDeclContext) {}

// EnterFunTopDecl is called when production FunTopDecl is entered.
func (s *BaseLanguageListener) EnterFunTopDecl(ctx *FunTopDeclContext) {}

// ExitFunTopDecl is called when production FunTopDecl is exited.
func (s *BaseLanguageListener) ExitFunTopDecl(ctx *FunTopDeclContext) {}

// EnterDataType is called when production DataType is entered.
func (s *BaseLanguageListener) EnterDataType(ctx *DataTypeContext) {}

// ExitDataType is called when production DataType is exited.
func (s *BaseLanguageListener) ExitDataType(ctx *DataTypeContext) {}

// EnterConstrList is called when production ConstrList is entered.
func (s *BaseLanguageListener) EnterConstrList(ctx *ConstrListContext) {}

// ExitConstrList is called when production ConstrList is exited.
func (s *BaseLanguageListener) ExitConstrList(ctx *ConstrListContext) {}

// EnterConstrType is called when production ConstrType is entered.
func (s *BaseLanguageListener) EnterConstrType(ctx *ConstrTypeContext) {}

// ExitConstrType is called when production ConstrType is exited.
func (s *BaseLanguageListener) ExitConstrType(ctx *ConstrTypeContext) {}

// EnterFunTypeDecl is called when production FunTypeDecl is entered.
func (s *BaseLanguageListener) EnterFunTypeDecl(ctx *FunTypeDeclContext) {}

// ExitFunTypeDecl is called when production FunTypeDecl is exited.
func (s *BaseLanguageListener) ExitFunTypeDecl(ctx *FunTypeDeclContext) {}

// EnterFunDecl is called when production FunDecl is entered.
func (s *BaseLanguageListener) EnterFunDecl(ctx *FunDeclContext) {}

// ExitFunDecl is called when production FunDecl is exited.
func (s *BaseLanguageListener) ExitFunDecl(ctx *FunDeclContext) {}

// EnterVarDecl is called when production VarDecl is entered.
func (s *BaseLanguageListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production VarDecl is exited.
func (s *BaseLanguageListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterTypeSignature is called when production TypeSignature is entered.
func (s *BaseLanguageListener) EnterTypeSignature(ctx *TypeSignatureContext) {}

// ExitTypeSignature is called when production TypeSignature is exited.
func (s *BaseLanguageListener) ExitTypeSignature(ctx *TypeSignatureContext) {}

// EnterVarList is called when production VarList is entered.
func (s *BaseLanguageListener) EnterVarList(ctx *VarListContext) {}

// ExitVarList is called when production VarList is exited.
func (s *BaseLanguageListener) ExitVarList(ctx *VarListContext) {}

// EnterFunType is called when production FunType is entered.
func (s *BaseLanguageListener) EnterFunType(ctx *FunTypeContext) {}

// ExitFunType is called when production FunType is exited.
func (s *BaseLanguageListener) ExitFunType(ctx *FunTypeContext) {}

// EnterTypeApp is called when production TypeApp is entered.
func (s *BaseLanguageListener) EnterTypeApp(ctx *TypeAppContext) {}

// ExitTypeApp is called when production TypeApp is exited.
func (s *BaseLanguageListener) ExitTypeApp(ctx *TypeAppContext) {}

// EnterConType is called when production ConType is entered.
func (s *BaseLanguageListener) EnterConType(ctx *ConTypeContext) {}

// ExitConType is called when production ConType is exited.
func (s *BaseLanguageListener) ExitConType(ctx *ConTypeContext) {}

// EnterVarType is called when production VarType is entered.
func (s *BaseLanguageListener) EnterVarType(ctx *VarTypeContext) {}

// ExitVarType is called when production VarType is exited.
func (s *BaseLanguageListener) ExitVarType(ctx *VarTypeContext) {}

// EnterTupleType is called when production TupleType is entered.
func (s *BaseLanguageListener) EnterTupleType(ctx *TupleTypeContext) {}

// ExitTupleType is called when production TupleType is exited.
func (s *BaseLanguageListener) ExitTupleType(ctx *TupleTypeContext) {}

// EnterParenType is called when production ParenType is entered.
func (s *BaseLanguageListener) EnterParenType(ctx *ParenTypeContext) {}

// ExitParenType is called when production ParenType is exited.
func (s *BaseLanguageListener) ExitParenType(ctx *ParenTypeContext) {}

// EnterUnitType is called when production UnitType is entered.
func (s *BaseLanguageListener) EnterUnitType(ctx *UnitTypeContext) {}

// ExitUnitType is called when production UnitType is exited.
func (s *BaseLanguageListener) ExitUnitType(ctx *UnitTypeContext) {}

// EnterDeclLhs is called when production DeclLhs is entered.
func (s *BaseLanguageListener) EnterDeclLhs(ctx *DeclLhsContext) {}

// ExitDeclLhs is called when production DeclLhs is exited.
func (s *BaseLanguageListener) ExitDeclLhs(ctx *DeclLhsContext) {}

// EnterDeclExp is called when production DeclExp is entered.
func (s *BaseLanguageListener) EnterDeclExp(ctx *DeclExpContext) {}

// ExitDeclExp is called when production DeclExp is exited.
func (s *BaseLanguageListener) ExitDeclExp(ctx *DeclExpContext) {}

// EnterEDo is called when production EDo is entered.
func (s *BaseLanguageListener) EnterEDo(ctx *EDoContext) {}

// ExitEDo is called when production EDo is exited.
func (s *BaseLanguageListener) ExitEDo(ctx *EDoContext) {}

// EnterEMulDiv is called when production EMulDiv is entered.
func (s *BaseLanguageListener) EnterEMulDiv(ctx *EMulDivContext) {}

// ExitEMulDiv is called when production EMulDiv is exited.
func (s *BaseLanguageListener) ExitEMulDiv(ctx *EMulDivContext) {}

// EnterELogical is called when production ELogical is entered.
func (s *BaseLanguageListener) EnterELogical(ctx *ELogicalContext) {}

// ExitELogical is called when production ELogical is exited.
func (s *BaseLanguageListener) ExitELogical(ctx *ELogicalContext) {}

// EnterEComp is called when production EComp is entered.
func (s *BaseLanguageListener) EnterEComp(ctx *ECompContext) {}

// ExitEComp is called when production EComp is exited.
func (s *BaseLanguageListener) ExitEComp(ctx *ECompContext) {}

// EnterECase is called when production ECase is entered.
func (s *BaseLanguageListener) EnterECase(ctx *ECaseContext) {}

// ExitECase is called when production ECase is exited.
func (s *BaseLanguageListener) ExitECase(ctx *ECaseContext) {}

// EnterEFun is called when production EFun is entered.
func (s *BaseLanguageListener) EnterEFun(ctx *EFunContext) {}

// ExitEFun is called when production EFun is exited.
func (s *BaseLanguageListener) ExitEFun(ctx *EFunContext) {}

// EnterEBitOr is called when production EBitOr is entered.
func (s *BaseLanguageListener) EnterEBitOr(ctx *EBitOrContext) {}

// ExitEBitOr is called when production EBitOr is exited.
func (s *BaseLanguageListener) ExitEBitOr(ctx *EBitOrContext) {}

// EnterEAddSub is called when production EAddSub is entered.
func (s *BaseLanguageListener) EnterEAddSub(ctx *EAddSubContext) {}

// ExitEAddSub is called when production EAddSub is exited.
func (s *BaseLanguageListener) ExitEAddSub(ctx *EAddSubContext) {}

// EnterFArg is called when production FArg is entered.
func (s *BaseLanguageListener) EnterFArg(ctx *FArgContext) {}

// ExitFArg is called when production FArg is exited.
func (s *BaseLanguageListener) ExitFArg(ctx *FArgContext) {}

// EnterFApp is called when production FApp is entered.
func (s *BaseLanguageListener) EnterFApp(ctx *FAppContext) {}

// ExitFApp is called when production FApp is exited.
func (s *BaseLanguageListener) ExitFApp(ctx *FAppContext) {}

// EnterVar is called when production Var is entered.
func (s *BaseLanguageListener) EnterVar(ctx *VarContext) {}

// ExitVar is called when production Var is exited.
func (s *BaseLanguageListener) ExitVar(ctx *VarContext) {}

// EnterConstr is called when production Constr is entered.
func (s *BaseLanguageListener) EnterConstr(ctx *ConstrContext) {}

// ExitConstr is called when production Constr is exited.
func (s *BaseLanguageListener) ExitConstr(ctx *ConstrContext) {}

// EnterLit is called when production Lit is entered.
func (s *BaseLanguageListener) EnterLit(ctx *LitContext) {}

// ExitLit is called when production Lit is exited.
func (s *BaseLanguageListener) ExitLit(ctx *LitContext) {}

// EnterParenExp is called when production ParenExp is entered.
func (s *BaseLanguageListener) EnterParenExp(ctx *ParenExpContext) {}

// ExitParenExp is called when production ParenExp is exited.
func (s *BaseLanguageListener) ExitParenExp(ctx *ParenExpContext) {}

// EnterTuple is called when production Tuple is entered.
func (s *BaseLanguageListener) EnterTuple(ctx *TupleContext) {}

// ExitTuple is called when production Tuple is exited.
func (s *BaseLanguageListener) ExitTuple(ctx *TupleContext) {}

// EnterStmtsList is called when production StmtsList is entered.
func (s *BaseLanguageListener) EnterStmtsList(ctx *StmtsListContext) {}

// ExitStmtsList is called when production StmtsList is exited.
func (s *BaseLanguageListener) ExitStmtsList(ctx *StmtsListContext) {}

// EnterSExp is called when production SExp is entered.
func (s *BaseLanguageListener) EnterSExp(ctx *SExpContext) {}

// ExitSExp is called when production SExp is exited.
func (s *BaseLanguageListener) ExitSExp(ctx *SExpContext) {}

// EnterSAssign is called when production SAssign is entered.
func (s *BaseLanguageListener) EnterSAssign(ctx *SAssignContext) {}

// ExitSAssign is called when production SAssign is exited.
func (s *BaseLanguageListener) ExitSAssign(ctx *SAssignContext) {}

// EnterAlternatives is called when production Alternatives is entered.
func (s *BaseLanguageListener) EnterAlternatives(ctx *AlternativesContext) {}

// ExitAlternatives is called when production Alternatives is exited.
func (s *BaseLanguageListener) ExitAlternatives(ctx *AlternativesContext) {}

// EnterAlternative is called when production Alternative is entered.
func (s *BaseLanguageListener) EnterAlternative(ctx *AlternativeContext) {}

// ExitAlternative is called when production Alternative is exited.
func (s *BaseLanguageListener) ExitAlternative(ctx *AlternativeContext) {}

// EnterPatArg is called when production PatArg is entered.
func (s *BaseLanguageListener) EnterPatArg(ctx *PatArgContext) {}

// ExitPatArg is called when production PatArg is exited.
func (s *BaseLanguageListener) ExitPatArg(ctx *PatArgContext) {}

// EnterPatCon is called when production PatCon is entered.
func (s *BaseLanguageListener) EnterPatCon(ctx *PatConContext) {}

// ExitPatCon is called when production PatCon is exited.
func (s *BaseLanguageListener) ExitPatCon(ctx *PatConContext) {}

// EnterApatVar is called when production ApatVar is entered.
func (s *BaseLanguageListener) EnterApatVar(ctx *ApatVarContext) {}

// ExitApatVar is called when production ApatVar is exited.
func (s *BaseLanguageListener) ExitApatVar(ctx *ApatVarContext) {}

// EnterApatCon is called when production ApatCon is entered.
func (s *BaseLanguageListener) EnterApatCon(ctx *ApatConContext) {}

// ExitApatCon is called when production ApatCon is exited.
func (s *BaseLanguageListener) ExitApatCon(ctx *ApatConContext) {}

// EnterApatLit is called when production ApatLit is entered.
func (s *BaseLanguageListener) EnterApatLit(ctx *ApatLitContext) {}

// ExitApatLit is called when production ApatLit is exited.
func (s *BaseLanguageListener) ExitApatLit(ctx *ApatLitContext) {}

// EnterInt is called when production Int is entered.
func (s *BaseLanguageListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production Int is exited.
func (s *BaseLanguageListener) ExitInt(ctx *IntContext) {}

// EnterChar is called when production Char is entered.
func (s *BaseLanguageListener) EnterChar(ctx *CharContext) {}

// ExitChar is called when production Char is exited.
func (s *BaseLanguageListener) ExitChar(ctx *CharContext) {}

// EnterString is called when production String is entered.
func (s *BaseLanguageListener) EnterString(ctx *StringContext) {}

// ExitString is called when production String is exited.
func (s *BaseLanguageListener) ExitString(ctx *StringContext) {}

// EnterDec is called when production Dec is entered.
func (s *BaseLanguageListener) EnterDec(ctx *DecContext) {}

// ExitDec is called when production Dec is exited.
func (s *BaseLanguageListener) ExitDec(ctx *DecContext) {}

// EnterHex is called when production Hex is entered.
func (s *BaseLanguageListener) EnterHex(ctx *HexContext) {}

// ExitHex is called when production Hex is exited.
func (s *BaseLanguageListener) ExitHex(ctx *HexContext) {}

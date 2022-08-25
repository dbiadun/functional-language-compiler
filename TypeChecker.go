package main

import "log"

type TypeChecker struct {
	types        map[string]*FunType    // types of used functions and variables
	constrs      map[string]*ConstrInfo // all defined constructors
	definedTypes map[string]*DataType   // all defined types
}

type ConstrInfo struct {
	t       *DataType
	constrT *ConstrType
}

type TypeCheckResult struct {
}

func (*TypeChecker) errFatal(v ASTNode, s string) {
	line, col := v.getPos()
	log.Fatalf("TypeChecker error at line %d, column %d: %s\n", line, col, s)
}

func (c *TypeChecker) checkTopDecls(v TopDecls) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkTopDeclsList(v *TopDeclsList) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkTopDecl(v TopDecl) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkDataTopDecl(v *DataTopDecl) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkFunTopDecl(v *FunTopDecl) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkSimpleType(v SimpleType) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkDataType(v *DataType) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkConstrs(v Constrs) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkConstrList(v *ConstrList) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkConstrDef(v ConstrDef) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkConstrType(v *ConstrType) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkDecl(v Decl) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkFunTypeDecl(v *FunTypeDecl) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkFunDecl(v *FunDecl) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkVarDecl(v *VarDecl) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkGenDecl(v GenDecl) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkTypeSignature(v *TypeSignature) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkType(v Type) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkFunType(v *FunType) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkBType(v BType) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkTypeApp(v *TypeApp) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkAType(v AType) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkConType(v *ConType) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkVarType(v *VarType) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkTupleType(v *TupleType) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkParenType(v *ParenType) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkFunLhs(v FunLhs) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkDeclLhs(v *DeclLhs) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkRhs(v Rhs) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkDeclExp(v *DeclExp) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkExp(v Exp) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkEFun(v *EFun) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkECase(v *ECase) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkEMul(v *EMul) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkEDiv(v *EDiv) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkEAdd(v *EAdd) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkESub(v *ESub) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkFExp(v FExp) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkFApp(v *FApp) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkFArg(v *FArg) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkAExp(v AExp) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkVar(v *Var) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkConstr(v *Constr) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkLit(v *Lit) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkParenExp(v *ParenExp) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkTuple(v *Tuple) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkAlternative(v *Alternative) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkPat(v Pat) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkPatArg(v *PatArg) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkPatCon(v *PatCon) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkAPat(v APat) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkAPatVar(v *APatVar) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkAPatCon(v *APatCon) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkAPatLit(v *APatLit) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkLiteral(v Literal) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkInt(v *Int) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkChar(v *Char) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkString(v *String) *TypeCheckResult {
	return &TypeCheckResult{}
}

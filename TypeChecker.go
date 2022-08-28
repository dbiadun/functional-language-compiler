package main

import (
	"fmt"
	"log"
)

////////////////////////////////////// CONTEXT ///////////////////////////////////////

type ConstrInfo struct {
	t       *DataType
	constrT *ConstrType
}

type ContextData interface {
	contextData()
}

func (*DataType) contextData() {}

type void struct{}

func (void) contextData() {}

var voidMember void

type DataTypeVars struct {
	vars map[string]void
}

func (*DataTypeVars) contextData() {}

type ParameterTypes struct {
	types map[string]void
}

func (*ParameterTypes) contextData() {}

const (
	CurDataType       = iota // *DataType
	CurDataTypeVars          // *DataTypeVars
	CurParameterTypes        // *ParameterTypes
	LastArrowType            // void
	TypesCheck               // void
	GlobalNamesCheck         // void
	LastCheck                // void
)

/////////////////////////////////// TYPE_CHECKER /////////////////////////////////////

type TypeChecker struct {
	types        map[string]*FunType    // types of declared functions and variables
	constrs      map[string]*ConstrInfo // all defined constructors
	definedTypes map[string]*DataType   // all defined types
	context      map[int]ContextData    // additional context
}

type TypeCheckResult struct {
	t *FunType
}

func (c *TypeChecker) init() {
	c.types = make(map[string]*FunType)
	c.constrs = make(map[string]*ConstrInfo)
	c.definedTypes = make(map[string]*DataType)
	c.context = make(map[int]ContextData)
}

func (*TypeChecker) errFatal(v ASTNode, s string) {
	line, col := v.getPos()
	log.Fatalf("TypeChecker error at line %d, column %d: %s\n", line, col, s)
}

// The actual check

func (c *TypeChecker) checkTopDecls(v TopDecls) *TypeCheckResult {
	switch v := v.(type) {
	case *TopDeclsList:
		c.context[TypesCheck] = voidMember
		c.checkTopDeclsList(v)
		delete(c.context, TypesCheck)

		c.context[GlobalNamesCheck] = voidMember
		c.checkTopDeclsList(v)
		delete(c.context, GlobalNamesCheck)

		c.context[LastCheck] = voidMember
		c.checkTopDeclsList(v)
		delete(c.context, LastCheck)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkTopDeclsList(v *TopDeclsList) *TypeCheckResult {
	for _, decl := range v.decls {
		c.checkTopDecl(decl)
	}

	return &TypeCheckResult{}
}

func (c *TypeChecker) checkTopDecl(v TopDecl) *TypeCheckResult {
	switch v := v.(type) {
	case *DataTopDecl:
		return c.checkDataTopDecl(v)
	case *FunTopDecl:
		return c.checkFunTopDecl(v)
	}

	return &TypeCheckResult{}
}

func (c *TypeChecker) checkDataTopDecl(v *DataTopDecl) *TypeCheckResult {
	c.checkSimpleType(v.t)

	_, globalNamesCheck := c.context[GlobalNamesCheck]
	if globalNamesCheck {
		c.checkConstrs(v.constrs)
	}

	return &TypeCheckResult{}
}

func (c *TypeChecker) checkFunTopDecl(v *FunTopDecl) *TypeCheckResult {
	return c.checkDecl(v.decl)
}

func (c *TypeChecker) checkSimpleType(v SimpleType) *TypeCheckResult {
	switch v := v.(type) {

	// Should place a *DataType at CurDataType field of context
	// and a *DataTypeVars at CurDataTypeVars field

	case *DataType:
		return c.checkDataType(v)
	}

	return &TypeCheckResult{}
}

func (c *TypeChecker) checkDataType(v *DataType) *TypeCheckResult {
	_, typesCheck := c.context[TypesCheck]
	_, globalNamesCheck := c.context[GlobalNamesCheck]

	if typesCheck {
		_, ok := c.definedTypes[v.constr]
		if ok {
			c.errFatal(v, fmt.Sprintf("Type '%s' already defined", v.constr))
		}
		c.definedTypes[v.constr] = v
	}

	if globalNamesCheck {
		c.context[CurDataType] = v

		vars := new(DataTypeVars)
		vars.vars = make(map[string]void)
		for _, curVar := range v.vars {
			_, ok := vars.vars[curVar]
			if ok {
				c.errFatal(v, fmt.Sprintf("Conflicting definitions for '%s'", curVar))
			}
			vars.vars[curVar] = voidMember
		}
		c.context[CurDataTypeVars] = vars
	}

	return &TypeCheckResult{}
}

func (c *TypeChecker) checkConstrs(v Constrs) *TypeCheckResult {
	switch v := v.(type) {
	case *ConstrList:
		return c.checkConstrList(v)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkConstrList(v *ConstrList) *TypeCheckResult {
	for _, def := range v.defs {
		c.checkConstrDef(def)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkConstrDef(v ConstrDef) *TypeCheckResult {
	switch v := v.(type) {
	case *ConstrType:
		return c.checkConstrType(v)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkConstrType(v *ConstrType) *TypeCheckResult {
	_, ok := c.constrs[v.constr]
	if ok {
		c.errFatal(v, fmt.Sprintf("Constructor '%s' already defined", v.constr))
	}

	for _, arg := range v.args {
		// argsCount is equal to 0, because type with parameters would be enclosed in parentheses
		c.checkConstrDefAType(arg, 0)
	}

	c.constrs[v.constr] = &ConstrInfo{c.context[CurDataType].(*DataType), v}

	return &TypeCheckResult{}
}

func (c *TypeChecker) checkConstrDefType(v Type) *TypeCheckResult {
	switch v := v.(type) {
	case *FunType:
		c.checkConstrDefFunType(v)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkConstrDefFunType(v *FunType) *TypeCheckResult {
	for _, t := range v.types {
		c.checkConstrDefBType(t)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkConstrDefBType(v BType) *TypeCheckResult {
	switch v := v.(type) {
	case *TypeApp:
		c.checkConstrDefTypeApp(v)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkConstrDefTypeApp(v *TypeApp) *TypeCheckResult {
	for i, t := range v.types {
		argsCount := 0
		if i == 0 {
			// The first type has parameters (the rest of the types), so we need to check if its arity is valid
			argsCount = len(v.types) - 1
		}
		c.checkConstrDefAType(t, argsCount)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkConstrDefAType(v AType, argsCount int) *TypeCheckResult {
	switch v := v.(type) {
	case *ConType:
		c.checkConstrDefConType(v, argsCount)
	case *VarType:
		c.checkConstrDefVarType(v, argsCount)
	case *TupleType:
		c.checkConstrDefTupleType(v, argsCount)
	case *ParenType:
		c.checkConstrDefParenType(v, argsCount)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkConstrDefConType(v *ConType, argsCount int) *TypeCheckResult {
	c.checkDataTypeExistence(v, v.id, argsCount)
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkConstrDefVarType(v *VarType, argsCount int) *TypeCheckResult {
	if argsCount > 0 {
		c.errFatal(v, fmt.Sprintf("Type '%s' can't have any parameters", v.id))
	}

	vars, ok := c.context[CurDataTypeVars]

	if ok {
		varsSet := vars.(*DataTypeVars).vars
		_, ok = varsSet[v.id]
		if ok {
			return &TypeCheckResult{}
		}
	}

	c.errFatal(v, fmt.Sprintf("Type '%s' not defined", v.id))

	return &TypeCheckResult{}
}

func (c *TypeChecker) checkConstrDefTupleType(v *TupleType, argsCount int) *TypeCheckResult {
	if argsCount > 0 {
		c.errFatal(v, "Tuple type can't have any parameters")
	}

	for _, t := range v.types {
		c.checkConstrDefType(t)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkConstrDefParenType(v *ParenType, argsCount int) *TypeCheckResult {
	if argsCount > 0 {
		c.errFatal(v, "Parenthesis type can't have any parameters")
	}

	c.checkConstrDefType(v.t)
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkDecl(v Decl) *TypeCheckResult {
	switch v := v.(type) {
	case *FunTypeDecl:
		c.checkFunTypeDecl(v)
	case *FunDecl:
		c.checkFunDecl(v)
	case *VarDecl:
		c.checkVarDecl(v)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkFunTypeDecl(v *FunTypeDecl) *TypeCheckResult {
	_, globalNamesCheck := c.context[GlobalNamesCheck]

	if !globalNamesCheck {
		return &TypeCheckResult{}
	}

	c.checkGenDecl(v.d)
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkFunDecl(v *FunDecl) *TypeCheckResult {
	_, globalNamesCheck := c.context[GlobalNamesCheck]

	if !globalNamesCheck {
		return &TypeCheckResult{}
	}

	return &TypeCheckResult{}
}

func (c *TypeChecker) checkVarDecl(v *VarDecl) *TypeCheckResult {
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkGenDecl(v GenDecl) *TypeCheckResult {
	switch v := v.(type) {
	case *TypeSignature:
		c.checkTypeSignature(v)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkTypeSignature(v *TypeSignature) *TypeCheckResult {
	c.context[CurParameterTypes] = &ParameterTypes{make(map[string]void)}
	c.context[LastArrowType] = voidMember
	t := c.checkType(v.t).t

	for _, curVar := range v.vars {
		_, ok := c.types[curVar]
		if ok {
			c.errFatal(v, fmt.Sprintf("Type of '%s' already defined", curVar))
		}

		c.types[curVar] = t
	}

	return &TypeCheckResult{}
}

func (c *TypeChecker) checkType(v Type) *TypeCheckResult {
	switch v := v.(type) {
	case *FunType:
		return c.checkFunType(v)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkFunType(v *FunType) *TypeCheckResult {
	_, lastArrowType := c.context[LastArrowType]
	delete(c.context, LastArrowType)

	for i, t := range v.types {
		// We need to know if the type is after the last arrow to know if we need its parameter types defined
		// We don't set LastArrowType on the last type always, because we want to guarantee that
		// only the real last type in the whole type tree is set to last
		if lastArrowType && i == len(v.types)-1 {
			c.context[LastArrowType] = voidMember
		}
		c.checkBType(t)
	}
	return &TypeCheckResult{v}
}

func (c *TypeChecker) checkBType(v BType) *TypeCheckResult {
	switch v := v.(type) {
	case *TypeApp:
		c.checkTypeApp(v)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkTypeApp(v *TypeApp) *TypeCheckResult {
	for i, t := range v.types {
		argsCount := 0
		if i == 0 {
			argsCount = len(v.types) - 1
		}
		c.checkAType(t, argsCount)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkAType(v AType, argsCount int) *TypeCheckResult {
	switch v := v.(type) {
	case *ConType:
		c.checkConType(v, argsCount)
	case *VarType:
		c.checkVarType(v, argsCount)
	case *TupleType:
		c.checkTupleType(v, argsCount)
	case *ParenType:
		c.checkParenType(v, argsCount)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkConType(v *ConType, argsCount int) *TypeCheckResult {
	c.checkDataTypeExistence(v, v.id, argsCount)
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkVarType(v *VarType, argsCount int) *TypeCheckResult {
	if argsCount > 0 {
		c.errFatal(v, fmt.Sprintf("Type '%s' can't have any parameters", v.id))
	}

	parameterTypes := c.context[CurParameterTypes].(*ParameterTypes).types

	_, lastArrowType := c.context[LastArrowType]
	if lastArrowType {
		_, defined := parameterTypes[v.id]
		if !defined {
			c.errFatal(v, fmt.Sprintf("Type '%s' not defined", v.id))
		}
	} else {
		parameterTypes[v.id] = voidMember
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkTupleType(v *TupleType, argsCount int) *TypeCheckResult {
	if argsCount > 0 {
		c.errFatal(v, "Tuple type can't have any parameters")
	}

	for _, t := range v.types {
		c.checkType(t)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkParenType(v *ParenType, argsCount int) *TypeCheckResult {
	if argsCount > 0 {
		c.errFatal(v, "Parenthesis type can't have any parameters")
	}
	c.checkType(v.t)
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

func (c *TypeChecker) checkDataTypeExistence(v ASTNode, constr string, argsCount int) *TypeCheckResult {
	t, ok := c.definedTypes[constr]

	if !ok {
		c.errFatal(v, fmt.Sprintf("Type '%s' not defined", constr))
	}

	// We check the parameters count of the type
	actualArity := len(t.vars)
	if argsCount != actualArity {
		c.errFatal(v, fmt.Sprintf("Type '%s' needs %d parameters", constr, actualArity))
	}

	return &TypeCheckResult{}
}

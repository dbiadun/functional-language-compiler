package main

import (
	"fmt"
	"log"
	"strings"
)

/////////////////////////////////// TYPE_CHECKER /////////////////////////////////////

type TypeChecker struct {
	types        map[string]*FunType    // types of declared functions and variables
	constrs      map[string]*ConstrInfo // all defined constructors
	definedTypes map[string]*DataType   // all defined types
	context      map[int]ContextData    // additional context
	uniuqeNum    int                    // for unique numbers
}

type ConstrInfo struct {
	t       *DataType
	constrT *ConstrType
}

func (c *TypeChecker) init() {
	c.types = make(map[string]*FunType)
	c.constrs = make(map[string]*ConstrInfo)
	c.definedTypes = make(map[string]*DataType)
	c.context = make(map[int]ContextData)

	c.addBuiltins()
}

func (*TypeChecker) errFatal(v ASTNode, s string) {
	line, col := v.getPos()
	log.Fatalf("TypeChecker error at line %d, column %d: %s\n", line, col, s)
}

func (c *TypeChecker) getNum() int {
	c.uniuqeNum++
	return c.uniuqeNum
}

////////////////////////////////////// CONTEXT ///////////////////////////////////////

const (
	CurDataType          = iota // *DataType
	CurDataTypeVars             // *DataTypeVars
	CurParameterTypes           // *ParameterTypes
	CurChangeBackup             // *ChangeBackup
	CurValidRhsType             // *FunType
	CurTypeSubstitutions        // *TypeSubstitutions
	LastArrowType               // void
	TypesCheck                  // void
	GlobalNamesCheck            // void
	LastCheck                   // void
)

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

type WantedType struct {
	t *FunType
}

func (*WantedType) contextData() {}

type TypeSubstitutions struct {
	subst map[string]*FunType
}

func (*TypeSubstitutions) contextData() {}

/////////////////////////////////// BUILTINS ///////////////////////////////////////

func (c *TypeChecker) addBuiltins() {
	intDataType := &DataType{constr: "Int", vars: []string{}}
	c.definedTypes["Int"] = intDataType

	charDataType := &DataType{constr: "Char", vars: []string{}}
	c.definedTypes["Char"] = charDataType

	stringDataType := &DataType{constr: "String", vars: []string{}}
	c.definedTypes["String"] = stringDataType
}

func getIntType() *FunType {
	return &FunType{types: []BType{&TypeApp{types: []AType{&ConType{id: "Int"}}}}}
}

func getCharType() *FunType {
	return &FunType{types: []BType{&TypeApp{types: []AType{&ConType{id: "Char"}}}}}
}

func getStringType() *FunType {
	return &FunType{types: []BType{&TypeApp{types: []AType{&ConType{id: "String"}}}}}
}

//////////////////////////////// TYPES CHANGING //////////////////////////////////////

// Things used to modify the environment and later revert the modifications.
// Useful, for example, when we get a new declared variable and we need to change
// the type assigned to it for the time needed to process a subtree of the AST.

type ChangeBackup struct {
	changedTypes map[string]*FunType
	addedTypes   map[string]void
}

func (*ChangeBackup) contextData() {}

func (c *TypeChecker) setTypes(varNames []string, types []*FunType) *ChangeBackup {
	backup := new(ChangeBackup)
	backup.changedTypes = make(map[string]*FunType)
	backup.addedTypes = make(map[string]void)

	for i, name := range varNames {
		newType := types[i]
		oldType, ok := c.types[name]
		if ok {
			backup.changedTypes[name] = oldType
		} else {
			backup.addedTypes[name] = voidMember
		}

		c.types[name] = newType
	}

	return backup
}

func (c *TypeChecker) revertTypes(backup *ChangeBackup) {
	if backup == nil {
		return
	}

	for varName, varType := range backup.changedTypes {
		c.types[varName] = varType
	}

	for varName := range backup.addedTypes {
		delete(c.types, varName)
	}
}

/////////////////////////////////// TYPE OPERATIONS //////////////////////////////////////

func (c *TypeChecker) detachTypes(parentNode ASTNode, funType *FunType, typesCount int) ([]*FunType, *FunType) {
	// Detaches the first `typesCount` types from the function type

	// We need to make sure that the type isn't wrapped in parentheses
	c.reduceParens(funType)
	if typesCount >= len(funType.types) {
		c.errFatal(parentNode, "Too many arguments to match the function type")
	}

	var detachedTypes []*FunType

	for i := 0; i < typesCount; i++ {
		nextBType := funType.types[i]
		nextFunType := &FunType{types: []BType{nextBType}}
		detachedTypes = append(detachedTypes, nextFunType)
	}

	remainingTypes := funType.types[typesCount:]
	modifiedType := &FunType{types: remainingTypes}

	return detachedTypes, modifiedType
}

func (c *TypeChecker) mergeTypes(parentNode ASTNode, funType *FunType, typesCount int) *FunType {
	// Merges the last `typesCount` types in the function type

	newFunType := &FunType{}
	oldTypesCount := len(funType.types)
	if typesCount > oldTypesCount {
		c.errFatal(parentNode, "Too few types to merge")
	}

	funType1 := c.typeDeepCopy(funType)
	funType2 := c.typeDeepCopy(funType)

	bTypes := funType1.types[:oldTypesCount-typesCount]

	resultFunType := &FunType{types: funType2.types[oldTypesCount-typesCount:]}

	resultBType := &TypeApp{types: []AType{&ParenType{t: resultFunType}}}

	bTypes = append(bTypes, resultBType)

	newFunType.types = bTypes

	return newFunType
}

func (c *TypeChecker) reduceParens(v Type) {
	// Reduces the unnecessary parentheses

	var replacementBTypes []BType

	funType := v.(*FunType)
	for i, bType := range funType.types {
		typeApp := bType.(*TypeApp)
		for j, aType := range typeApp.types {
			// First, we reduce recursively types within tuples and parens
			_, ok := aType.(*TupleType)
			if ok {
				tupleType := aType.(*TupleType)
				for _, t := range tupleType.types {
					c.reduceParens(t)
				}
			}

			_, ok = aType.(*ParenType)
			if ok {
				parenType := aType.(*ParenType)
				c.reduceParens(parenType.t)
				innerFunType := parenType.t.(*FunType)
				if len(innerFunType.types) == 1 {
					innerBType := innerFunType.types[0]
					innerTypeApp := innerBType.(*TypeApp)
					if len(innerTypeApp.types) == 1 {
						// If type within parenthesis is just AType, we can put it directly inside typeApp.types
						typeApp.types[j] = innerTypeApp.types[0]
					} else if len(typeApp.types) == 1 {
						// If type within parenthesis is BType, and our ParenType is the only AType in the application,
						// we can put the type from parenthesis inside funType
						funType.types[i] = innerBType
					}
				} else if i == len(funType.types)-1 && len(typeApp.types) == 1 {
					// We can reduce parens in the last part of arrow divided type
					replacementBTypes = innerFunType.types
				}
			}
		}
	}

	// Replacement of the type after the last arrow with types inside parenthesis
	if replacementBTypes != nil {
		funType.types = funType.types[:len(funType.types)-1]
		funType.types = append(funType.types, replacementBTypes...)
	}
}

func (c *TypeChecker) isVarType(t BType) bool {
	// Checks if the type is a single variable type

	typeApp := t.(*TypeApp)

	if len(typeApp.types) != 1 {
		return false
	}

	aType := typeApp.types[0]
	_, isVar := aType.(*VarType)

	return isVar
}

func (c *TypeChecker) createUniqueParams(funType *FunType, paramTypeChar byte) {
	// Changes parameter types names so that they are unique

	substitutions := make(map[string]string)
	c.changeParams(funType, paramTypeChar, substitutions)
}

func (c *TypeChecker) changeParams(funType *FunType, paramTypeChar byte, substitutions map[string]string) {
	for _, bType := range funType.types {
		typeApp := bType.(*TypeApp)
		for _, aType := range typeApp.types {
			switch aType := aType.(type) {
			case *ConType:
				continue
			case *VarType:
				subst, ok := substitutions[aType.id]
				if ok {
					aType.id = subst
				} else {
					if strings.ContainsAny(aType.id, "%$") {
						continue
					}
					varNum := c.getNum()
					newId := fmt.Sprintf("%s%c%d", aType.id, paramTypeChar, varNum)
					oldId := aType.id
					aType.id = newId

					substitutions[oldId] = newId
				}
			case *TupleType:
				for _, t := range aType.types {
					c.changeParams(t.(*FunType), paramTypeChar, substitutions)
				}
			case *ParenType:
				c.changeParams(aType.t.(*FunType), paramTypeChar, substitutions)
			}
		}
	}
}

func (c *TypeChecker) typeDeepCopy(oldType Type) *FunType {
	// Creates a deep copy of a type

	funType := &FunType{}
	var bTypes []BType
	for _, oldBType := range oldType.(*FunType).types {
		oldTypeApp := oldBType.(*TypeApp)
		var aTypes []AType
		for _, oldAType := range oldTypeApp.types {
			switch oldAType := oldAType.(type) {
			case *ConType:
				aTypes = append(aTypes, &ConType{id: strings.Clone(oldAType.id)})
			case *VarType:
				aTypes = append(aTypes, &VarType{id: strings.Clone(oldAType.id)})
			case *TupleType:
				var types []Type
				for _, oldType := range oldAType.types {
					types = append(types, c.typeDeepCopy(oldType))
				}
				aTypes = append(aTypes, &TupleType{types: types})
			case *ParenType:
				aTypes = append(aTypes, &ParenType{t: c.typeDeepCopy(oldAType.t)})
			}
		}
		typeApp := &TypeApp{types: aTypes}
		bTypes = append(bTypes, typeApp)
	}
	funType.types = bTypes

	return funType
}

func (c *TypeChecker) applySubstitutions(t Type) {
	// Replaces parameter types with its calculated substitutions

	s, ok := c.context[CurTypeSubstitutions]
	if !ok {
		c.errFatal(t, "No substitutions")
	}
	substs := s.(*TypeSubstitutions)

	funType := t.(*FunType)
	for _, bType := range funType.types {
		typeApp := bType.(*TypeApp)

		for j, aType := range typeApp.types {
			switch aType := aType.(type) {
			case *ConType:
				continue
			case *VarType:
				t, ok := substs.subst[aType.id]
				if ok {
					c.applySubstitutions(t)
					typeApp.types[j] = &ParenType{t: t}
				}
			case *TupleType:
				for _, t := range aType.types {
					c.applySubstitutions(t)
				}
			case *ParenType:
				c.applySubstitutions(aType.t)
			}
		}
	}
}

func (c *TypeChecker) printType(t Type) {
	funType := t.(*FunType)
	for i, bType := range funType.types {
		typeApp := bType.(*TypeApp)
		if i > 0 {
			fmt.Print(" -> ")
		}
		for j, aType := range typeApp.types {
			if j > 0 {
				fmt.Print(" ")
			}
			switch aType := aType.(type) {
			case *ConType:
				fmt.Print(aType.id)
			case *VarType:
				fmt.Print(aType.id)
			case *TupleType:
				fmt.Print("(")
				for k, subtype := range aType.types {
					if k > 0 {
						fmt.Print(", ")
					}
					c.printType(subtype)
				}
				fmt.Print(")")
			case *ParenType:
				fmt.Print("(")
				c.printType(aType.t)
				fmt.Print(")")
			}
		}
	}
}

///////////////////////////////////// TYPES MATCHING ///////////////////////////////////////

// When matching we allow substitution of parameter types.
// Parameter types originating from the currently declared function type have '%' sign in
// their names and those originating from other sources (for example function calls of case
// expressions) have '$' in their names. The first ones can't be substituted by any type,
// because they are defined so that they can fit any type provided by the user when calling
// the function. The second ones may be substituted so that the whole types match each other.

func (c *TypeChecker) checkMatch(v ASTNode, validType Type, actualType Type, sameLevel bool) {
	// Checks if two types are matching, and applies needed substitutions so that they are the same

	// We need to reduce parentheses in both types, so that they have the same form
	c.reduceParens(validType)
	c.reduceParens(actualType)
	c.checkTypeMatch(v, validType, actualType, sameLevel)
}

func (c *TypeChecker) checkTypeMatch(v ASTNode, validType Type, actualType Type, sameLevel bool) {
	switch validType := validType.(type) {
	case *FunType:
		actualType, ok := actualType.(*FunType)
		if ok {
			c.checkFunTypeMatch(v, validType, actualType, sameLevel)
			return
		}
	}
	c.errFatal(v, "Types not matching")
}

func (c *TypeChecker) checkFunTypeMatch(v ASTNode, validType *FunType, actualType *FunType, sameLevel bool) {
	validTypeLen := len(validType.types)
	actualTypeLen := len(actualType.types)
	if validTypeLen != actualTypeLen {
		// If function types has different arity we merge the last subtypes (put them in parentheses)
		// in the type with the bigger arity in hopes that we will match those subtypes with some
		// parameter type
		if validTypeLen > actualTypeLen && c.isVarType(actualType.types[actualTypeLen-1]) {
			validType = c.mergeTypes(v, validType, validTypeLen-actualTypeLen+1)
		} else if actualTypeLen > validTypeLen && c.isVarType(validType.types[validTypeLen-1]) {
			actualType = c.mergeTypes(v, actualType, actualTypeLen-validTypeLen+1)
		} else {
			c.errFatal(v, "Types arity not matching")
		}
	}

	for i := range validType.types {
		c.checkBTypeMatch(v, validType.types[i], actualType.types[i], sameLevel)
	}
}

func (c *TypeChecker) checkBTypeMatch(v ASTNode, validType BType, actualType BType, sameLevel bool) {
	switch validType := validType.(type) {
	case *TypeApp:
		actualType, ok := actualType.(*TypeApp)
		if ok {
			c.checkTypeAppMatch(v, validType, actualType, sameLevel)
			return
		}
	}
	c.errFatal(v, "Types between arrows not matching")
}

func (c *TypeChecker) checkTypeAppMatch(v ASTNode, validType *TypeApp, actualType *TypeApp, sameLevel bool) {
	if len(validType.types) != len(actualType.types) {
		// When comparing function's lhs and rhs type we can substitute only rhs type parameters, but when comparing
		// function call type with its arguments, case branches etc. we can substitute parameters on both sides
		if len(actualType.types) == 1 || (sameLevel && len(validType.types) == 1) {
			var v1, v2 *TypeApp
			if len(actualType.types) == 1 {
				v1 = validType
				v2 = actualType
			} else {
				v1 = actualType
				v2 = validType
			}
			validFunType := &FunType{types: []BType{v1}}
			actualType, isVarType := v2.types[0].(*VarType)

			if !isVarType {
				c.errFatal(v, "Types not matching")
			}

			if !c.mayBeSubstituted(actualType) {
				c.errFatal(v, "Types not matching (trying to substitute type that can't be substituted)")
			}

			substs := c.context[CurTypeSubstitutions].(*TypeSubstitutions)
			t, ok := substs.subst[actualType.id]
			if ok {
				// If types were swapped, then they had to on the same level, so there is no difference
				// in the order of passing them to this function
				c.checkMatch(v, validFunType, t, sameLevel)
			} else {
				substs.subst[actualType.id] = validFunType
			}
		} else {
			c.errFatal(v, "Types not matching (different number of args)")
		}
	} else {
		// Substitution can't be allowed on a type name when it has parameters
		// It can be done only if we have a type without parameters or on one of the parameters
		var substitutionAllowed bool
		if len(validType.types) == 1 {
			substitutionAllowed = true
		}
		for i := range validType.types {
			if i > 0 {
				substitutionAllowed = true
			}
			c.checkATypeMatch(v, validType.types[i], actualType.types[i], substitutionAllowed, sameLevel)
		}
	}

}

func (c *TypeChecker) checkATypeMatch(v ASTNode, validType AType, actualType AType, substitutionAllowed bool, sameLevel bool) {
	validTypeBackup := validType
	validType, ok := validType.(*VarType)

	mayBeSubstituted := true
	if ok {
		// Only parameter types originating from the rhs may be substituted
		mayBeSubstituted = c.mayBeSubstituted(validType.(*VarType))
	}

	if sameLevel && ok && mayBeSubstituted {
		c.checkVarTypeMatch(v, actualType, validType.(*VarType), substitutionAllowed, sameLevel)
	} else {
		validType = validTypeBackup
		switch actualType := actualType.(type) {
		case *ConType:
			c.checkConTypeMatch(v, validType, actualType)
		case *VarType:
			c.checkVarTypeMatch(v, validType, actualType, substitutionAllowed, sameLevel)
		case *TupleType:
			c.checkTupleTypeMatch(v, validType, actualType, sameLevel)
		case *ParenType:
			c.checkParenTypeMatch(v, validType, actualType, sameLevel)
		}
	}
}

func (c *TypeChecker) checkConTypeMatch(v ASTNode, validType AType, actualType *ConType) {
	validType, ok := validType.(*ConType)
	if ok {
		if validType.(*ConType).id != actualType.id {
			c.errFatal(v, "Types not matching (different data types)")
		}
	} else {
		c.errFatal(v, "Types not matching (type constructor with something else)")
	}
}

func (c *TypeChecker) checkVarTypeMatch(v ASTNode, validType AType, actualType *VarType, substitutionAllowed bool, sameLevel bool) {
	if !substitutionAllowed {
		c.errFatal(v, "Types not matching (type constructor matching with parameter type)")
	}

	if c.mayBeSubstituted(actualType) {
		substs := c.context[CurTypeSubstitutions].(*TypeSubstitutions)
		validFunType := &FunType{types: []BType{&TypeApp{types: []AType{validType}}}}
		c.reduceParens(validFunType)

		t, ok := substs.subst[actualType.id]
		if ok {
			c.checkMatch(v, validFunType, t, sameLevel)
		} else {
			substs.subst[actualType.id] = validFunType
		}
	} else {
		validType, ok := validType.(*VarType)
		if ok {
			if validType.id != actualType.id {
				c.errFatal(v, "Types not matching (different type parameters)")
			}
		} else {
			c.errFatal(v, "Types not matching (parameter type that can't be substituted and something else)")
		}
	}
}

func (c *TypeChecker) checkTupleTypeMatch(v ASTNode, validType AType, actualType *TupleType, sameLevel bool) {
	validType, ok := validType.(*TupleType)
	if ok {
		validType := validType.(*TupleType)
		if len(validType.types) == len(actualType.types) {
			for i := range validType.types {
				c.checkTypeMatch(v, validType.types[i], actualType.types[i], sameLevel)
			}
		} else {
			c.errFatal(v, "Types not matching (tuples of different lengths)")
		}
	} else {
		c.errFatal(v, "Types not matching (tuple type with something else)")
	}
}

func (c *TypeChecker) checkParenTypeMatch(v ASTNode, validType AType, actualType *ParenType, sameLevel bool) {
	validType, ok := validType.(*ParenType)
	if ok {
		validType := validType.(*ParenType)
		c.checkTypeMatch(v, validType.t, actualType.t, sameLevel)
	} else {
		c.errFatal(v, "Types not matching (paren type with something else)")
	}
}

func (c *TypeChecker) mayBeSubstituted(v *VarType) bool {
	// Checks if this parameter type may be substituted (even if that's true we don't know
	// if it doesn't already have a substitute).
	return strings.Contains(v.id, "$")
}

//////////////////////////////// THE ACTUAL CHECK /////////////////////////////////////

type TypeCheckResult struct {
	t *FunType
}

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
		c.errFatal(v, "Parentheses type can't have any parameters")
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
	_, lastCheck := c.context[LastCheck]

	if lastCheck {
		c.context[CurTypeSubstitutions] = &TypeSubstitutions{subst: make(map[string]*FunType)}

		// Should set CurChangeBackup and CurValidRhsType
		c.checkFunLhs(v.lhs)

		// We save the types changed while checking lhs, to revert them later
		typesBackup := c.context[CurChangeBackup].(*ChangeBackup)
		validRhsType := c.context[CurValidRhsType].(*WantedType).t
		rhsType := c.checkRhs(v.rhs).t
		c.revertTypes(typesBackup)

		c.checkMatch(v, validRhsType, rhsType, false)
	}

	return &TypeCheckResult{}
}

func (c *TypeChecker) checkVarDecl(v *VarDecl) *TypeCheckResult {
	_, lastCheck := c.context[LastCheck]

	if lastCheck {
		c.context[CurTypeSubstitutions] = &TypeSubstitutions{subst: make(map[string]*FunType)}

		var lhsType *FunType
		pat, ok := v.pat.(*PatArg)
		if ok {
			v, ok := pat.arg.(*APatVar)
			if ok {
				t, ok := c.types[v.id]
				if ok {
					lhsType = t
				} else {
					c.errFatal(v, fmt.Sprintf("Function '%s' has no declared type", v.id))
				}
			} else {
				c.errFatal(v, "Constructor name instead of a function name")
			}
		} else {
			c.errFatal(v, "Constructor pattern instead of a function declaration")
		}

		lhsTypeCopy := c.typeDeepCopy(lhsType)
		c.createUniqueParams(lhsTypeCopy, '%')
		rhsType := c.checkRhs(v.rhs).t
		c.checkMatch(v, lhsTypeCopy, rhsType, false)
	}

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
		c.errFatal(v, "Parentheses type can't have any parameters")
	}
	c.checkType(v.t)
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkFunLhs(v FunLhs) *TypeCheckResult {
	switch v := v.(type) {
	case *DeclLhs:
		c.checkDeclLhs(v)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkDeclLhs(v *DeclLhs) *TypeCheckResult {
	t, ok := c.types[v.fun]
	if !ok {
		c.errFatal(v, fmt.Sprintf("Function '%s' has no declared type", v.fun))
	}
	tCopy := c.typeDeepCopy(t)
	c.createUniqueParams(tCopy, '%')

	var patternVars []string
	for _, arg := range v.args {
		switch arg := arg.(type) {
		case *APatVar:
			patternVars = append(patternVars, arg.id)
		default:
			c.errFatal(v, fmt.Sprintf("Complex pattern in '%s' function declaration", v.fun))
		}
	}

	paramsCount := len(v.args)
	patternTypes, modifiedType := c.detachTypes(v, tCopy, paramsCount)

	typesBackup := c.setTypes(patternVars, patternTypes)
	c.context[CurChangeBackup] = typesBackup
	c.context[CurValidRhsType] = &WantedType{modifiedType}

	return &TypeCheckResult{}
}

func (c *TypeChecker) checkRhs(v Rhs) *TypeCheckResult {
	switch v := v.(type) {
	case *DeclExp:
		return c.checkDeclExp(v)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkDeclExp(v *DeclExp) *TypeCheckResult {
	return c.checkExp(v.e)
}

func (c *TypeChecker) checkExp(v Exp) *TypeCheckResult {
	switch v := v.(type) {
	case *EFun:
		return c.checkEFun(v)
	case *ECase:
		return c.checkECase(v)
	case *EMul:
		return c.checkEMul(v)
	case *EDiv:
		return c.checkEDiv(v)
	case *EAdd:
		return c.checkEAdd(v)
	case *ESub:
		return c.checkESub(v)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkEFun(v *EFun) *TypeCheckResult {
	return c.checkFExp(v.fExp)
}

func (c *TypeChecker) checkECase(v *ECase) *TypeCheckResult {
	expType := c.checkExp(v.e).t

	if len(v.alts) == 0 {
		c.errFatal(v, "Empty case")
	}

	var altType *FunType
	for i, alt := range v.alts {
		t := c.checkAlternative(alt, expType).t

		if i == 0 {
			altType = t
		} else {
			// We need to check if all the alternatives return results of the same type
			c.checkMatch(v, altType, t, true)
		}
	}

	return &TypeCheckResult{altType}
}

func (c *TypeChecker) checkEMul(v *EMul) *TypeCheckResult {
	return c.checkBinOp(v, v.e1, v.e2)
}

func (c *TypeChecker) checkEDiv(v *EDiv) *TypeCheckResult {
	return c.checkBinOp(v, v.e1, v.e2)
}

func (c *TypeChecker) checkEAdd(v *EAdd) *TypeCheckResult {
	return c.checkBinOp(v, v.e1, v.e2)
}

func (c *TypeChecker) checkESub(v *ESub) *TypeCheckResult {
	return c.checkBinOp(v, v.e1, v.e2)
}

func (c *TypeChecker) checkBinOp(v ASTNode, e1 Exp, e2 Exp) *TypeCheckResult {
	intType := getIntType()

	t1 := c.checkExp(e1)
	t2 := c.checkExp(e2)

	c.checkMatch(v, intType, t1.t, false)
	c.checkMatch(v, intType, t2.t, false)
	return &TypeCheckResult{intType}
}

func (c *TypeChecker) checkFExp(v FExp) *TypeCheckResult {
	switch v := v.(type) {
	case *FApp:
		return c.checkFApp(v)
	case *FArg:
		return c.checkFArg(v)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkFApp(v *FApp) *TypeCheckResult {
	funType := c.checkFExp(v.fun).t
	// We need to apply all substitutions to be sure, that the type is not folded before trying to apply arguments
	c.applySubstitutions(funType)
	c.reduceParens(funType)

	argType := c.checkAExp(v.arg).t

	firstTypeSlice, resultType := c.detachTypes(v, funType, 1)

	c.checkMatch(v, firstTypeSlice[0], argType, true)
	return &TypeCheckResult{resultType}
}

func (c *TypeChecker) checkFArg(v *FArg) *TypeCheckResult {
	return c.checkAExp(v.arg)
}

func (c *TypeChecker) checkAExp(v AExp) *TypeCheckResult {
	switch v := v.(type) {
	case *Var:
		return c.checkVar(v)
	case *Constr:
		return c.checkConstr(v)
	case *Lit:
		return c.checkLit(v)
	case *ParenExp:
		return c.checkParenExp(v)
	case *Tuple:
		return c.checkTuple(v)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkVar(v *Var) *TypeCheckResult {
	t, ok := c.types[v.id]
	if ok {
		c.reduceParens(t)
		if len(t.types) > 1 {
			// We need to create unique type params in the type of the called function
			tCopy := c.typeDeepCopy(t)
			c.createUniqueParams(tCopy, '$')
			return &TypeCheckResult{tCopy}
		}
		return &TypeCheckResult{t}
	}

	c.errFatal(v, fmt.Sprintf("Variable '%s' not defined", v.id))
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkConstr(v *Constr) *TypeCheckResult {
	// Checks constructor existence and type (in the form of a function type)
	// and changes parameter types for unique ones
	constrInfo, ok := c.constrs[v.id]
	if !ok {
		c.errFatal(v, fmt.Sprintf("Constructor '%s' not defined", v.id))
	}

	constrType := constrInfo.constrT
	constrResultType := constrInfo.t

	var arrowTypes []BType
	for _, t := range constrType.args {
		nextType := &TypeApp{types: []AType{t}}
		arrowTypes = append(arrowTypes, nextType)
	}

	var resultTypeArgs []AType
	resultTypeArgs = append(resultTypeArgs, &ConType{id: constrResultType.constr})
	for _, argVar := range constrResultType.vars {
		arg := &VarType{id: argVar}
		resultTypeArgs = append(resultTypeArgs, arg)
	}

	resultType := &TypeApp{types: resultTypeArgs}
	arrowTypes = append(arrowTypes, resultType)

	exprType := &FunType{types: arrowTypes}

	exprTypeCopy := c.typeDeepCopy(exprType)
	c.createUniqueParams(exprTypeCopy, '$')
	return &TypeCheckResult{exprTypeCopy}
}

func (c *TypeChecker) checkLit(v *Lit) *TypeCheckResult {
	return c.checkLiteral(v.lit)
}

func (c *TypeChecker) checkParenExp(v *ParenExp) *TypeCheckResult {
	return c.checkExp(v.e)
}

func (c *TypeChecker) checkTuple(v *Tuple) *TypeCheckResult {
	var types []Type
	for _, exp := range v.exps {
		nextType := c.checkExp(exp).t
		types = append(types, nextType)
	}
	tupleType := &TupleType{types: types}
	typeApp := &TypeApp{types: []AType{tupleType}}
	funType := &FunType{types: []BType{typeApp}}

	return &TypeCheckResult{funType}
}

func (c *TypeChecker) checkAlternative(v *Alternative, caseType *FunType) *TypeCheckResult {
	c.context[CurChangeBackup] = &ChangeBackup{make(map[string]*FunType), make(map[string]void)}
	// Should add vars to CurChangeBackup
	c.checkPat(v.pat, caseType)

	typesBackup := c.context[CurChangeBackup].(*ChangeBackup)
	expType := c.checkExp(v.exp).t
	c.revertTypes(typesBackup)

	return &TypeCheckResult{expType}
}

func (c *TypeChecker) checkPat(v Pat, caseType *FunType) *TypeCheckResult {
	switch v := v.(type) {
	case *PatArg:
		c.checkPatArg(v, caseType)
	case *PatCon:
		c.checkPatCon(v, caseType)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkPatArg(v *PatArg, caseType *FunType) *TypeCheckResult {
	c.checkAPat(v.arg, caseType)
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkPatCon(v *PatCon, caseType *FunType) *TypeCheckResult {
	vars := make(map[string]void)

	constr := &Constr{id: v.conId}
	constr.setPos(v.getPos())
	funType := c.checkConstr(constr).t

	argTypes, resultType := c.detachTypes(v, funType, len(funType.types)-1)

	if len(v.args) < 1 {
		c.errFatal(v, "No constructor arguments in the pattern")
	}

	if len(argTypes) != len(v.args) {
		c.errFatal(v, "Invalid number of constructor args")
	}

	for i, arg := range v.args {
		c.checkAPat(arg, argTypes[i])

		// We need to check if vars in the pattern arent repeated
		vArg := arg.(*APatVar)
		_, ok := vars[vArg.id]
		if ok {
			c.errFatal(vArg, fmt.Sprintf("Repeated '%s' pattern variable", vArg.id))
		} else {
			vars[vArg.id] = voidMember
		}
	}

	c.checkMatch(v, caseType, resultType, false)

	return &TypeCheckResult{}
}

func (c *TypeChecker) checkAPat(v APat, argType *FunType) *TypeCheckResult {
	switch v := v.(type) {
	case *APatVar:
		c.checkAPatVar(v, argType)
	case *APatCon:
		c.checkAPatCon(v, argType)
	case *APatLit:
		c.checkAPatLit(v, argType)
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkAPatVar(v *APatVar, argType *FunType) *TypeCheckResult {
	changeBackup := c.context[CurChangeBackup].(*ChangeBackup)
	oldType, ok := c.types[v.id]
	if ok {
		changeBackup.changedTypes[v.id] = oldType
	} else {
		changeBackup.addedTypes[v.id] = voidMember
	}
	c.types[v.id] = argType
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkAPatCon(v *APatCon, argType *FunType) *TypeCheckResult {
	constr := &Constr{id: v.id}
	constr.setPos(v.getPos())

	funType := c.checkConstr(constr).t
	if len(funType.types) > 1 {
		c.errFatal(v, "Constructor used as a pattern arg can't require any parameters")
	}

	c.checkMatch(v, argType, funType, false)

	return &TypeCheckResult{}
}

func (c *TypeChecker) checkAPatLit(v *APatLit, argType *FunType) *TypeCheckResult {
	litType := c.checkLiteral(v.lit).t
	c.checkMatch(v, argType, litType, false)
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkLiteral(v Literal) *TypeCheckResult {
	switch x := v.(type) {
	case *Int:
		_ = x
		return c.checkInt()
	case *Char:
		return c.checkChar()
	case *String:
		return c.checkString()
	}
	return &TypeCheckResult{}
}

func (c *TypeChecker) checkInt() *TypeCheckResult {
	return &TypeCheckResult{getIntType()}
}

func (c *TypeChecker) checkChar() *TypeCheckResult {
	return &TypeCheckResult{getCharType()}
}

func (c *TypeChecker) checkString() *TypeCheckResult {
	return &TypeCheckResult{getStringType()}
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

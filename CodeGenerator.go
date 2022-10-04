package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

////////////////////////////////////// GENERATOR ///////////////////////////////////////////

type CodeGenerator struct {
	code            []*Instr
	codeDump        [][]*Instr
	functionDefs    map[string][]*Instr
	functionArities map[string]int
	instrSchemes    map[int]*InstrScheme
	constrTags      map[string]int
	env             *VarEnv
	iteration       int
}

func newCodeGenerator() *CodeGenerator {
	codeGenerator := new(CodeGenerator)
	codeGenerator.init()
	return codeGenerator
}

func (g *CodeGenerator) init() {
	g.functionDefs = make(map[string][]*Instr)
	g.functionArities = make(map[string]int)
	g.initInstrSchemes()
	g.constrTags = make(map[string]int)
	g.env = createVarEnv()

	g.addCompiledConstructors()
}

func (g *CodeGenerator) addCompiledConstructors() {
	g.constrTags["False"] = 0
	g.constrTags["True"] = 1
}

func (g *CodeGenerator) genCodePart(decls TopDecls) {
	g.iteration = 0
	g.genTopDecls(decls)
}

func (g *CodeGenerator) emit(outputFile string) {
	output := g.genOutput()

	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Println(err)
	} else {
		file.WriteString(output)
	}
	file.Close()
}

func (g *CodeGenerator) saveFunction(fun string) {
	g.functionDefs[fun] = g.code
}

func (g *CodeGenerator) reset() {
	g.code = []*Instr{}
	g.env = createVarEnv()
}

func (g *CodeGenerator) saveCode() {
	g.codeDump = append(g.codeDump, g.code)
	g.code = []*Instr{}
}

func (g *CodeGenerator) restoreCode() {
	l := len(g.codeDump)
	g.code = g.codeDump[l-1]
	g.codeDump = g.codeDump[:l-1]
}

func (g *CodeGenerator) pushInstr(instrType int, args ...string) {
	scheme, ok := g.instrSchemes[instrType]
	if !ok {
		errFatalAnywhere(fmt.Sprintf("Unknown instruction type (%d)", instrType))
	}

	instr := scheme.create(args...)
	g.code = append(g.code, instr)
}

//////////////////////////////////////// ERRORS ////////////////////////////////////////////

func errFatal(v ASTNode, s string) {
	line, col := v.getPos()
	log.Fatalf("Code generator error at line %d, column %d: %s\n", line, col, s)
}

func errFatalAnywhere(s string) {
	log.Fatalf("Code generator error: %s\n", s)
}

////////////////////////////////////// ENVIRONMENT /////////////////////////////////////////

type VarEnv struct {
	vars   map[string]int
	offset int
}

func createVarEnv() *VarEnv {
	e := new(VarEnv)
	e.vars = make(map[string]int)

	return e
}

func (e *VarEnv) addVar(v string) {
	e.offset++
	e.vars[v] = e.offset
}

func (e *VarEnv) getVar(v string) int {
	n, ok := e.vars[v]
	if !ok {
		errFatalAnywhere("Accessed variable not defined")
	}
	return e.offset - n
}

func (e *VarEnv) varDefined(v string) bool {
	_, defined := e.vars[v]
	return defined
}

func (e *VarEnv) changeStackSize(n int) {
	e.offset += n
}

/////////////////////////////////// ENVIRONMENT CHANGE /////////////////////////////////////

type EnvBackup struct {
	changed map[string]int
	added   []string
}

func (b *EnvBackup) size() int {
	return len(b.changed) + len(b.added)
}

func (e *VarEnv) applyChange(vars []string) *EnvBackup {
	backup := new(EnvBackup)
	backup.changed = make(map[string]int)

	for i := len(vars) - 1; i >= 0; i-- {
		v := vars[i]
		if e.varDefined(v) {
			tag := e.getVar(v)
			backup.changed[v] = tag
		} else {
			backup.added = append(backup.added, v)
		}

		e.addVar(v)
	}

	return backup
}

func (e *VarEnv) revertChange(backup *EnvBackup) {
	for v, tag := range backup.changed {
		e.changeStackSize(-1)
		e.vars[v] = tag
	}

	for _, v := range backup.added {
		e.changeStackSize(-1)
		delete(e.vars, v)
	}
}

////////////////////////////////////// INSTRUCTIONS ////////////////////////////////////////

type Instr struct {
	name     string
	argNames []string
	args     []string
}

func (i *Instr) show() string {
	if len(i.argNames) != len(i.args) {
		errFatalAnywhere("Invalid number of args passed to instruction")
	}

	var res strings.Builder

	res.WriteString("&")
	res.WriteString(i.name)
	res.WriteString("{")

	for j := range i.argNames {
		if j > 0 {
			res.WriteString(",")
		}

		res.WriteString(i.argNames[j])
		res.WriteString(":")
		res.WriteString(i.args[j])
	}

	res.WriteString("}")

	return res.String()
}

func showCode(code []*Instr) string {
	var res strings.Builder

	res.WriteString("[]GInstr{\n")

	for _, instr := range code {
		res.WriteString(fmt.Sprintf("%s,\n", instr.show()))
	}

	res.WriteString("}")
	return res.String()
}

type InstrScheme struct {
	name     string
	argNames []string
}

func (s *InstrScheme) create(args ...string) *Instr {
	if len(s.argNames) != len(args) {
		errFatalAnywhere("Invalid number of args passed to instruction creation")
	}

	return &Instr{s.name, s.argNames, args}
}

func (g *CodeGenerator) initInstrSchemes() {
	g.instrSchemes = make(map[int]*InstrScheme)
	g.instrSchemes[cPushInt] = &InstrScheme{"IPushInt", []string{"n"}}
	g.instrSchemes[cPushString] = &InstrScheme{"IPushString", []string{"s"}}
	g.instrSchemes[cPushGlobal] = &InstrScheme{"IPushGlobal", []string{"f"}}
	g.instrSchemes[cPush] = &InstrScheme{"IPush", []string{"n"}}
	g.instrSchemes[cMkApp] = &InstrScheme{"IMkApp", []string{}}
	g.instrSchemes[cUnwind] = &InstrScheme{"IUnwind", []string{}}
	g.instrSchemes[cUpdate] = &InstrScheme{"IUpdate", []string{"n"}}
	g.instrSchemes[cPack] = &InstrScheme{"IPack", []string{"t", "n"}}
	g.instrSchemes[cSplit] = &InstrScheme{"ISplit", []string{}}
	g.instrSchemes[cJump] = &InstrScheme{"IJump", []string{"m"}}
	g.instrSchemes[cSlide] = &InstrScheme{"ISlide", []string{"n"}}
	g.instrSchemes[cBinOp] = &InstrScheme{"IBinOp", []string{"op"}}
	g.instrSchemes[cEval] = &InstrScheme{"IEval", []string{}}
	g.instrSchemes[cAlloc] = &InstrScheme{"IAlloc", []string{"n"}}
	g.instrSchemes[cPop] = &InstrScheme{"IPop", []string{"n"}}
}

const (
	cPushInt = iota
	cPushString
	cPushGlobal
	cPush
	cMkApp
	cUnwind
	cUpdate
	cPack
	cSplit
	cJump
	cSlide
	cBinOp
	cEval
	cAlloc
	cPop
)

/////////////////////////////////////// CASE MAP ///////////////////////////////////////////

type CaseMap struct {
	m           map[int][]*Instr
	defaultCode []*Instr
}

func createCaseMap() *CaseMap {
	m := new(CaseMap)
	m.m = make(map[int][]*Instr)

	return m
}

func (m *CaseMap) addAlt(i int, code []*Instr) {
	_, defined := m.m[i]
	if !defined {
		m.m[i] = code
	}
}

func (m *CaseMap) show() string {
	var res strings.Builder

	res.WriteString("map[int][]GInstr {\n")

	for i, code := range m.m {
		res.WriteString(fmt.Sprintf("%d: ", i))
		res.WriteString(showCode(code))
		res.WriteString(",\n")
	}

	if m.defaultCode != nil {
		res.WriteString("-1: ")
		res.WriteString(showCode(m.defaultCode))
		res.WriteString(",\n")
	}

	res.WriteString("}")
	return res.String()
}

//////////////////////////////////// CODE GENERATION ///////////////////////////////////////

func (g *CodeGenerator) genTopDecls(v TopDecls) {
	g.genTopDeclsList(v.(*TopDeclsList))
	g.iteration++
	g.genTopDeclsList(v.(*TopDeclsList))
}

func (g *CodeGenerator) genTopDeclsList(v *TopDeclsList) {
	for _, d := range v.decls {
		g.genTopDecl(d)
	}
}

func (g *CodeGenerator) genTopDecl(v TopDecl) {
	switch v := v.(type) {
	case *ImportTopDecl:
		return
	case *DataTopDecl:
		g.genDataTopDecl(v)
	case *FunTopDecl:
		g.genFunTopDecl(v)
	}
}

func (g *CodeGenerator) genDataTopDecl(v *DataTopDecl) {
	if g.iteration != 0 {
		return
	}

	constrs := v.constrs
	constrList := constrs.(*ConstrList)
	for i, def := range constrList.defs {
		constrType := def.(*ConstrType)
		constrName := constrType.constr
		g.constrTags[constrName] = i
	}
}

func (g *CodeGenerator) genFunTopDecl(v *FunTopDecl) {
	if g.iteration != 1 {
		return
	}

	decl := v.decl
	switch decl := decl.(type) {
	case *FunTypeDecl:
		return
	case *FunDecl:
		g.genFunDecl(decl)
	case *VarDecl:
		g.genVarDecl(decl)
	}
}

func (g *CodeGenerator) genFunDecl(v *FunDecl) {
	declLhs := v.lhs.(*DeclLhs)
	funName := declLhs.fun
	isIO := v.getIsIO()

	g.reset()

	var vars []string
	for _, arg := range declLhs.args {
		variable := arg.(*APatVar)
		vars = append(vars, variable.id)
	}
	g.env.applyChange(vars)

	g.genRhs(v.rhs)

	argsCount := len(declLhs.args)

	if isIO {
		g.pushInstr(cSlide, fmt.Sprintf("%d", argsCount+1))

		// We need it to prevent from exiting the program
		g.pushInstr(cUnwind)
	} else {
		g.pushInstr(cUpdate, fmt.Sprintf("%d", argsCount))
		g.pushInstr(cPop, fmt.Sprintf("%d", argsCount))
	}

	g.functionArities[funName] = argsCount
	g.saveFunction(funName)
}

func (g *CodeGenerator) genVarDecl(v *VarDecl) {
	patArg := v.pat.(*PatArg)
	aPat := patArg.arg
	variable := aPat.(*APatVar)
	funName := variable.id
	isIO := v.getIsIO()

	g.reset()
	g.genRhs(v.rhs)

	if isIO {
		g.pushInstr(cSlide, "1")

		// We need it to prevent from exiting the program
		g.pushInstr(cUnwind)
	} else {
		g.pushInstr(cUpdate, "0")
	}

	g.functionArities[funName] = 0
	g.saveFunction(funName)
}

func (g *CodeGenerator) genRhs(v Rhs) {
	exp := v.(*DeclExp)
	g.genExp(exp.e)
}

func (g *CodeGenerator) genExp(v Exp) {
	switch v := v.(type) {
	case *EFun:
		g.genEFun(v)
	case *EDo:
		g.genEDo(v)
	case *ECase:
		g.genECase(v)
	case *EMul:
		g.genEBinary(v, "*")
	case *EDiv:
		g.genEBinary(v, "/")
	case *EAdd:
		g.genEBinary(v, "+")
	case *ESub:
		g.genEBinary(v, "-")
	case *EBit:
		g.genEBit(v)
	case *EComp:
		g.genEComp(v)
	case *ELogical:
		g.genELogical(v)
	}
}

func (g *CodeGenerator) genEFun(v *EFun) {
	g.genFExp(v.fExp, 0)
}

func (g *CodeGenerator) genEDo(v *EDo) {
	g.genStmts(v.stmts)
}

func (g *CodeGenerator) genECase(v *ECase) {
	g.genExp(v.e)

	caseMap := createCaseMap()
	for _, alt := range v.alts {
		i, code := g.genAlternative(alt)
		caseMap.addAlt(i, code)
	}

	g.pushInstr(cEval)
	g.pushInstr(cJump, caseMap.show())
	// After performing Jump the constructor address at the stack is replaced with the result address,
	// so we don't have to worry about changing stack size
}

func (g *CodeGenerator) genEBinary(v EBinary, op string) {
	e1, e2 := v.getExps()
	g.genExp(e2)
	g.genExp(e1)

	// We need to print the operator in quotes
	g.pushInstr(cPushGlobal, fmt.Sprintf("%q", op))
	g.pushInstr(cMkApp)
	g.pushInstr(cMkApp)
	g.env.changeStackSize(-1) // Two args are replaced by one result
}

func (g *CodeGenerator) genEBit(v *EBit) {
	g.genExp(v.e2)
	g.genExp(v.e1)

	// We need to print the operator in quotes
	g.pushInstr(cPushGlobal, fmt.Sprintf("%q", v.op))
	g.pushInstr(cMkApp)
	g.pushInstr(cMkApp)
	g.env.changeStackSize(-1) // Two args are replaced by one result
}

func (g *CodeGenerator) genEComp(v *EComp) {
	g.genExp(v.e2)
	g.genExp(v.e1)

	// We need to print the operator in quotes
	g.pushInstr(cPushGlobal, fmt.Sprintf("%q", v.op))
	g.pushInstr(cMkApp)
	g.pushInstr(cMkApp)
	g.env.changeStackSize(-1) // Two args are replaced by one result
}

func (g *CodeGenerator) genELogical(v *ELogical) {
	g.genExp(v.e2)
	g.genExp(v.e1)

	// We need to print the operator in quotes
	g.pushInstr(cPushGlobal, fmt.Sprintf("%q", v.op))
	g.pushInstr(cMkApp)
	g.pushInstr(cMkApp)
	g.env.changeStackSize(-1) // Two args are replaced by one result
}

func (g *CodeGenerator) genFExp(v FExp, argsCount int) {
	switch v := v.(type) {
	case *FApp:
		g.genFApp(v, argsCount)
	case *FArg:
		g.genFArg(v, argsCount)
	}
}

func (g *CodeGenerator) genFApp(v *FApp, argsCount int) {
	g.genAExp(v.arg, 0)
	g.genFExp(v.fun, argsCount+1)
}

func (g *CodeGenerator) genFArg(v *FArg, argsCount int) {
	g.genAExp(v.arg, argsCount)
}

func (g *CodeGenerator) genAExp(v AExp, argsCount int) {
	switch v := v.(type) {
	case *Var:
		if g.env.varDefined(v.id) {
			// `v` is a variable
			offset := g.env.getVar(v.id)
			g.pushInstr(cPush, fmt.Sprintf("%d", offset))
			g.env.changeStackSize(1)
		} else {
			// `v` is a global function
			g.pushInstr(cPushGlobal, fmt.Sprintf("%q", v.id))
			g.env.changeStackSize(1)
		}
	case *Constr:
		tag := g.constrTags[v.id]
		g.pushInstr(cPack, fmt.Sprintf("%d", tag), fmt.Sprintf("%d", argsCount))
		g.env.changeStackSize(1 - argsCount)
		return
	case *Lit:
		g.genLiteral(v.lit)
		return
	case *ParenExp:
		g.genExp(v.e)
	case *Tuple:
		errFatal(v, "Tuples not supported by runtime yet")
	}

	for i := 0; i < argsCount; i++ {
		g.pushInstr(cMkApp)
		g.env.changeStackSize(-1)
	}
}

func (g *CodeGenerator) genLiteral(v Literal) {
	switch v := v.(type) {
	case *Int:
		g.pushInstr(cPushInt, fmt.Sprintf("%d", v.n))
		g.env.changeStackSize(1)
	case *Char:
		errFatal(v, "Chars not supported by runtime yet")
	case *String:
		g.pushInstr(cPushString, fmt.Sprintf("%q", v.s))
		g.env.changeStackSize(1)
	}
}

func (g *CodeGenerator) genStmts(v Stmts) {
	stmtsList := v.(*StmtsList)

	var backups []*EnvBackup
	envChanges := 0

	for _, stmt := range stmtsList.statements {
		backup := g.genStmt(stmt)
		backups = append(backups, backup)
		envChanges += backup.size()
	}
	g.genExp(stmtsList.exp)

	// We omit the following eval to get effective tail calls.
	// The last expression is evaluated later (in many cases after cleaning the stack at the end of a function)
	// g.pushInstr(cEval)

	// Remove all variables from do from the stack
	g.pushInstr(cSlide, fmt.Sprintf("%d", envChanges))

	for i := len(backups) - 1; i >= 0; i-- {
		g.env.revertChange(backups[i])
	}
}

func (g *CodeGenerator) genStmt(v Stmt) *EnvBackup {
	switch v := v.(type) {
	case *SExp:
		return g.genSExp(v)
	case *SAssign:
		return g.genSAssign(v)
	}

	return &EnvBackup{changed: make(map[string]int)}
}

func (g *CodeGenerator) genSExp(v *SExp) *EnvBackup {
	g.genExp(v.exp)
	g.pushInstr(cEval)

	// Even if exp is of type `IO ()`, evaluating it leaves unit value at the stack
	g.pushInstr(cPop, "1")
	g.env.changeStackSize(-1)

	return &EnvBackup{changed: make(map[string]int)}
}

func (g *CodeGenerator) genSAssign(v *SAssign) *EnvBackup {
	varName := v.pat.(*PatArg).arg.(*APatVar).id

	g.genExp(v.exp)
	g.pushInstr(cEval)

	// We need to decrease stored stack size, to save the correct variable location
	g.env.changeStackSize(-1)
	backup := g.env.applyChange([]string{varName})

	return backup
}

func (g *CodeGenerator) genAlternative(v *Alternative) (int, []*Instr) {
	g.saveCode()

	// We split the data, so before putting pattern vars at the stack we remove the structure
	g.env.changeStackSize(-1)
	tag, backup := g.genPat(v.pat)

	if tag >= 0 {
		g.pushInstr(cSplit)
	}
	g.genExp(v.exp)
	if tag >= 0 {
		g.pushInstr(cSlide, fmt.Sprintf("%d", backup.size()))
	}

	code := g.code
	g.env.revertChange(backup)

	g.restoreCode()
	return tag, code
}

func (g *CodeGenerator) genPat(v Pat) (int, *EnvBackup) {
	switch v := v.(type) {
	case *PatArg:
		return g.genPatArg(v)
	case *PatCon:
		return g.genPatCon(v)
	}
	return 0, nil
}

func (g *CodeGenerator) genPatArg(v *PatArg) (int, *EnvBackup) {
	switch arg := v.arg.(type) {
	case *APatVar:
		backup := g.env.applyChange([]string{arg.id})
		return -1, backup
	case *APatCon:
		backup := g.env.applyChange([]string{})
		tag := g.constrTags[arg.id]
		return tag, backup
	case *APatLit:
		errFatal(v, "Literals in patterns not supported by runtime yet")
	}
	return 0, nil
}

func (g *CodeGenerator) genPatCon(v *PatCon) (int, *EnvBackup) {
	tag := g.constrTags[v.conId]
	var args []string

	for _, arg := range v.args {
		switch arg := arg.(type) {
		case *APatVar:
			args = append(args, arg.id)
		case *APatCon:
			errFatal(v, "Complex patterns not supported")
		case *APatLit:
			errFatal(v, "Literals in patterns not supported by runtime yet")
		}
	}

	backup := g.env.applyChange(args)
	return tag, backup
}

//////////////////////////////////// OUTPUT CODE ///////////////////////////////////////////

func (g *CodeGenerator) genOutput() string {
	var res strings.Builder

	res.WriteString("package main\n\n")
	res.WriteString("func main() {\n")

	res.WriteString("gMachine := NewGMachine()\n")

	res.WriteString("initialInstructions := []GInstr{\n")
	res.WriteString("&IPushGlobal{f: \"main\"},\n")
	res.WriteString("&IUnwind{},\n")
	res.WriteString("}\n")
	res.WriteString("gMachine.instrQueue.putN(initialInstructions)\n\n")

	for f, code := range g.functionDefs {
		res.WriteString(fmt.Sprintf("gMachine.addGlobal(%q, %d, %s)\n", f, g.functionArities[f], showCode(code)))
	}

	res.WriteString("gMachine.run()\n")

	res.WriteString("}\n")

	return res.String()
}

package main

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Base node

type ASTNode interface {
	astNode()
	getPos() (int, int)
	setPos(int, int)
	setPosFromCtx(antlr.ParserRuleContext)
}

type BaseASTNode struct {
	posLine int
	posCol  int
}

func (BaseASTNode) astNode() {}

func (node *BaseASTNode) getPos() (int, int) {
	return node.posLine, node.posCol
}
func (node *BaseASTNode) setPos(line int, col int) {
	node.posLine = line
	node.posCol = col + 1 // change counting from 0 to counting from 1
}

func (node *BaseASTNode) setPosFromCtx(ctx antlr.ParserRuleContext) {
	line := ctx.GetStart().GetLine()
	col := ctx.GetStart().GetColumn()

	node.setPos(line, col)
}

// topdecls

type TopDecls interface {
	ASTNode
	topDecls()
}

type BaseTopDecls struct {
	BaseASTNode
}

func (*BaseTopDecls) topDecls() {}

type TopDeclsList struct {
	BaseTopDecls
	decls []TopDecl
}

// topdecl

type TopDecl interface {
	ASTNode
	topDecl()
}

type BaseTopDecl struct {
	BaseASTNode
}

func (*BaseTopDecl) topDecl() {}

type DataTopDecl struct {
	BaseTopDecl
	t       SimpleType
	constrs Constrs
}

type FunTopDecl struct {
	BaseTopDecl
	decl Decl
}

// simpletype

type SimpleType interface {
	ASTNode
	simpleType()
}

type BaseSimpleType struct {
	BaseASTNode
}

func (*BaseSimpleType) simpleType() {}

type DataType struct {
	BaseSimpleType
	constr string
	vars   []string
}

// constrs

type Constrs interface {
	ASTNode
	constrs()
}

type BaseConstrs struct {
	BaseASTNode
}

func (*BaseConstrs) constrs() {}

type ConstrList struct {
	BaseConstrs
	defs []ConstrDef
}

// constrdef

type ConstrDef interface {
	ASTNode
	constrDef()
}

type BaseConstrDef struct {
	BaseASTNode
}

func (*BaseConstrDef) constrDef() {}

type ConstrType struct {
	BaseConstrDef
	constr string
	args   []AType
}

// decl

type Decl interface {
	ASTNode
	decl()
}

type BaseDecl struct {
	BaseASTNode
}

func (*BaseDecl) decl() {}

type FunTypeDecl struct {
	BaseDecl
	d GenDecl
}

type FunDecl struct {
	BaseDecl
	lhs FunLhs
	rhs Rhs
}

type VarDecl struct {
	BaseDecl
	pat Pat
	rhs Rhs
}

// gendecl

type GenDecl interface {
	ASTNode
	genDecl()
}

type BaseGenDecl struct {
	BaseASTNode
}

func (*BaseGenDecl) genDecl() {}

type TypeSignature struct {
	BaseGenDecl
	vars []string
	t    Type
}

// type

type Type interface {
	ASTNode
	type_()
}

type BaseType struct {
	BaseASTNode
}

func (*BaseType) type_() {}

type FunType struct {
	BaseType
	types []BType
}

// btype

type BType interface {
	ASTNode
	bType()
}

type BaseBType struct {
	BaseASTNode
}

func (*BaseBType) bType() {}

type TypeApp struct {
	BaseBType
	types []AType
}

// atype

type AType interface {
	ASTNode
	aType()
}

type BaseAType struct {
	BaseASTNode
}

func (*BaseAType) aType() {}

type ConType struct {
	BaseAType
	id string
}

type VarType struct {
	BaseAType
	id string
}

type TupleType struct {
	BaseAType
	types []Type
}

type ParenType struct {
	BaseAType
	t Type
}

// funlhs

type FunLhs interface {
	ASTNode
	funLhs()
}

type BaseFunLhs struct {
	BaseASTNode
}

func (*BaseFunLhs) funLhs() {}

type DeclLhs struct {
	BaseFunLhs
	fun  string
	args []APat
}

// rhs

type Rhs interface {
	ASTNode
	rhs()
}

type BaseRhs struct {
	BaseASTNode
}

func (*BaseRhs) rhs() {}

type DeclExp struct {
	BaseRhs
	e Exp
}

// exp

type Exp interface {
	ASTNode
	exp()
}

type BaseExp struct {
	BaseASTNode
}

func (*BaseExp) exp() {}

type EFun struct {
	BaseExp
	fExp FExp
}

type ECase struct {
	BaseExp
	e    Exp
	alts []*Alternative
}

type EBinary interface {
	Exp
	eBinary()
	setExps(Exp, Exp)
	getExps() (Exp, Exp)
}

type BaseEBinary struct {
	BaseExp
	e1 Exp
	e2 Exp
}

func (*BaseEBinary) eBinary() {}

func (e *BaseEBinary) setExps(e1 Exp, e2 Exp) {
	e.e1 = e1
	e.e2 = e2
}

func (e *BaseEBinary) getExps() (Exp, Exp) {
	return e.e1, e.e2
}

type EMul struct {
	BaseEBinary
}

type EDiv struct {
	BaseEBinary
}

type EAdd struct {
	BaseEBinary
}

type ESub struct {
	BaseEBinary
}

type EComp struct {
	BaseExp
	e1 Exp
	e2 Exp
	op string
}

type ELogical struct {
	BaseExp
	e1 Exp
	e2 Exp
	op string
}

// fExp

type FExp interface {
	ASTNode
	fExp()
}

type BaseFExp struct {
	BaseASTNode
}

func (*BaseFExp) fExp() {}

type FApp struct {
	BaseFExp
	fun FExp
	arg AExp
}

type FArg struct {
	BaseFExp
	arg AExp
}

// aexp

type AExp interface {
	ASTNode
	aExp()
}

type BaseAExp struct {
	BaseASTNode
}

func (*BaseAExp) aExp() {}

type Var struct {
	BaseAExp
	id string
}

type Constr struct {
	BaseAExp
	id string
}

type Lit struct {
	BaseAExp
	lit Literal
}

type ParenExp struct {
	BaseAExp
	e Exp
}

type Tuple struct {
	BaseAExp
	exps []Exp
}

// alt

type Alternative struct {
	BaseASTNode
	pat Pat
	exp Exp
}

// pat

type Pat interface {
	ASTNode
	pat()
}

type BasePat struct {
	BaseASTNode
}

func (*BasePat) pat() {}

type PatArg struct {
	BasePat
	arg APat
}

type PatCon struct {
	BasePat
	conId string
	args  []APat
}

// apat

type APat interface {
	ASTNode
	aPat()
}

type BaseAPat struct {
	BaseASTNode
}

func (*BaseAPat) aPat() {}

type APatVar struct {
	BaseAPat
	id string
}

type APatCon struct {
	BaseAPat
	id string
}

type APatLit struct {
	BaseAPat
	lit Literal
}

// literal

type Literal interface {
	ASTNode
	literal()
}

type BaseLiteral struct {
	BaseASTNode
}

func (*BaseLiteral) literal() {}

type Int struct {
	BaseLiteral
	n int
}

type Char struct {
	BaseLiteral
	c byte
}

type String struct {
	BaseLiteral
	s string
}

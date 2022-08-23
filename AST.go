package main

// Base node

type ASTNode interface {
	astNode()
	getPos() (int, int)
	setPos(int, int)
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
}

type BaseEBinary struct {
	BaseExp
	e1 Exp
	e2 Exp
}

func (*BaseEBinary) eBinary() {}

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

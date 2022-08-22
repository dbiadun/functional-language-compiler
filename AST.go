package main

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

type Expr interface {
	ASTNode
	expr()
}

type BaseExpr struct {
	BaseASTNode
}

func (*BaseExpr) expr() {}

type EBinary interface {
	Expr
	eBinary()
}

type BaseEBinary struct {
	BaseExpr
	e1 Expr
	e2 Expr
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

type Number struct {
	BaseExpr
	n int
}

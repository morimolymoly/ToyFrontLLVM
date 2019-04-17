package ast

type BinaryAst struct {
	name string
	op   int
	lhs  ExprAST
	rhs  ExprAST
}

func NewBinaryAst(op int, lhs ExprAST, rhs ExprAST) BinaryAst {
	return BinaryAst{op: op, lhs: lhs, rhs: rhs}
}

func (a BinaryAst) Mock() {
}

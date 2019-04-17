package ast

type CallExprAST struct {
	callee string
	args   []ExprAST
}

func NewCallExprAST(callee string, args []ExprAST) *CallExprAST {
	return &CallExprAST{callee: callee, args: args}
}

func (a CallExprAST) Mock() {
}

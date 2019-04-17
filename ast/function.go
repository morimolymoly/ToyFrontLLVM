package ast

type FunctionAST struct {
	proto PrototypeAST
	body  ExprAST
}

func NewFunctionAST(proto PrototypeAST, body ExprAST) FunctionAST {
	return FunctionAST{proto: proto, body: body}
}

func (a FunctionAST) Mock() {

}

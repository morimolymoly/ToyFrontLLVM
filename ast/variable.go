package ast

type VariableAst struct {
	name string
}

func NewVariableAst(name string) VariableAst {
	return VariableAst{name: name}
}

func (a VariableAst) Mock() {
}

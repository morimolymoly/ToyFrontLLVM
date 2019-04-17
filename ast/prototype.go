package ast

type PrototypeAST struct {
	name string
	args []string
}

func NewPrototypeAST(name string, args []string) PrototypeAST {
	return PrototypeAST{name: name, args: args}
}

func (a PrototypeAST) GetName() string {
	return a.name
}

func (a PrototypeAST) Mock() {

}

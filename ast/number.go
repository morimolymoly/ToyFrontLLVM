package ast

type NumberAst struct {
	val float64
}

func NewNumberAst(val float64) NumberAst {
	return NumberAst{val: val}
}

func (a NumberAst) Mock() {
}

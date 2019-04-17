package parser

import (
	"fmt"

	"github.com/morimolymoly/ToyFrontLLVM/input"
	"github.com/morimolymoly/ToyFrontLLVM/lex"
)

type Handler struct {
	p *Parser
}

func NewHandler(i input.Input) Handler {
	return Handler{
		NewParser(i),
	}
}

func (h Handler) HandleDef() {
	if _, ok := h.p.ParseDef(); ok {
		fmt.Println("HANDLE DEF")
	} else {
		h.p.GetNextToken()
	}
}
func (h Handler) HandleExtern() {
	if _, ok := h.p.ParseExtern(); ok {
		fmt.Println("HANDLE EXTERN")
	} else {
		h.p.GetNextToken()
	}
}

func (h Handler) HandleTopLevelExpr() {
	if _, ok := h.p.ParseTopLevelExpr(); ok {
		fmt.Println("HANDLE TOPLEVEL EXPR")
	} else {
		h.p.GetNextToken()
	}
}

func (h Handler) MainLoop() {
	h.p.GetNextToken()
	for {
		switch h.p.currentToken {
		case lex.TokenEof:
			return
		case ';':
			h.p.GetNextToken()
			continue
		case lex.TokenDef:
			h.HandleDef()
		case lex.TokenExtern:
			h.HandleExtern()
		default:
			h.HandleTopLevelExpr()
		}
	}
}

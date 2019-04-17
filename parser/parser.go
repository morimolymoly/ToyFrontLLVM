package parser

import (
	"github.com/morimolymoly/ToyFrontLLVM/ast"
	"github.com/morimolymoly/ToyFrontLLVM/input"
	"github.com/morimolymoly/ToyFrontLLVM/lex"
)

type Parser struct {
	currentToken int
	input        input.Input
}

func NewParser(i input.Input) *Parser {
	return &Parser{input: i}
}

func (p *Parser) GetNextToken() {
	p.currentToken = lex.GetToken(p.input)
}

func (p *Parser) parseNumberExpr() (ast.NumberAst, bool) {
	r := ast.NewNumberAst(lex.NumVal)
	p.GetNextToken()
	return r, true
}

func (p *Parser) ParseParenExpr() (ast.ExprAST, bool) {
	p.GetNextToken() // '('
	v, ok := p.ParseExpression()
	if !ok {
		return nil, false
	}
	if p.currentToken != ')' {
		panic("expected )")
	}
	p.GetNextToken()
	return v, true //v
}

// identifier
// identifer () <=function call
func (p *Parser) ParseIdentifierExpr() (ast.ExprAST, bool) {
	Idname := lex.IdentifierStr
	p.GetNextToken()

	if p.currentToken != '(' {
		return ast.NewVariableAst(Idname), true
	}

	p.GetNextToken()
	args := []ast.ExprAST{}

	if p.currentToken != ')' {
		for {
			if arg, ok := p.ParseExpression(); ok {
				args = append(args, arg)
			} else {
				return nil, false
			}

			if p.currentToken == ')' {
				break
			}
			if p.currentToken != ',' {
				panic("")
			}
			p.GetNextToken()
		}
	}
	p.GetNextToken()
	return ast.NewCallExprAST(Idname, args), true
}

func (p *Parser) getTokenPrec() int {
	if p.currentToken > 127 || p.currentToken <= 0 {
		return -1
	}
	switch p.currentToken {
	default:
		return -1
	case '<':
		return 10
	case '+':
		return 20
	case '-':
		return 20
	case '*':
		return 40
	}
}

func (p *Parser) ParseExpression() (ast.ExprAST, bool) {
	LHS, ok := p.ParsePrimary()
	if !ok {
		return nil, false
	}
	return p.ParseBinOpRHS(0, LHS)
}

func (p *Parser) ParseBinOpRHS(execPrec int, LHS ast.ExprAST) (ast.ExprAST, bool) {
	for {
		currentPrec := p.getTokenPrec()
		if currentPrec < execPrec {
			return LHS, true
		}

		binOP := p.currentToken
		p.GetNextToken()

		RHS, ok := p.ParsePrimary()
		if !ok {
			return nil, false
		}
		nextPrec := p.getTokenPrec()
		if currentPrec < nextPrec {
			RHS, ok = p.ParseBinOpRHS(currentPrec+1, RHS)
			if !ok {
				return nil, false
			}
		}
		LHS = ast.NewBinaryAst(binOP, LHS, RHS)
	}
}

func (p *Parser) ParseExtern() (ast.PrototypeAST, bool) {
	p.GetNextToken()
	return p.ParsePrototype()
}

func (p *Parser) ParsePrototype() (ast.PrototypeAST, bool) {
	if p.currentToken != lex.TokenIdentifier {
		panic("")
	}

	functionName := lex.IdentifierStr
	p.GetNextToken()
	if p.currentToken != '(' {
		panic("")
	}

	args := []string{}
	for {
		if p.GetNextToken(); p.currentToken == lex.TokenIdentifier {
			args = append(args, lex.IdentifierStr)
			continue
		}
		break
	}
	if p.currentToken != ')' {
		panic("")
	}
	p.GetNextToken()
	return ast.NewPrototypeAST(functionName, args), true
}

func (p *Parser) ParseDef() (ast.FunctionAST, bool) {
	p.GetNextToken()
	proto, ok := p.ParsePrototype()
	if !ok {
		return ast.NewFunctionAST(ast.NewPrototypeAST("", nil), ast.NewNumberAst(1)), false
	}

	exp, ok := p.ParseExpression()
	if !ok {
		return ast.NewFunctionAST(ast.NewPrototypeAST("", nil), ast.NewNumberAst(1)), false
	}
	return ast.NewFunctionAST(proto, exp), true
}

func (p *Parser) ParseTopLevelExpr() (ast.FunctionAST, bool) {
	e, ok := p.ParseExpression()
	proto := ast.NewPrototypeAST("", []string{})
	if ok {
		return ast.NewFunctionAST(proto, e), true
	}
	return ast.NewFunctionAST(proto, e), false
}

func (p *Parser) ParsePrimary() (ast.ExprAST, bool) {
	switch p.currentToken {
	case lex.TokenIdentifier:
		return p.ParseIdentifierExpr()
	case lex.TokenNumber:
		return p.parseNumberExpr()
	case '(':
		return p.ParseParenExpr()
	}
	return nil, false
}

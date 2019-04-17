package lex

import (
	"strconv"

	"github.com/morimolymoly/ToyFrontLLVM/input"
	u "github.com/morimolymoly/ToyFrontLLVM/utils"
)

const (
	TokenEof int = iota
	TokenDef
	TokenExtern
	TokenIdentifier
	TokenNumber
)

var NumVal float64
var IdentifierStr string

func GetToken(i input.Input) int {
	var lastChar int
	lastChar = ' '
	var err error

	for {
		if lastChar != ' ' {
			break
		}
		lastChar, err = i.GetChar()
	}
	if u.IsAlpha(lastChar) {
		IdentifierStr = string(lastChar)
		for {
			lastChar, err = i.GetChar()
			if u.IsAlphaOrNum(lastChar) {
				IdentifierStr += string(lastChar)
				continue
			}
			break
		}
		if IdentifierStr == "def" {
			return TokenDef
		} else if IdentifierStr == "extern" {
			return TokenExtern
		}
		return TokenIdentifier
	}
	if u.IsDigit(lastChar) || lastChar == '.' {
		numString := ""
		numString += string(lastChar)
		for {
			lastChar, err = i.GetChar()
			if u.IsDigit(lastChar) || lastChar == '.' {
				numString += string(lastChar)
				continue
			}
			break
		}
		f, err := strconv.ParseFloat(numString, 64)
		if err != nil {
			panic(err)
		}
		NumVal = f
		return TokenNumber
	}

	if lastChar == '#' {
		i.ReadLine()
		return GetToken(i)
	}
	if lastChar == '\n' || lastChar == '\r' {
		i.ReadLine()
		return GetToken(i)
	}

	if err == input.EOF {
		return TokenEof
	}
	return lastChar
}

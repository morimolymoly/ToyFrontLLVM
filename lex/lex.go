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

func GetToken(i input.Input) int {
	var lastChar int
	lastChar = ' '
	identifiedChar := ""
	var err error

	for {
		if lastChar != ' ' {
			break
		}
		lastChar, err = i.GetChar()
	}
	if u.IsAlpha(lastChar) {
		identifiedChar = string(lastChar)
		for {
			lastChar, err = i.GetChar()
			if u.IsAlphaOrNum(lastChar) {
				identifiedChar += string(lastChar)
				continue
			}
			break
		}
		if identifiedChar == "def" {
			return TokenDef
		} else if identifiedChar == "extern" {
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
		if !i.ReadLine() {
			return TokenEof
		}
		return GetToken(i)
	}
	if lastChar == '\n' || lastChar == '\r' {
		if !i.ReadLine() {
			return TokenEof
		}
		return GetToken(i)
	}

	if err == input.EOF {
		return TokenEof
	}
	return lastChar
}

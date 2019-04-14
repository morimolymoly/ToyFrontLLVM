package lex

import (
	"bytes"
	"testing"

	"github.com/morimolymoly/ToyFrontLLVM/input"
)

func TestRex(t *testing.T) {
	{
		i := input.NewStdin(bytes.NewBufferString("def unko"))
		i.ReadLine()
		r := GetToken(i)
		if r != TokenDef {
			t.Fatal("token should be a def")
		}
	}
	{
		i := input.NewStdin(bytes.NewBufferString("extern shit"))
		i.ReadLine()
		r := GetToken(i)
		if r != TokenExtern {
			t.Fatal("token should be a Extern")
		}
	}
	{
		i := input.NewStdin(bytes.NewBufferString("128.1"))
		i.ReadLine()
		r := GetToken(i)
		if r != TokenNumber {
			t.Fatal("token should be a TokenNumber")
		}
	}
	{
		i := input.NewStdin(bytes.NewBufferString("superman"))
		i.ReadLine()
		r := GetToken(i)
		if r != TokenIdentifier {
			t.Fatal("token should be a TokenIdentifier")
		}
	}
	{
		i := input.NewStdin(bytes.NewBufferString("# zakozako"))
		i.ReadLine()
		r := GetToken(i)
		if r != TokenEof {
			t.Fatal("token should be a EOF")
		}
	}
	{
		i := input.NewStdin(bytes.NewBufferString("# zakozako\nsuperman"))
		i.ReadLine()
		r := GetToken(i)
		if r != TokenIdentifier {
			t.Fatal("token should be a TokenIdentifier")
		}
	}
	{
		i := input.NewStdin(bytes.NewBufferString("+"))
		i.ReadLine()
		r := GetToken(i)
		if r != '+' {
			t.Fatal("token should be a +")
		}
	}
	{
		i := input.NewStdin(bytes.NewBufferString("#zakozako\n\n+"))
		i.ReadLine()
		r := GetToken(i)
		if r == '+' {
			t.Fatal("Empty Line is not allowed!")
		}
	}
}

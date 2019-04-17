package main

import (
	"bytes"
	"fmt"

	"github.com/morimolymoly/ToyFrontLLVM/input"
	"github.com/morimolymoly/ToyFrontLLVM/parser"
)

func main() {
	fmt.Println("HelloWorld!")
	{
		i := input.NewStdin(bytes.NewBufferString("def uoo ( x y ) 1+4"))
		i.ReadLine()
		h := parser.NewHandler(i)
		h.MainLoop()
	}
	{
		i := input.NewStdin(bytes.NewBufferString("1 + 10"))
		i.ReadLine()
		h := parser.NewHandler(i)
		h.MainLoop()
	}
}

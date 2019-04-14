package input

import (
	"bufio"
	"bytes"
	"errors"
	"os"
)

var EOF = errors.New("EOF")

type Input interface {
	GetChar() (int, error)
	ReadLine() bool
}

type Stdin struct {
	s   *bufio.Scanner
	pos int
}

func NewStdin(buf *bytes.Buffer) *Stdin {
	s := &Stdin{}
	if buf != nil {
		s.s = bufio.NewScanner(buf)
		return s
	}

	s.s = bufio.NewScanner(os.Stdin)
	return s
}

func (s *Stdin) GetChar() (int, error) {
	defer func() {
		s.pos++
	}()

	if s.pos >= len(s.s.Text()) {
		return 0, EOF
	}

	return int(s.s.Text()[s.pos]), nil
}

func (s *Stdin) ReadLine() bool {
	s.pos = 0
	return s.s.Scan()
}

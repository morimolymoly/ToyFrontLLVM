package input

import (
	"bytes"
	"testing"
)

func TestGetChar(t *testing.T) {
	i := NewStdin(bytes.NewBufferString("abcde"))

	if r := i.ReadLine(); !r {
		t.Fatal("NOOOO")
	}

	o, e := i.GetChar()
	if 'a' != o && e != nil {
		t.Fatalf("char shoule be a! not %d %s %s", o, e, i.s.Text())
	}
	o, e = i.GetChar()
	if 'b' != o && e != nil {
		t.Fatalf("char shoule be b! not %d %s", o, e)
	}
	o, e = i.GetChar()
	if 'c' != o && e != nil {
		t.Fatalf("char shoule be c! not %d %s", o, e)
	}
	o, e = i.GetChar()
	if 'd' != o && e != nil {
		t.Fatalf("char shoule be d! not %d %s", o, e)
	}
	o, e = i.GetChar()
	if 'e' != o && e != nil {
		t.Fatalf("char shoule be e! not %d %s", o, e)
	}

	_, e = i.GetChar()
	if e != EOF {
		t.Fatalf("error EOF should be returned")
	}
}

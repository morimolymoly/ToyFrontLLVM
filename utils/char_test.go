package utils

import "testing"

func TestisAlpha(t *testing.T) {
	res := IsAlpha('a')
	if !res {
		t.Fatal("0 should be treated as a alpha")
	}
	res = IsAlpha('Z')
	if !res {
		t.Fatal("Z should be treated as a alpha")
	}
	res = IsAlpha('s')
	if !res {
		t.Fatal("s should be treated as a alpha")
	}
	res = IsAlpha('T')
	if !res {
		t.Fatal("T should be treated as a alpha")
	}

	res = IsAlpha('0')
	if res {
		t.Fatal("0 should not be treated as a alpha")
	}
	res = IsAlpha('6')
	if res {
		t.Fatal("6 should not be treated as a alpha")
	}
	res = IsAlpha('!')
	if res {
		t.Fatal("! should not be treated as a alpha")
	}
}

func TestisAlphaOrNum(t *testing.T) {
	res := IsAlphaOrNum(0)
	if !res {
		t.Fatal("0 should be treated as a alphaornum")
	}
	res = IsAlphaOrNum(5)
	if !res {
		t.Fatal("5 should be treated as a alphaornum")
	}
	res = IsAlphaOrNum('a')
	if !res {
		t.Fatal("0 should be treated as a alphaornum")
	}
	res = IsAlphaOrNum('Z')
	if !res {
		t.Fatal("Z should be treated as a alphaornum")
	}
	res = IsAlphaOrNum('s')
	if !res {
		t.Fatal("s should be treated as a alphaornum")
	}
	res = IsAlphaOrNum('T')
	if !res {
		t.Fatal("T should be treated as a alphaornum")
	}

	res = IsAlphaOrNum('!')
	if res {
		t.Fatal("! should not be treated as a alphaornum")
	}
}

func TestIsDigit(t *testing.T) {
	res := IsDigit('0')
	if !res {
		t.Fatal("0 should be treated as digit")
	}

	res = IsDigit('9')
	if !res {
		t.Fatal("9 should be treated as digit")
	}

	res = IsDigit('5')
	if !res {
		t.Fatal("5 should be treated as digit")
	}

	res = IsDigit('a')
	if res {
		t.Fatal("a should not be treated as digit")
	}
}

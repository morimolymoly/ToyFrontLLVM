package utils

func IsAlpha(char int) bool {
	if char >= 'A' && char <= 'z' {
		return true
	}
	return false
}

func IsAlphaOrNum(char int) bool {
	if '0' <= char && char >= 'A' && char <= 'z' {
		return true
	}
	return false
}

func IsDigit(char int) bool {
	if '0' <= char && char <= '9' {
		return true
	}
	return false
}

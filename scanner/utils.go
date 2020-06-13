package scanner

import (
	"unicode"
)

func isSpace(token rune) bool {
	return unicode.IsSpace(token)
}

func isLetter(token rune) bool {
	return unicode.IsLetter(token)
}

func isSymbol(token rune) bool {
	switch token {
	case '+', '-', '*', '/':
		return true
	default:
		return false
	}
}

func isDigit(token rune) bool {
	return unicode.IsDigit(token)
}

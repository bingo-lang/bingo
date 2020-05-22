package scanner

import (
	"unicode"
)

func isSpace(buffer rune) bool {
	return unicode.IsSpace(buffer)
}

func isLetter(buffer rune) bool {
	return unicode.IsLetter(buffer)
}

func isSymbol(buffer rune) bool {
	return buffer == '-'
}

func isDigit(buffer rune) bool {
	return unicode.IsDigit(buffer)
}

package tokenizer

import (
	"unicode"
)

func isSpace(char rune) bool {
	return unicode.IsSpace(char)
}

func isLetter(char rune) bool {
	return unicode.IsLetter(char)
}

func isSymbol(char rune) bool {
	return char == ';' || unicode.IsSymbol(char)
}

func isDigit(char rune) bool {
	return unicode.IsDigit(char)
}

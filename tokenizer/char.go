package tokenizer

import (
	"unicode/utf8"
)

func (t *Tokenizer) char() rune {
	c, _ := utf8.DecodeRuneInString(t.source[t.end:])
	return c
}

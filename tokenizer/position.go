package tokenizer

import (
	"unicode/utf8"
)

func (t *Tokenizer) advance() {
	_, w := utf8.DecodeRuneInString(t.source[t.end:])
	t.end += w
}
func (t *Tokenizer) confirm() {
	t.pos = t.end
}

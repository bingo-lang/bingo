package tokenizer

func (t *Tokenizer) next() string {
	if isLetter(t.char()) {
		return t.nextWord()
	}
	if isDigit(t.char()) {
		return t.nextNumber()
	}
	if isSymbol(t.char()) {
		return t.nextSymbol()
	}
	return ""
}

func (t *Tokenizer) nextSpace() string {
	space := ""
	for ; isSpace(t.char()); t.advance() {
	}
	t.confirm()
	return space
}

func (t *Tokenizer) nextWord() string {
	word := ""
	for ; isLetter(t.char()); t.advance() {
		word += string(t.char())
	}
	t.confirm()
	return word
}

func (t *Tokenizer) nextNumber() string {
	number := ""
	for ; isDigit(t.char()); t.advance() {
		number += string(t.char())
	}
	t.confirm()
	return number

}

func (t *Tokenizer) nextSymbol() string {
	symbol := string(t.char())
	t.advance()
	t.confirm()
	return symbol
}

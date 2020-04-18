package tokenizer

type Tokenizer struct {
	source string

	// position pointers
	pos int
	end int
}

func New(source string) (*Tokenizer, error) {
	return &Tokenizer{source: source}, nil
}

func (t *Tokenizer) Next() string {
	t.nextSpace()
	token := t.next()
	t.nextSpace()
	return token
}

func (t *Tokenizer) EOF() bool {
	return t.pos >= len(t.source)
}

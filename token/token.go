package token

import "fmt"

type Token struct {
	Type  Type
	Value string
}

func New(t Type, v string) Token {
	return Token{t, v}
}

func (t Token) String() string {
	return fmt.Sprintf("(%s, %s)", t.Type, t.Value)
}

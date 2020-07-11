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

func (t1 Token) Equals(t2 Token) bool {
	return t1.Type == t2.Type && t1.Value == t2.Value
}

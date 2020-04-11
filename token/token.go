package token

type Type uint

const (
	UNDEFINED Type = iota
)

type Token struct {
	Type  Type
	Value string
}

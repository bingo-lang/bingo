package token

type Type uint

const (
	UNDEFINED Type = iota
	EOF
	// Symbols.
	MINUS
	// Numbers.
	INTEGER
)

func (t Type) String() string {
	switch t {
	case EOF:
		return "EOF"
	//Symbols.
	case MINUS:
		return "MINUS"
	//Numbers.
	case INTEGER:
		return "INTEGER"
	default:
		return "UNDEFINED"
	}
}

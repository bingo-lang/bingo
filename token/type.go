package token

type Type uint

const (
	UNDEFINED Type = iota

	SPACE

	// Words.
	LET
	IDENTIFIER

	// Symbols.
	ASSIGNMENT
	SEMICOLON

	// Numbers.
	INTEGER
)

func (t Type) String() string {
	switch t {
	case SPACE:
		return "SPACE"

	case LET:
		return "LET"
	case IDENTIFIER:
		return "IDENTIFIER"

	case ASSIGNMENT:
		return "ASSIGNMENT"
	case SEMICOLON:
		return "SEMICOLON"

	case INTEGER:
		return "INTEGER"

	default:
		return "UNDEFINED"
	}
}

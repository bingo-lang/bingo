package token

type Type uint

const (
	UNDEFINED Type = iota
	EOF

	// Symbols.
	PLUS
	MINUS
	ASTERISK
	SLASH

	// Numbers.
	INTEGER
)

func (t Type) String() string {
	switch t {
	case EOF:
		return "EOF"

		// Symbols.
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case ASTERISK:
		return "ASTERISK"
	case SLASH:
		return "SLASH"

	// Numbers.
	case INTEGER:
		return "INTEGER"

	default:
		return "UNDEFINED"
	}
}

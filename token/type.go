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
	SEMICOLON

	// Numbers.
	INTEGER
)

func (t Type) String() string {
	switch t {
	case EOF:
		return "EOF"

	// Symbols.
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	case ASTERISK:
		return "*"
	case SLASH:
		return "/"
	case SEMICOLON:
		return ";"

	// Numbers.
	case INTEGER:
		return "INTEGER"

	default:
		return "UNDEFINED"
	}
}

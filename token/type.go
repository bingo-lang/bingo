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

	EQUALS
	SEMICOLON

	// Numbers.
	INTEGER

	// Words.
	LET
	IDENTIFIER

	INT
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

	case EQUALS:
		return "="
	case SEMICOLON:
		return ";"

	// Numbers.
	case INTEGER:
		return "INTEGER"

		//Words.
	case LET:
		return "let"
	case IDENTIFIER:
		return "identifier"

	case INT:
		return "int"

	default:
		return "UNDEFINED"
	}
}

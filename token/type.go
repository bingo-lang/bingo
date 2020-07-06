package token

type Type string

const (
	UNDEFINED = "undefined"
	EOF       = "EOF"

	// Symbols.
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	OR       = "||"
	AND      = "&&"
	GT       = ">"
	GTE      = ">="
	LT       = "<"
	LTE      = "<="
	EQUAL    = "=="
	BANG     = "!"

	ASSIGN    = "="
	SEMICOLON = ";"

	// Numbers.
	INTEGER = "integer"

	// Words.
	LET        = "let"
	IDENTIFIER = "identifier"

	INT  = "int"
	BOOL = "bool"

	BOOLEAN = "boolean" // It reprensets the true or false constants
)

func (t Type) String() string {
	return string(t)
}

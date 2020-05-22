package token

type Type uint

const (
	UNDEFINED Type = iota
	EOF
	// Numbers.
	INTEGER
)

func (t Type) String() string {
	switch t {
	//Numbers.
	case INTEGER:
		return "INTEGER"
	case EOF:
		return "EOF"
	default:
		return "UNDEFINED"
	}
}

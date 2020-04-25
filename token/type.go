package token

type Type uint

const (
	UNDEFINED Type = iota
	// Numbers.
	INTEGER
)

func (t Type) String() string {
	switch t {
	//Numbers.
	case INTEGER:
		return "INTEGER"
	default:
		return "UNDEFINED"
	}
}

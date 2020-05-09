package object

type Integer struct {
	Value int
}

func (i *Integer) Type() Type {
	return INTEGER
}

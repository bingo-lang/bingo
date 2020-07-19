package object

import (
	"fmt"
)

type Integer struct {
	Value int
}

func (i Integer) Type() Type {
	return INTEGER
}

func (i Integer) String() string {
	return fmt.Sprintf("%d", i.Value)
}

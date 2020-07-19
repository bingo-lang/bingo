package object

import (
	"fmt"
)

type Boolean struct {
	Value bool
}

func (b Boolean) Type() Type {
	return BOOLEAN
}

func (b Boolean) String() string {
	return fmt.Sprintf("%t", b.Value)
}

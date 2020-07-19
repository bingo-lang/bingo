package ast

import (
	"fmt"
	"github.com/bingo-lang/bingo/token"
)

type Expression interface {
	String() string
}

// Binary.
type ExpressionBinary struct {
	Operator        token.Token
	ExpressionLeft  Expression
	ExpressionRight Expression
}

func (eb ExpressionBinary) String() string {
	return fmt.Sprintf("(%s %s %s)", eb.ExpressionLeft, eb.Operator.Type, eb.ExpressionRight)
}

// Unary.
type ExpressionUnary struct {
	Operator   token.Token
	Expression Expression
	Invalid    bool
}

func (eu ExpressionUnary) String() string {
	return fmt.Sprintf("%s%s", eu.Operator.Type, eu.Expression)
}

// Integer.
type ExpressionInteger struct {
	Value string
}

func (ei ExpressionInteger) String() string {
	return ei.Value
}

// Boolean.
type ExpressionBoolean struct {
	Value string
}

func (eb ExpressionBoolean) String() string {
	return eb.Value
}

package ast

import (
	"fmt"
	"github.com/bingo-lang/bingo/token"
)

type Expression interface {
	String() string
	HasErrors() bool
}

// Binary.
type ExpressionBinary struct {
	Operator        token.Token
	ExpressionLeft  Expression
	ExpressionRight Expression
	Invalid         bool
}

func (eb ExpressionBinary) String() string {
	return fmt.Sprintf("(%s %s %s)", eb.ExpressionLeft, eb.Operator.Type, eb.ExpressionRight)
}

func (eb ExpressionBinary) HasErrors() bool {
	return eb.Invalid || eb.ExpressionLeft.HasErrors() || eb.ExpressionRight.HasErrors()
}

// Unary.
type ExpressionUnary struct {
	Operator   token.Token
	Expression Expression
	Invalid    bool
}

func (eu ExpressionUnary) HasErrors() bool {
	return eu.Invalid || eu.Expression.HasErrors()
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

func (ei ExpressionInteger) HasErrors() bool {
	return false
}

// Boolean.
type ExpressionBoolean struct {
	Value string
}

func (eb ExpressionBoolean) String() string {
	return eb.Value
}

func (eb ExpressionBoolean) HasErrors() bool {
	return false
}

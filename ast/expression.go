package ast

import (
	"github.com/bingo-lang/bingo/token"
)

type Expression interface {
}

type ExpressionBinary struct {
	Operator        token.Token
	ExpressionLeft  Expression
	ExpressionRight Expression
}

type ExpressionUnary struct {
	Operator   token.Token
	Expression Expression
}

type ExpressionInteger struct {
	Value string
}

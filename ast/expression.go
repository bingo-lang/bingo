package ast

import (
	"github.com/bingo-lang/bingo/token"
)

type Expression interface {
}

type ExpressionInfix struct {
	Operator        token.Token
	ExpressionLeft  Expression
	ExpressionRight Expression
}

type ExpressionPrefix struct {
	Operator   token.Token
	Expression Expression
}

type ExpressionInteger struct {
	Value string
}

package ast

import (
	"github.com/bingo-lang/bingo/token"
)

type Expression interface {
}

type ExpressionPrefix struct {
	Operator   token.Token
	Expression Expression
}

type ExpressionInteger struct {
	Value string
}

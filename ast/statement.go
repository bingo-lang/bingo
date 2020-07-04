package ast

import (
	"fmt"
	"github.com/bingo-lang/bingo/token"
)

type Statement interface {
	Node
}

type StatementExpression struct {
	Expression
}

type StatementError struct {
	InvalidToken token.Token
}

func (se *StatementError) String() string {
	return fmt.Sprintf("SyntaxError: Invalid token %s", se.InvalidToken.Value)
}

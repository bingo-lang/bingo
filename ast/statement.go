package ast

import (
	"fmt"
	"github.com/bingo-lang/bingo/token"
)

type Statement interface {
	String() string
}

type StatementExpression struct {
	Expression
}

func (se StatementExpression) String() string {
	return se.Expression.String()
}

type StatementLet struct {
	Identifier token.Token
	Expression
}

func (se StatementLet) String() string {
	return fmt.Sprintf("let %s = %s", se.Identifier.Value, se.Expression)
}

type StatementBlock struct {
	Statements []Statement
}

func (sb StatementBlock) String() string {
	// TODO(tugorez): Improve this section
	return ""
}

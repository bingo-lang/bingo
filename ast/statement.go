package ast

type Statement interface {
	String() string
}

type StatementExpression struct {
	Expression
}

func (se StatementExpression) String() string {
	return se.Expression.String()
}

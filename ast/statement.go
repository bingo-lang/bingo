package ast

type Statement interface {
	String() string
	HasErrors() bool
}

type StatementExpression struct {
	Expression
	Invalid bool
}

func (se StatementExpression) HasErrors() bool {
	return se.Invalid || se.Expression.HasErrors()
}

func (se StatementExpression) String() string {
	return se.Expression.String()
}

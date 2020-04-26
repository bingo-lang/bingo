package ast

type Statement interface {
	Node
}

type StatementExpression struct {
	Expression
}

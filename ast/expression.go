package ast

type Expression interface {
	Node
}

type ExpressionInfix interface {
}

type ExpressionInfixInteger struct {
	Value string
}

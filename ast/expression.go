package ast

type Expression interface {
	Node
}

type ExpressionPrefix interface {
}

type ExpressionPrefixInteger struct {
	Value string
}

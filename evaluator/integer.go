package evaluator

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/object"
	"strconv"
)

func evalExpressionPrefixInteger(node ast.ExpressionPrefixInteger) *object.Integer {
	value, _ := strconv.Atoi(node.Value)
	return &object.Integer{Value: value}
}

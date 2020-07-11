package evaluator

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/object"
	"strconv"
)

func evalExpressionInteger(node ast.ExpressionInteger) *object.Integer {
	value, _ := strconv.Atoi(node.Value)
	return &object.Integer{Value: value}
}

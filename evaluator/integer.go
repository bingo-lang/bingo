package evaluator

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/object"
	"strconv"
)

func evalExpressionInteger(node ast.ExpressionInteger) (integer object.Integer, err error) {
	value, _ := strconv.Atoi(node.Value)
	integer = object.Integer{Value: value}
	return
}

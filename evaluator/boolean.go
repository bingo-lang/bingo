package evaluator

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/object"
)

func evalExpressionBoolean(node ast.ExpressionBoolean) (boolean object.Boolean, err error) {
	value := node.Value == "true"
	boolean = object.Boolean{Value: value}
	return
}

package evaluator

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/object"
)

func evalExpressionBoolean(node *ast.ExpressionBoolean) *object.Boolean {
	value := node.Value == "true"
	return &object.Boolean{Value: value}
}

package evaluator

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case ast.ExpressionInfixInteger:
		return evalExpressionInfixInteger(node)
	default:
		return nil
	}
}

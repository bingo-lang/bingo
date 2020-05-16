package evaluator

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.StatementExpression:
		return evalExpression(node.Expression)
	default:
		return nil
	}
}

func evalExpression(node ast.Expression) object.Object {
	switch node := node.(type) {
	case ast.ExpressionPrefix:
		return evalExpressionPrefix(node)
	default:
		return nil
	}
}

func evalExpressionPrefix(node ast.ExpressionPrefix) object.Object {
	switch node := node.(type) {
	case *ast.ExpressionPrefixInteger:
		return evalExpressionPrefixInteger(node)
	default:
		return nil
	}
}

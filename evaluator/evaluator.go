package evaluator

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/object"
	"github.com/bingo-lang/bingo/token"
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
	case *ast.ExpressionInteger:
		return evalExpressionInteger(node)
	case *ast.ExpressionUnary:
		return evalExpressionUnary(node)
	default:
		return nil
	}
}

func evalExpressionUnary(node *ast.ExpressionUnary) object.Object {
	switch node.Operator.Type {
	case token.MINUS:
		return evalMinus(node.Expression)
	default:
		return nil
	}
}

func evalMinus(node ast.Expression) object.Object {
	switch ob := evalExpression(node).(type) {
	case *object.Integer:
		ob.Value = -ob.Value
		return ob
	default:
		return nil
	}
}

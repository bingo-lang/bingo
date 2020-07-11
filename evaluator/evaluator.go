package evaluator

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/object"
	"github.com/bingo-lang/bingo/token"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case ast.StatementExpression:
		return evalExpression(node.Expression)
	default:
		return nil
	}
}

func evalExpression(node ast.Expression) object.Object {
	switch node := node.(type) {
	case ast.ExpressionBoolean:
		return evalExpressionBoolean(node)
	case ast.ExpressionInteger:
		return evalExpressionInteger(node)
	case ast.ExpressionUnary:
		return evalExpressionUnary(node)
	case ast.ExpressionBinary:
		return evalExpressionBinary(node)
	default:
		// This is an error
		return nil
	}
}

func evalExpressionBinary(node ast.ExpressionBinary) object.Object {
	left, _ := evalExpression(node.ExpressionLeft).(*object.Integer)
	right, _ := evalExpression(node.ExpressionRight).(*object.Integer)
	switch node.Operator.Type {
	case token.PLUS:
		return &object.Integer{Value: left.Value + right.Value}
	case token.MINUS:
		return &object.Integer{Value: left.Value - right.Value}
	case token.ASTERISK:
		return &object.Integer{Value: left.Value * right.Value}
	case token.SLASH:
		return &object.Integer{Value: left.Value / right.Value}
	default:
		// This is an error
		return nil
	}
}

func evalExpressionUnary(node ast.ExpressionUnary) object.Object {
	switch node.Operator.Type {
	case token.PLUS:
		return evalExpression(node.Expression)
	case token.MINUS:
		return evalMinus(node.Expression)
	case token.BANG:
		return evalBang(node.Expression)
	default:
		// This is an error
		return nil
	}
}

func evalMinus(node ast.Expression) object.Object {
	switch ob := evalExpression(node).(type) {
	case *object.Integer:
		ob.Value = -ob.Value
		return ob
	default:
		// This is an error
		return nil
	}
}

func evalBang(node ast.Expression) object.Object {
	switch ob := evalExpression(node).(type) {
	case *object.Boolean:
		ob.Value = !ob.Value
		return ob
	default:
		// This is an error
		return nil
	}
}

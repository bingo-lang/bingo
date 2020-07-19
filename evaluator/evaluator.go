package evaluator

import (
	"fmt"
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/object"
	"github.com/bingo-lang/bingo/token"
)

func Eval(statement ast.Statement) (object.Object, error) {
	switch statement := statement.(type) {
	case ast.StatementExpression:
		return evalExpression(statement.Expression)
	default:
		return nil, fmt.Errorf("Unsupported statement %T", statement)
	}
}

func evalExpression(expression ast.Expression) (object.Object, error) {
	switch expression := expression.(type) {
	case ast.ExpressionBinary:
		return evalExpressionBinary(expression)
	case ast.ExpressionUnary:
		return evalExpressionUnary(expression)
	case ast.ExpressionInteger:
		return evalExpressionInteger(expression)
	case ast.ExpressionBoolean:
		return evalExpressionBoolean(expression)
	default:
		return nil, fmt.Errorf("Unsupported expression %q of type %T", expression, expression)
	}
}

func evalExpressionBinary(expressionBinary ast.ExpressionBinary) (object.Object, error) {
	left, err := evalExpression(expressionBinary.ExpressionLeft)
	if err != nil {
		return nil, err
	}
	right, err := evalExpression(expressionBinary.ExpressionRight)
	if err != nil {
		return nil, err
	}
	switch expressionBinary.Operator.Type {
	case token.PLUS:
		return evalBinaryPlus(left, right)
	case token.MINUS:
		return evalBinaryMinus(left, right)
	case token.ASTERISK:
		return evalBinaryAsterisk(left, right)
	case token.SLASH:
		return evalBinarySlash(left, right)
	case token.GT:
		return evalBinaryGT(left, right)
	case token.GTE:
		return evalBinaryGTE(left, right)
	case token.LT:
		return evalBinaryLT(left, right)
	case token.LTE:
		return evalBinaryLTE(left, right)
	case token.EQUAL:
		return evalBinaryEqual(left, right)
	case token.AND:
		return evalBinaryAnd(left, right)
	case token.OR:
		return evalBinaryOr(left, right)
	default:
		return nil, fmt.Errorf("Unsupported binary operator %q", expressionBinary.Operator.Value)
	}
}

func evalExpressionUnary(expressionUnary ast.ExpressionUnary) (object.Object, error) {
	switch expressionUnary.Operator.Type {
	case token.PLUS:
		return evalExpression(expressionUnary.Expression)
	case token.MINUS:
		return evalUnaryMinus(expressionUnary.Expression)
	case token.BANG:
		return evalUnaryBang(expressionUnary.Expression)
	default:
		return nil, fmt.Errorf("Unsupported binary operator %q", expressionUnary.Operator.Value)
	}
}

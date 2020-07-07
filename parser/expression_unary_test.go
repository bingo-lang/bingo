package parser

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/token"
	"reflect"
	"strings"
	"testing"
)

func TestExpressionUnary(t *testing.T) {
	testCases := []struct {
		source     string
		operator   token.Type
		expression ast.Expression
	}{
		{"-143", token.MINUS, &ast.ExpressionInteger{Value: "143"}},
		{"+256", token.PLUS, &ast.ExpressionInteger{Value: "256"}},
		{"!true", token.BANG, &ast.ExpressionBoolean{Value: "true"}},
	}
	for _, testCase := range testCases {
		testExpressionUnary(t, testCase.source, testCase.operator, testCase.expression)
	}
}

func TestExpressionValue(t *testing.T) {
	testCases := []struct {
		source     string
		expression ast.Expression
	}{
		{"1", &ast.ExpressionInteger{Value: "1"}},
		{"10", &ast.ExpressionInteger{Value: "10"}},
		{"true", &ast.ExpressionBoolean{Value: "true"}},
		{"false", &ast.ExpressionBoolean{Value: "false"}},
	}
	for _, testCase := range testCases {
		source := strings.NewReader(testCase.source)
		parser := New(source)
		statement := parser.parseStatement()
		statementExpression, ok := (statement).(*ast.StatementExpression)
		if !ok {
			t.Fatalf("Expecting statement to be %s, got %T", "StatementExpression", statement)
		}
		expression := statementExpression.Expression
		if reflect.TypeOf(expression) != reflect.TypeOf(testCase.expression) {
			t.Fatalf("Expecting expression %s to be of type %T, got %T", testCase.source, testCase.expression, expression)

		}
		if expression.String() != testCase.expression.String() {
			t.Fatalf("Expecting expression %s value to be %s, got %s", testCase.source, testCase.expression, expression)
		}
	}
}

func testExpressionUnary(t *testing.T, source string, operator token.Type, expectedExpression ast.Expression) {
	reader := strings.NewReader(source)
	parser := New(reader)
	statement := parser.parseStatement()
	statementExpression, ok := (statement).(*ast.StatementExpression)
	if !ok {
		t.Fatalf("Expecting statement to be %s, got %T", "StatementExpression", statement)
	}
	expressionPrefix, ok := statementExpression.Expression.(*ast.ExpressionUnary)
	if !ok {
		t.Fatalf("Expecting expression to be %s, got %T", "ExpressionUnary", expressionPrefix)
	}
	if expressionPrefix.Operator.Type != operator {
		t.Fatalf("Expecting operator to be of type %q, got %q", operator, expressionPrefix.Operator.Type)
	}
	expression := expressionPrefix.Expression
	if reflect.TypeOf(expression) != reflect.TypeOf(expectedExpression) {
		t.Fatalf("Expecting expression %s to be of type %T, got %T", source, expectedExpression, expression)
	}
	if expression.String() != expectedExpression.String() {
		t.Fatalf("Expecting expression %s value to be %s, got %s", source, expectedExpression, expression)
	}

}

package parser

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/token"
	"strings"
	"testing"
)

func TestExpressionUnary(t *testing.T) {
	testCases := []struct {
		source   string
		operator token.Type
		value    string
	}{
		{"-143", token.MINUS, "143"},
		{"+256", token.PLUS, "256"},
		{"-10", token.MINUS, "10"},
	}
	for _, testCase := range testCases {
		testExpressionUnary(t, testCase.source, testCase.operator, testCase.value)
	}
}

func TestExpressionInteger(t *testing.T) {
	source := strings.NewReader(`237`)
	parser := New(source)
	statement := parser.parseStatement()
	statementExpression, ok := (statement).(*ast.StatementExpression)
	if !ok {
		t.Fatalf("Expecting statement to be %s, got %T", "StatementExpression", statement)
	}
	expressionInteger, ok := statementExpression.Expression.(*ast.ExpressionInteger)
	if !ok {
		t.Fatalf("Expecting expressionInteger to be %s, got %T", "ExpressionInteger", expressionInteger)
	}
	if expressionInteger.Value != "237" {
		t.Fatalf("Expecting expressionInteger prefix integer value to be %q, got %q", "237", "237")
	}
}

func testExpressionUnary(t *testing.T, source string, operator token.Type, value string) {
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
	expressionInteger, ok := expressionPrefix.Expression.(*ast.ExpressionInteger)
	if !ok {
		t.Fatalf("Expecting expression to be %s, got %T", "ExpressionInteger", expressionInteger)
	}
	if expressionInteger.Value != value {
		t.Fatalf("Expecting expression prefix integer value to be %q, got %q", value, expressionInteger.Value)
	}
}

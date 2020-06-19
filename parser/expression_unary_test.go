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
	gotten := parser.Parse()
	if len(gotten.Statements) != 1 {
		t.Fatalf("Expecting 1 statement, got %d", len(gotten.Statements))
	}
	statement, ok := (gotten.Statements[0]).(*ast.StatementExpression)
	if !ok {
		t.Fatalf("Expecting statement to be %q, got %q", "StatementExpression", reflect.TypeOf(statement))
	}
	expressionInteger, ok := statement.Expression.(*ast.ExpressionInteger)
	if !ok {
		t.Fatalf("Expecting expressionInteger to be %q, got %q", "ExpressionInteger", reflect.TypeOf(expressionInteger))
	}
	if expressionInteger.Value != "237" {
		t.Fatalf("Expecting expressionInteger prefix integer value to be %q, got %q", "237", "237")
	}
}

func testExpressionUnary(t *testing.T, source string, operator token.Type, value string) {
	reader := strings.NewReader(source)
	parser := New(reader)
	gotten := parser.Parse()
	if len(gotten.Statements) != 1 {
		t.Fatalf("Expecting 1 statement, got %d", len(gotten.Statements))
	}
	statement, ok := (gotten.Statements[0]).(*ast.StatementExpression)
	if !ok {
		t.Fatalf("Expecting statement to be %q, got %q", "StatementExpression", reflect.TypeOf(statement))
	}
	expressionPrefix, ok := statement.Expression.(*ast.ExpressionUnary)
	if !ok {
		t.Fatalf("Expecting expression to be %q, got %q", "ExpressionUnary", reflect.TypeOf(expressionPrefix))
	}
	if expressionPrefix.Operator.Type != operator {
		t.Fatalf("Expecting operator to be of type %q, got %q", operator, expressionPrefix.Operator.Type)
	}
	expressionInteger, ok := expressionPrefix.Expression.(*ast.ExpressionInteger)
	if !ok {
		t.Fatalf("Expecting expression to be %q, got %q", "ExpressionInteger", reflect.TypeOf(expressionInteger))
	}
	if expressionInteger.Value != value {
		t.Fatalf("Expecting expression prefix integer value to be %q, got %q", value, expressionInteger.Value)
	}
}

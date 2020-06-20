package parser

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/token"
	"reflect"
	"strings"
	"testing"
)

func TestOperatorPrecedence(t *testing.T) {
	testCases := []struct{ source, expected string }{
		{"1", "1"},
		{"+2", "+2"},
		{"-3", "-3"},
		{"1 + 2", "(1 + 2)"},
		{"1 + 2 * 3", "(1 + (2 * 3))"},
		{"1 / 2 * 3", "((1 / 2) * 3)"},
		{"1 / 2 * 3 + 5", "(((1 / 2) * 3) + 5)"},
	}
	for _, testCase := range testCases {
		testOperatorPrecedence(t, testCase.source, testCase.expected)
	}
}

func testOperatorPrecedence(t *testing.T, source string, expected string) {
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
	expressionStr := statement.Expression.String()
	if expressionStr != expected {
		t.Fatalf("Expecting expression to be %q, got %q", expected, expressionStr)
	}
}

func TestExpressionBinary(t *testing.T) {
	testCases := []struct {
		source                string
		operator              token.Type
		leftValue, rightValue string
	}{
		{"143 - 1", token.MINUS, "143", "1"},
		{"120 + 3", token.PLUS, "120", "2"},
		{"1 - 0", token.MINUS, "1", "0"},
		{"1 * 0", token.ASTERISK, "1", "0"},
		{"0 / 1", token.SLASH, "0", "1"},
	}
	for _, testCase := range testCases {
		testExpressionBinary(t, testCase.source, testCase.operator, testCase.leftValue, testCase.rightValue)
	}
}

func testExpressionBinary(t *testing.T, source string, operator token.Type, valueLeft, valueRight string) {
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
	expressionPrefix, ok := statement.Expression.(*ast.ExpressionBinary)
	if !ok {
		t.Fatalf("Expecting expression to be %q, got %q", "ExpressionBinary", reflect.TypeOf(expressionPrefix))
	}
	if expressionPrefix.Operator.Type != operator {
		t.Fatalf("Expecting operator to be of type %q, got %q", operator, expressionPrefix.Operator.Type)
	}
	// Left expression.
	expressionIntegerLeft, ok := expressionPrefix.ExpressionLeft.(*ast.ExpressionInteger)
	if !ok {
		t.Fatalf("Expecting expression left to be %q, got %q", "ExpressionInteger", reflect.TypeOf(expressionIntegerLeft))
	}
	if expressionIntegerLeft.Value != valueLeft {
		t.Fatalf("Expecting expression left integer value to be %q, got %q", valueLeft, expressionIntegerLeft.Value)
	}
	// Right expression.
	expressionIntegerRight, ok := expressionPrefix.ExpressionLeft.(*ast.ExpressionInteger)
	if !ok {
		t.Fatalf("Expecting expression right to be %q, got %q", "ExpressionInteger", reflect.TypeOf(expressionIntegerRight))
	}
	if expressionIntegerRight.Value != valueLeft {
		t.Fatalf("Expecting expression right integer value to be %q, got %q", valueRight, expressionIntegerLeft.Value)
	}
}

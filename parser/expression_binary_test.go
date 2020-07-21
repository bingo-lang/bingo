package parser

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/token"
	"strings"
	"testing"
)

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
		{"1 > 0", token.GT, "1", "0"},
		{"1 >= 0", token.GTE, "1", "0"},
		{"1 < 0", token.LT, "1", "0"},
		{"1 <= 0", token.LTE, "1", "0"},
		{"1 == 0", token.EQUAL, "1", "0"},
	}
	for _, testCase := range testCases {
		testExpressionBinary(t, testCase.source, testCase.operator, testCase.leftValue, testCase.rightValue)
	}
}

func TestOperatorPrecedence(t *testing.T) {
	testCases := []struct{ source, expected string }{
		{"1", "1"},
		{"+2", "+2"},
		{"-3", "-3"},
		{"1 + 2", "(1 + 2)"},
		{"1 + 2 * 3", "(1 + (2 * 3))"},
		{"1 / 2 * 3", "((1 / 2) * 3)"},
		{"1 / 2 * 3 + 5", "(((1 / 2) * 3) + 5)"},
		{"1 > 2", "(1 > 2)"},
		{"1 > 2 * 3", "(1 > (2 * 3))"},
		{"1 > 2 + 3", "(1 > (2 + 3))"},
		{"1 == 2 / 2", "(1 == (2 / 2))"},
		{"1 < 2 || 1 > 2", "((1 < 2) || (1 > 2))"},
		{"(1 + 2) * 3", "((1 + 2) * 3)"},
	}
	for _, testCase := range testCases {
		testOperatorPrecedence(t, testCase.source, testCase.expected)
	}
}

func testExpressionBinary(t *testing.T, source string, operator token.Type, valueLeft, valueRight string) {
	reader := strings.NewReader(source)
	parser := New(reader)
	statement, err := parser.parseStatement()
	if err != nil {
		t.Fatal(err)
	}
	statementExpression, ok := (statement).(ast.StatementExpression)
	if !ok {
		t.Fatalf("Expecting %q statement to be %s, got %T", source, "StatementExpression", statement)
	}
	expressionPrefix, ok := statementExpression.Expression.(ast.ExpressionBinary)
	if !ok {
		t.Fatalf("Expecting %q expression to be %s, got %T", source, "ExpressionBinary", expressionPrefix)
	}
	if expressionPrefix.Operator.Type != operator {
		t.Fatalf("Expecting operator to be of type %q, got %q", operator, expressionPrefix.Operator.Type)
	}
	// Left expression.
	expressionIntegerLeft, ok := expressionPrefix.ExpressionLeft.(ast.ExpressionInteger)
	if !ok {
		t.Fatalf("Expecting expression left to be %s, got %T", "ExpressionInteger", expressionIntegerLeft)
	}
	if expressionIntegerLeft.Value != valueLeft {
		t.Fatalf("Expecting expression's left value to be %q, got %q", valueLeft, expressionIntegerLeft.Value)
	}
	// Right expression.
	expressionIntegerRight, ok := expressionPrefix.ExpressionLeft.(ast.ExpressionInteger)
	if !ok {
		t.Fatalf("Expecting expression right to be %s, got %T", "ExpressionInteger", expressionIntegerRight)
	}
	if expressionIntegerRight.Value != valueLeft {
		t.Fatalf("Expecting expression right integer value to be %q, got %q", valueRight, expressionIntegerLeft.Value)
	}
}

func testOperatorPrecedence(t *testing.T, source string, expected string) {
	reader := strings.NewReader(source)
	parser := New(reader)
	statement, err := parser.parseStatement()
	if err != nil {
		t.Fatal(err)
	}
	statementExpression, ok := (statement).(ast.StatementExpression)
	if !ok {
		t.Fatalf("Expecting statement to be %s, got %T", "StatementExpression", statement)
	}
	expressionStr := statementExpression.Expression.String()
	if expressionStr != expected {
		t.Fatalf("Expecting expression to be %q, got %q", expected, expressionStr)
	}
}

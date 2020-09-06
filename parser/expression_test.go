package parser

import (
	"strings"
	"testing"
)

// TODO(tugorez): Validate it throws exception when invalid.
// TODO(tugorez): Validate it calls the boolean expression.
// TODO(tugorez): Validate it calls the integer expression.
// TODO(tugorez): Validate it calls the grouped expression.
// TODO(tugorez): Validate it calls the unary expression.
// TODO(tugorez): Validate it calls the binary operator.
func TestExpressionOperatorPrecedence(t *testing.T) {
	testCases := []struct{ source, expected string }{
		{"1", "1"},
		{"+2", "+2"},
		{"-3", "-3"},
		{"1 + 2", "(1 + 2)"},
		{"1 + 2 * 3", "(1 + (2 * 3))"},
		{"1 / 2 * 3", "(1 / (2 * 3))"},
		{"1 / 2 * 3 + 5", "((1 / (2 * 3)) + 5)"},
		{"1 > 2", "(1 > 2)"},
		{"1 > 2 * 3", "(1 > (2 * 3))"},
		{"1 > 2 + 3", "(1 > (2 + 3))"},
		{"1 == 2 / 2", "(1 == (2 / 2))"},
		{"1 < 2 || 1 > 2", "((1 < 2) || (1 > 2))"},
		{"(1 + 2) * 3", "((1 + 2) * 3)"},
		{"(-1) == (-1)", "(-1 == -1)"},
		{"(-1) == -1", "(-1 == -1)"},
		{"-1 == -1", "(-1 == -1)"},
		{"1 - 1 == 1", "((1 - 1) == 1)"},
	}
	for _, testCase := range testCases {
		source := strings.NewReader(testCase.source)
		parser := New(source)
		expression, err := parser.parseExpression(LOWEST)
		if err != nil {
			t.Fatal(err)
		}
		if expression.String() != testCase.expected {
			t.Fatalf("Expecting expression %s precedence to be %q got %q", testCase.source, testCase.expected, expression.String())
		}
	}
}

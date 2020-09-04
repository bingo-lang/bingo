package parser

import (
	"strings"
	"testing"
)

func TestExpressionInteger(t *testing.T) {
	testCases := []string{
		"1",
		"10",
	}
	for _, testCase := range testCases {
		source := strings.NewReader(testCase)
		parser := New(source)
		expressionInteger, err := parser.parseExpressionInteger()
		if err != nil {
			t.Fatal(err)
		}
		if expressionInteger.Value != testCase {
			t.Fatalf("Expected ExpressionInteger value to be %s, got %s", testCase, expressionInteger.Value)
		}
	}
}

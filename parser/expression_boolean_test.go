package parser

import (
	"strings"
	"testing"
)

func TestExpressionBoolean(t *testing.T) {
	testCases := []string{
		"true",
		"false",
	}
	for _, testCase := range testCases {
		source := strings.NewReader(testCase)
		parser := New(source)
		expressionBoolean, err := parser.parseExpressionBoolean()
		if err != nil {
			t.Fatal(err)
		}
		if expressionBoolean.Value != testCase {
			t.Fatalf("Expected ExpressionBoolean value to be %s, got %s", testCase, expressionBoolean.Value)
		}
	}
}

func TestExpressionBooleanErrors(t *testing.T) {
	testCases := []string{
		"@",
		"1",
		"pi",
	}
	for _, testCase := range testCases {
		source := strings.NewReader(testCase)
		parser := New(source)
		expressionBoolean, err := parser.parseExpressionBoolean()
		if err == nil {
			t.Fatalf("Expecting invalid expression %q to throw an error. Got a boolean with value %q instead", testCase, expressionBoolean.Value)
		}
	}
}

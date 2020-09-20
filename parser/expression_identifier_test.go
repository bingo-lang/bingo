package parser

import (
	"strings"
	"testing"
)

func TestExpressionIdentifier(t *testing.T) {
	testCases := []string{
		"x",
		"y",
	}
	for _, testCase := range testCases {
		source := strings.NewReader(testCase)
		parser := New(source)
		expressionIdentifier, err := parser.parseExpressionIdentifier()
		if err != nil {
			t.Fatal(err)
		}
		if expressionIdentifier.Value != testCase {
			t.Fatalf("Expected ExpressionIdentifier value to be %s, got %s", testCase, expressionIdentifier.Value)
		}
	}
}

func TestExpressionIdentifierError(t *testing.T) {
	testCases := []string{
		"@",
		"1",
	}
	for _, testCase := range testCases {
		source := strings.NewReader(testCase)
		parser := New(source)
		expressionIdentifier, err := parser.parseExpressionIdentifier()
		if err == nil {
			t.Fatalf("Expecting invalid expression %q to throw an error. Got an identifier with value %q instead", testCase, expressionIdentifier.Value)
		}
	}
}

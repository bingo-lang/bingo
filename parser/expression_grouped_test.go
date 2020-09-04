package parser

import (
	"strings"
	"testing"
)

func TestExpressionGrouped(t *testing.T) {
	// TODO(tugorez): Validate it throws an error with invalid parentheses.
	testCases := []string{
		"(1)",
		"(true)",
	}
	for _, testCase := range testCases {
		source := strings.NewReader(testCase)
		parser := New(source)
		_, err := parser.parseExpressionGrouped()
		if err != nil {
			t.Fatal(err)
		}
	}
}

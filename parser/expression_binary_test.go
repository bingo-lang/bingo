package parser

import (
	"github.com/bingo-lang/bingo/token"
	"strings"
	"testing"
)

func TestExpressionBinary(t *testing.T) {
	testCases := []struct {
		source   string
		operator token.Type
	}{
		{"+ 1", token.PLUS},
		{"- 1", token.MINUS},
		{"* 1", token.ASTERISK},
		{"/ 1", token.SLASH},
		{"> 1", token.GT},
		{">= 1", token.GTE},
		{"< 0", token.LT},
		{"<= 0", token.LTE},
		{"== 0", token.EQUAL},
	}
	for _, testCase := range testCases {
		source := strings.NewReader(testCase.source)
		parser := New(source)
		expressionBinary, err := parser.parseExpressionBinary(nil, LOWEST)
		if err != nil {
			t.Fatal(err)
		}
		if expressionBinary.Operator.Type != testCase.operator {
			t.Fatalf("Expecting binary expression %s operator to be %q got %q", testCase.source, testCase.operator, expressionBinary.Operator.Type)
		}
	}
}

func TestExpressionBinaryError(t *testing.T) {
	testCases := []string{
		"- /",
		"@ /",
		"@ 1",
	}
	for _, testCase := range testCases {
		source := strings.NewReader(testCase)
		parser := New(source)
		expressionBinary, err := parser.parseExpressionBinary(nil, LOWEST)
		if err == nil {
			t.Fatalf("Expecting invalid expression %q to throw an error. Got %s instead", testCase, expressionBinary)
		}
	}
}

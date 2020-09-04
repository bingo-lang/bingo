package parser

import (
	"github.com/bingo-lang/bingo/token"
	"strings"
	"testing"
)

func TestExpressionUnary(t *testing.T) {
	testCases := []struct {
		source   string
		operator token.Type
	}{
		{"-10", token.MINUS},
		{"+1", token.PLUS},
		{"!false", token.BANG},
	}
	for _, testCase := range testCases {
		source := strings.NewReader(testCase.source)
		parser := New(source)
		expressionUnary, err := parser.parseExpressionUnary()
		if err != nil {
			t.Fatal(err)
		}
		if expressionUnary.Operator.Type != testCase.operator {
			t.Fatalf("Expecting unary expression %s operator to be %q got %q", testCase.source, testCase.operator, expressionUnary.Operator.Type)
		}
	}
}

func TestExpressionUnaryError(t *testing.T) {
	testCases := []string{
		"@1",
	}
	for _, testCase := range testCases {
		source := strings.NewReader(testCase)
		parser := New(source)
		expressionUnary, err := parser.parseExpressionUnary()
		if err == nil {
			t.Fatalf("Expecting invalid expression %q to throw an error. Got %s instead", testCase, expressionUnary)
		}
	}
}

package parser

import (
	"strings"
	"testing"
)

func TestParseStatementLet(t *testing.T) {
	testCases := []struct{ source, identifier string }{
		{"let x = 1;", "x"},
		{"let x = 2; ", "x"},
	}
	for _, testCase := range testCases {
		reader := strings.NewReader(testCase.source)
		parser := New(reader)
		statement, err := parser.parseStatementLet()
		if err != nil {
			t.Fatalf("Error on %q, %s", testCase.source, err)
		}
		if statement.Identifier.Value != testCase.identifier {
			t.Fatalf("Expecting identifier %q, got %q instead", testCase.identifier, statement.Identifier.Value)
		}
	}
}

func TestParseStatementLetError(t *testing.T) {
	testCases := []string{
		"let",
		"let = 1",
		"let x 1",
	}
	for _, testCase := range testCases {
		reader := strings.NewReader(testCase)
		parser := New(reader)
		statement, err := parser.parseStatementLet()
		if err == nil {
			t.Fatalf("Expecting invalid expression %q to throw an error. Got %s instead", testCase, statement)
		}
	}
}

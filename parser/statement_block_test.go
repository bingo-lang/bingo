package parser

import (
	"strings"
	"testing"
)

func TestParseStatementBlock(t *testing.T) {
	testCases := []struct {
		source             string
		numberOfStatements int
	}{
		{"{1;}", 1},
		{"{1;2;}", 2},
		{"{let x = 10;1;}", 2},
	}
	for _, testCase := range testCases {
		reader := strings.NewReader(testCase.source)
		parser := New(reader)
		statement, err := parser.parseStatementBlock()
		if err != nil {
			t.Fatal(err)
		}
		if len(statement.Statements) != testCase.numberOfStatements {
			t.Fatalf("Expecting block statement %s to have %d statements, got %d instead", testCase.source, testCase.numberOfStatements, len(statement.Statements))
		}
	}
}

func TestParseStatementBlockError(t *testing.T) {
	testCases := []string{
		"let",
		"let = 1",
		"let x 1",
		"let x = @;",
	}
	for _, testCase := range testCases {
		reader := strings.NewReader(testCase)
		parser := New(reader)
		statement, err := parser.parseStatementBlock()
		if err == nil {
			t.Fatalf("Expecting invalid expression %q to throw an error. Got %s instead", testCase, statement)
		}
	}
}

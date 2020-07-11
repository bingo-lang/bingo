package parser

import (
	"strings"
	"testing"
)

func TestParseStatementError(t *testing.T) {
	testCases := []string{"-$", "$10 + 5", "10 $ 20"}
	for _, testCase := range testCases {
		reader := strings.NewReader(testCase)
		parser := New(reader)
		statement := parser.parseStatement()
		if !statement.HasErrors() {
			t.Fatalf("Expecting %q to has errors got %s", testCase, statement)
		}
		if !parser.IsEOF() {
			t.Fatal("Expecting parser to consume all the tokens until statement separator")
		}
	}
}

package parser

import (
	"github.com/bingo-lang/bingo/ast"
	"strings"
	"testing"
)

func TestParseStatementError(t *testing.T) {
	testCases := []string{"-$", "$10 + 5", "10 $ 20"}
	for _, testCase := range testCases {
		reader := strings.NewReader(testCase)
		parser := New(reader)
		statement := parser.parseStatement()
		_, ok := statement.(*ast.StatementError)
		if !ok {
			t.Fatalf("Expecting %q to be an StatementError got %T", testCase, statement)
		}
		if !parser.IsEOF() {
			t.Fatal("Expecting parser to consume all the tokens until statement separator")
		}
	}
}

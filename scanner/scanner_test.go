package scanner

import (
	"github.com/bingo-lang/bingo/token"
	"strings"
	"testing"
)

func TestScanner(t *testing.T) {
	source := strings.NewReader(`
	1
	2
	3
	4
	5
	`)
	expectedTokens := []token.Token{
		token.New(token.INTEGER, "1"),
		token.New(token.INTEGER, "2"),
		token.New(token.INTEGER, "3"),
		token.New(token.INTEGER, "4"),
		token.New(token.INTEGER, "5"),
	}
	scanner, err := New(source)
	if err != nil {
		t.Fatal(err)
	}
	for _, expectedToken := range expectedTokens {
		nextToken := scanner.NextToken()
		if nextToken.Type != expectedToken.Type {
			t.Fatalf("Invalid token type. Expecting %q, got %q", expectedToken.Type, nextToken.Type)
		}
		if nextToken.Value != expectedToken.Value {
			t.Fatalf("Invalid token value. Expecting %q, got %q", expectedToken.Value, nextToken.Value)
		}
	}
	if !scanner.EOF() {
		t.Fatalf("Expecting no more tokens but got %q", scanner.NextToken())
	}
}

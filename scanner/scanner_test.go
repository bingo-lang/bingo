package scanner

import (
	"github.com/bingo-lang/bingo/token"
	"strings"
	"testing"
)

func TestScanner(t *testing.T) {
	source := strings.NewReader(`
	let x = 10;
	`)
	expectedTokens := []token.Token{
		token.New(token.LET, "let"),
		token.New(token.IDENTIFIER, "x"),
		token.New(token.ASSIGNMENT, "="),
		token.New(token.INTEGER, "10"),
		token.New(token.SEMICOLON, ";"),
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

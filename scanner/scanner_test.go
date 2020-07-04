package scanner

import (
	"github.com/bingo-lang/bingo/token"
	"strings"
	"testing"
)

func TestScanner(t *testing.T) {
	source := strings.NewReader(`10 + 1 - 1 * 1 / 1;`)
	expectedTokens := []token.Token{
		token.New(token.INTEGER, "10"),
		token.New(token.PLUS, "+"),
		token.New(token.INTEGER, "1"),
		token.New(token.MINUS, "-"),
		token.New(token.INTEGER, "1"),
		token.New(token.ASTERISK, "*"),
		token.New(token.INTEGER, "1"),
		token.New(token.SLASH, "/"),
		token.New(token.INTEGER, "1"),
		token.New(token.SEMICOLON, ";"),
		token.New(token.EOF, ""),
	}
	scanner := New(source)
	for _, expectedToken := range expectedTokens {
		nextToken := scanner.ScanToken()
		if nextToken.Type != expectedToken.Type {
			t.Fatalf("Invalid token type. Expecting %q, got %q", expectedToken.Type, nextToken.Type)
		}
		if nextToken.Value != expectedToken.Value {
			t.Fatalf("Invalid token value. Expecting %q, got %q", expectedToken.Value, nextToken.Value)
		}
	}
}

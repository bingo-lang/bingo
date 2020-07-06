package scanner

import (
	"github.com/bingo-lang/bingo/token"
	"strings"
	"testing"
)

func TestScanner(t *testing.T) {
	source := strings.NewReader(`
		let x int = 10 + 1 - 1 * 1 / 1;
		let y bool = true || false && true;
		let z bool = 1 > 2 || 1 >= 2 && 1 < 2 && 1 <= 2 && 1 == 2;
		let a bool = !true;
	`)
	expectedTokens := []token.Token{
		// First statement.
		token.New(token.LET, "let"),
		token.New(token.IDENTIFIER, "x"),
		token.New(token.INT, "int"),
		token.New(token.ASSIGN, "="),
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

		// Second statement.
		token.New(token.LET, "let"),
		token.New(token.IDENTIFIER, "y"),
		token.New(token.BOOL, "bool"),
		token.New(token.ASSIGN, "="),
		token.New(token.BOOLEAN, "true"),
		token.New(token.OR, "||"),
		token.New(token.BOOLEAN, "false"),
		token.New(token.AND, "&&"),
		token.New(token.BOOLEAN, "true"),
		token.New(token.SEMICOLON, ";"),

		// Third statement
		token.New(token.LET, "let"),
		token.New(token.IDENTIFIER, "z"),
		token.New(token.BOOL, "bool"),
		token.New(token.ASSIGN, "="),
		token.New(token.INTEGER, "1"),
		token.New(token.GT, ">"),
		token.New(token.INTEGER, "2"),
		token.New(token.OR, "||"),
		token.New(token.INTEGER, "1"),
		token.New(token.GTE, ">="),
		token.New(token.INTEGER, "2"),
		token.New(token.AND, "&&"),
		token.New(token.INTEGER, "1"),
		token.New(token.LT, "<"),
		token.New(token.INTEGER, "2"),
		token.New(token.AND, "&&"),
		token.New(token.INTEGER, "1"),
		token.New(token.LTE, "<="),
		token.New(token.INTEGER, "2"),
		token.New(token.AND, "&&"),
		token.New(token.INTEGER, "1"),
		token.New(token.EQUAL, "=="),
		token.New(token.INTEGER, "2"),
		token.New(token.SEMICOLON, ";"),

		// Fourth statement
		token.New(token.LET, "let"),
		token.New(token.IDENTIFIER, "a"),
		token.New(token.BOOL, "bool"),
		token.New(token.ASSIGN, "="),
		token.New(token.BANG, "!"),
		token.New(token.BOOLEAN, "true"),
		token.New(token.SEMICOLON, ";"),

		token.New(token.EOF, ""),
	}
	scanner := New(source)
	for _, expectedToken := range expectedTokens {
		nextToken := scanner.ScanToken()
		if nextToken.Type != expectedToken.Type || nextToken.Value != expectedToken.Value {
			t.Fatalf("Invalid token. Expecting %q, got %q", expectedToken, nextToken)
		}
	}
}

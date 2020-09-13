package scanner

import (
	"github.com/bingo-lang/bingo/token"
	"strings"
	"testing"
)

func TestScanner(t *testing.T) {
	source := strings.NewReader(`
		let a int = 10 + 1 - 1 * 1 / 1;
		let b bool = true || false && true;
		let c bool = 1 > 2 || 1 >= 2 && 1 < 2 && 1 <= 2 && 1 == 2;
		let d bool = !true;
		let e bool = true == true;
		let f int = (10);
		{1};
	`)
	expectedTokens := []token.Token{
		// First statement.
		token.New(token.LET, "let"),
		token.New(token.IDENTIFIER, "a"),
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
		token.New(token.IDENTIFIER, "b"),
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
		token.New(token.IDENTIFIER, "c"),
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
		token.New(token.IDENTIFIER, "d"),
		token.New(token.BOOL, "bool"),
		token.New(token.ASSIGN, "="),
		token.New(token.BANG, "!"),
		token.New(token.BOOLEAN, "true"),
		token.New(token.SEMICOLON, ";"),

		// Fifth statement
		token.New(token.LET, "let"),
		token.New(token.IDENTIFIER, "e"),
		token.New(token.BOOL, "bool"),
		token.New(token.ASSIGN, "="),
		token.New(token.BOOLEAN, "true"),
		token.New(token.EQUAL, "=="),
		token.New(token.BOOLEAN, "true"),
		token.New(token.SEMICOLON, ";"),

		// Sixth statement
		token.New(token.LET, "let"),
		token.New(token.IDENTIFIER, "f"),
		token.New(token.INT, "int"),
		token.New(token.ASSIGN, "="),
		token.New(token.LPAREN, "("),
		token.New(token.INTEGER, "10"),
		token.New(token.RPAREN, ")"),
		token.New(token.SEMICOLON, ";"),

		// Seventh statement
		token.New(token.LBRACE, "{"),
		token.New(token.INTEGER, "1"),
		token.New(token.RBRACE, "}"),
		token.New(token.SEMICOLON, ";"),

		token.New(token.EOF, ""),
	}
	scanner := New(source)
	for _, expectedToken := range expectedTokens {
		nextToken := scanner.ScanToken()
		if !expectedToken.Equals(nextToken) {
			t.Fatalf("Invalid token. Expecting %q, got %q", expectedToken, nextToken)
		}
	}
}

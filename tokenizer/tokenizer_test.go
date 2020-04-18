package tokenizer

import (
	"testing"
)

func TestTokenizer(t *testing.T) {
	source := `
	let x = 10;
	let y = x == 10;
	`
	expectedTokens := []string{
		"let", "x", "=", "10", ";",
		"let", "y", "=", "x", "=", "=", "10", ";",
	}
	tokenizer, err := New(source)
	if err != nil {
		t.Fatal(err)
	}
	for _, expectedToken := range expectedTokens {
		nextToken := tokenizer.Next()
		if nextToken != expectedToken {
			t.Fatalf("Invalid token expecting %q, got %q", expectedToken, nextToken)
		}
	}
	if !tokenizer.EOF() {
		t.Fatalf("Expecting no more tokens but got %q", tokenizer.Next())
	}
}

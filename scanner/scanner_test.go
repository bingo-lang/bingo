package scanner

import (
	"github.com/bingo-lang/bingo/token"
	"testing"
)

func TestScanner(t *testing.T) {
	source := ""
	expected := []token.Token{}
	scanner, err := New(source)
	if err != nil {
		t.Fatal(err)
	}
	for _, token := range expected {
		if !equal(scanner.Token(), token) {
			t.Fatal("nooo")
		}
	}
	if !scanner.IsDone() {
		t.Fatal("nooo")
	}
}

func equal(a, b token.Token) bool {
	if a.Type != b.Type {
		return false
	}
	return true
}

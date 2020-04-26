package parser

import (
	"github.com/bingo-lang/bingo/ast"
	"strings"
	"testing"
)

func TestExpressionStatement(t *testing.T) {
	source := strings.NewReader(`1`)
	parser := New(source)
	expectedAst := &ast.Program{}
	program := parser.Parse()

	if program != expectedAst {
		t.Fatal("Different")
	}
}

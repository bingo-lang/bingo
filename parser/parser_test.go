package parser

import (
	"github.com/bingo-lang/bingo/ast"
	"strings"
	"testing"
)

func TestExpressionStatement(t *testing.T) {
	source := strings.NewReader(`1`)
	parser := New(source)
	expected := &ast.Program{
		Statements: []ast.Statement{
			&ast.StatementExpression{
				Expression: &ast.ExpressionPrefixInteger{
					Value: "1",
				},
			},
		},
	}
	gotten := parser.Parse()
	comparePrograms(t, expected, gotten)
}

func comparePrograms(t *testing.T, expected, gotten *ast.Program) {
	if len(gotten.Statements) != len(expected.Statements) {
		t.Fatalf("Expecting %d statements got %d", len(expected.Statements), len(gotten.Statements))
	}
}

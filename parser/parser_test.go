package parser

import (
	"github.com/bingo-lang/bingo/ast"
	"reflect"
	"strings"
	"testing"
)

func TestExpressionStatement(t *testing.T) {
	source := strings.NewReader(`237`)
	parser := New(source)
	gotten := parser.Parse()
	if len(gotten.Statements) != 1 {
		t.Fatalf("Expecting 1 statement, got %d", len(gotten.Statements))
	}
	statement, ok := (gotten.Statements[0]).(*ast.StatementExpression)
	if !ok {
		t.Fatalf("Expecting statement to be %q, got %q", "StatementExpression", reflect.TypeOf(statement))
	}
	expression, ok := statement.Expression.(*ast.ExpressionPrefixInteger)
	if !ok {
		t.Fatalf("Expecting expression to be %q, got %q", "ExpressionPrefixInteger", reflect.TypeOf(expression))
	}
	if expression.Value != "237" {
		t.Fatalf("Expecting expression prefix integer value to be %q, got %q", "237", "237")
	}
}

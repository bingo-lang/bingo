package parser

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/token"
	"reflect"
	"strings"
	"testing"
)

func TestExpressionPrefix(t *testing.T) {
	source := strings.NewReader(`+237`)
	parser := New(source)
	gotten := parser.Parse()
	if len(gotten.Statements) != 1 {
		t.Fatalf("Expecting 1 statement, got %d", len(gotten.Statements))
	}
	statement, ok := (gotten.Statements[0]).(*ast.StatementExpression)
	if !ok {
		t.Fatalf("Expecting statement to be %q, got %q", "StatementExpression", reflect.TypeOf(statement))
	}
	expressionPrefix, ok := statement.Expression.(*ast.ExpressionPrefix)
	if !ok {
		t.Fatalf("Expecting expression to be %q, got %q", "ExpressionPrefix", reflect.TypeOf(expressionPrefix))
	}
	if expressionPrefix.Operator.Type != token.PLUS {
		t.Fatalf("Expecting operator to be of type %q, got %q", token.MINUS, expressionPrefix.Operator.Type)
	}
	expressionInteger, ok := expressionPrefix.Expression.(*ast.ExpressionInteger)
	if !ok {
		t.Fatalf("Expecting expression to be %q, got %q", "ExpressionInteger", reflect.TypeOf(expressionInteger))
	}
	if expressionInteger.Value != "237" {
		t.Fatalf("Expecting expression prefix integer value to be %q, got %q", "237", expressionInteger.Value)
	}
}

func TestExpressionInteger(t *testing.T) {
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
	expressionInteger, ok := statement.Expression.(*ast.ExpressionInteger)
	if !ok {
		t.Fatalf("Expecting expressionInteger to be %q, got %q", "ExpressionInteger", reflect.TypeOf(expressionInteger))
	}
	if expressionInteger.Value != "237" {
		t.Fatalf("Expecting expressionInteger prefix integer value to be %q, got %q", "237", "237")
	}
}

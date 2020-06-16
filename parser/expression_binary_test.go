package parser

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/token"
	"reflect"
	"strings"
	"testing"
)

func TestExpressionInfixMinus(t *testing.T) {
	source := strings.NewReader(`11 - 10`)
	parser := New(source)
	gotten := parser.Parse()
	if len(gotten.Statements) != 1 {
		t.Fatalf("Expecting 1 statement, got %d", len(gotten.Statements))
	}
	statement, ok := (gotten.Statements[0]).(*ast.StatementExpression)
	if !ok {
		t.Fatalf("Expecting statement to be %q, got %q", "StatementExpression", reflect.TypeOf(statement))
	}
	expressionInfix, ok := statement.Expression.(*ast.ExpressionInfix)
	if !ok {
		t.Fatalf("Expecting expression to be %q, got %q", "ExpressionPrefix", reflect.TypeOf(expressionInfix))
	}
	if expressionInfix.Operator.Type != token.MINUS {
		t.Fatalf("Expecting operator to be of type %q, got %q", token.MINUS, expressionInfix.Operator.Type)
	}
	expressionLeft, ok := expressionInfix.ExpressionLeft.(*ast.ExpressionInteger)
	if !ok {
		t.Fatalf("Expecting expression to be %q, got %q", "ExpressionInteger", reflect.TypeOf(expressionLeft))
	}
	if expressionLeft.Value != "11" {
		t.Fatalf("Expecting expression prefix integer value to be %q, got %q", "11", expressionLeft.Value)
	}
	expressionRight, ok := expressionInfix.ExpressionRight.(*ast.ExpressionInteger)
	if !ok {
		t.Fatalf("Expecting expression to be %q, got %q", "ExpressionInteger", reflect.TypeOf(expressionRight))
	}
	if expressionRight.Value != "10" {
		t.Fatalf("Expecting expression prefix integer value to be %q, got %q", "10", expressionRight.Value)
	}

}

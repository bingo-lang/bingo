package evaluator

import (
	"github.com/bingo-lang/bingo/object"
	"github.com/bingo-lang/bingo/parser"
	"strings"
	"testing"
)

/*
func TestEvalExpressionPrefixMinus(t *testing.T) {
	source := strings.NewReader(`-1`)
	expected := object.Integer{Value: -1}
	parser := parser.New(source)
	program := parser.Parse()
	gotten, ok := Eval(program.Statements[0]).(*object.Integer)
	fmt.Println(gotten)
	if !ok {
		t.Fatalf("Expect an integer, got %q", gotten)
	}
	if gotten.Value != expected.Value {
		t.Fatalf("Expect value to be %q, got %q", expected, gotten)
	}
}
*/

func TestEvalExpressionPrefixInteger(t *testing.T) {
	source := strings.NewReader(`1`)
	expected := object.Integer{Value: 1}
	parser := parser.New(source)
	program := parser.Parse()
	gotten, ok := Eval(program.Statements[0]).(*object.Integer)
	if !ok {
		t.Fatalf("Expect an integer, got %q", gotten)
	}
	if gotten.Value != expected.Value {
		t.Fatalf("Expect value to be %d, got %d", expected.Value, gotten.Value)
	}
}

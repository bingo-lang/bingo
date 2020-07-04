package evaluator

import (
	"github.com/bingo-lang/bingo/object"
	"github.com/bingo-lang/bingo/parser"
	"strings"
	"testing"
)

func TestEvalExpression(t *testing.T) {
	testCases := []struct {
		source   string
		expected int
	}{
		{"-11", -11},
		{"10 + 1", 11},
		{"10 - 1", 9},
		{"10 * 1", 10},
		{"10 / 2", 5},
		{"+10 / 2", 5},
		{"10 + 1 * 1000", 1010},
	}
	for _, testCase := range testCases {
		source := strings.NewReader(testCase.source)
		expected := object.Integer{Value: testCase.expected}
		parser := parser.New(source)
		program := parser.ParseProgram()
		gotten, ok := Eval(program.Statements[0]).(*object.Integer)
		if !ok {
			t.Fatalf("Expect an integer, got %T", gotten)
		}
		if gotten.Value != expected.Value {
			t.Fatalf("Expect value to be %d, got %d", expected.Value, gotten.Value)
		}
	}
}

func TestEvalExpressionInteger(t *testing.T) {
	source := strings.NewReader(`1`)
	expected := object.Integer{Value: 1}
	parser := parser.New(source)
	program := parser.ParseProgram()
	gotten, ok := Eval(program.Statements[0]).(*object.Integer)
	if !ok {
		t.Fatalf("Expect an integer, got %q", gotten)
	}
	if gotten.Value != expected.Value {
		t.Fatalf("Expect value to be %d, got %d", expected.Value, gotten.Value)
	}
}

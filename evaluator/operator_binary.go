package evaluator

import (
	"fmt"
	"github.com/bingo-lang/bingo/object"
	"reflect"
)

func evalBinaryPlus(left, right object.Object) (object.Object, error) {
	op1, ok1 := left.(object.Integer)
	op2, ok2 := right.(object.Integer)
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("Expecting left and right operands to be Integers got %q of type %q and %q of type %q", left, left.Type(), right, right.Type())
	}
	return object.Integer{Value: op1.Value + op2.Value}, nil
}

func evalBinaryMinus(left, right object.Object) (object.Object, error) {
	op1, ok1 := left.(object.Integer)
	op2, ok2 := right.(object.Integer)
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("Expecting left and right operands to be Integers got %q of type %q and %q of type %q", left, left.Type(), right, right.Type())
	}
	return object.Integer{Value: op1.Value - op2.Value}, nil
}

func evalBinaryAsterisk(left, right object.Object) (object.Object, error) {
	op1, ok1 := left.(object.Integer)
	op2, ok2 := right.(object.Integer)
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("Expecting left and right operands to be Integers got %q of type %q and %q of type %q", left, left.Type(), right, right.Type())
	}
	return object.Integer{Value: op1.Value * op2.Value}, nil
}

func evalBinarySlash(left, right object.Object) (object.Object, error) {
	op1, ok1 := left.(object.Integer)
	op2, ok2 := right.(object.Integer)
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("Expecting left and right operands to be Integers got %q of type %q and %q of type %q", left, left.Type(), right, right.Type())
	}
	return object.Integer{Value: op1.Value / op2.Value}, nil
}

func evalBinaryGT(left, right object.Object) (object.Object, error) {
	op1, ok1 := left.(object.Integer)
	op2, ok2 := right.(object.Integer)
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("Expecting left and right operands to be Integers got %q of type %q and %q of type %q", left, left.Type(), right, right.Type())
	}
	return object.Boolean{Value: op1.Value > op2.Value}, nil
}

func evalBinaryGTE(left, right object.Object) (object.Object, error) {
	op1, ok1 := left.(object.Integer)
	op2, ok2 := right.(object.Integer)
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("Expecting left and right operands to be Integers got %q of type %q and %q of type %q", left, left.Type(), right, right.Type())
	}
	return object.Boolean{Value: op1.Value >= op2.Value}, nil
}

func evalBinaryLT(left, right object.Object) (object.Object, error) {
	op1, ok1 := left.(object.Integer)
	op2, ok2 := right.(object.Integer)
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("Expecting left and right operands to be Integers got %q of type %q and %q of type %q", left, left.Type(), right, right.Type())
	}
	return object.Boolean{Value: op1.Value < op2.Value}, nil
}

func evalBinaryLTE(left, right object.Object) (object.Object, error) {
	op1, ok1 := left.(object.Integer)
	op2, ok2 := right.(object.Integer)
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("Expecting left and right operands to be Integers got %q of type %q and %q of type %q", left, left.Type(), right, right.Type())
	}
	return object.Boolean{Value: op1.Value <= op2.Value}, nil
}

func evalBinaryEqual(left, right object.Object) (object.Object, error) {
	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		return nil, fmt.Errorf("Expecting left and right operands to be of same type got %q of type %q and %q of type %q", left, left.Type(), right, right.Type())
	}
	switch left.(type) {
	case object.Integer:
		op1 := left.(object.Integer)
		op2 := right.(object.Integer)
		return object.Boolean{Value: op1.Value == op2.Value}, nil
	case object.Boolean:
		op1 := left.(object.Boolean)
		op2 := right.(object.Boolean)
		return object.Boolean{Value: op1.Value == op2.Value}, nil
	default:
		return nil, fmt.Errorf("Expecting left and right operands to be Integers got %q of type %q and %q of type %q", left, left.Type(), right, right.Type())
	}
}

func evalBinaryAnd(left, right object.Object) (object.Object, error) {
	op1, ok1 := left.(object.Boolean)
	op2, ok2 := right.(object.Boolean)
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("Expecting left and right operands to be Booleans got %q of type %q and %q of type %q", left, left.Type(), right, right.Type())
	}
	return object.Boolean{Value: op1.Value && op2.Value}, nil
}

func evalBinaryOr(left, right object.Object) (object.Object, error) {
	op1, ok1 := left.(object.Boolean)
	op2, ok2 := right.(object.Boolean)
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("Expecting left and right operands to be Booleans got %q of type %q and %q of type %q", left, left.Type(), right, right.Type())
	}
	return object.Boolean{Value: op1.Value || op2.Value}, nil
}

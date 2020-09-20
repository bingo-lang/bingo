package evaluator

import (
	"github.com/bingo-lang/bingo/ast"
	"github.com/bingo-lang/bingo/environment"
	"github.com/bingo-lang/bingo/object"
)

func evalUnaryMinus(node ast.Expression, env *environment.Environment) (object.Object, error) {
	node, err := evalExpression(node, env)
	if err != nil {
		return nil, err
	}
	switch ob := node.(type) {
	case object.Integer:
		ob.Value = -ob.Value
		return ob, nil
	default:
		// This is an error
		return nil, nil
	}
}

func evalUnaryBang(node ast.Expression, env *environment.Environment) (object.Object, error) {
	node, err := evalExpression(node, env)
	if err != nil {
		return nil, err
	}
	switch ob := node.(type) {
	case object.Boolean:
		ob.Value = !ob.Value
		return ob, nil
	default:
		// This is an error
		return nil, nil
	}
}

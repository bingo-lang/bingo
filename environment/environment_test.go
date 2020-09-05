package environment

import (
	"testing"

	"github.com/bingo-lang/bingo/object"
)

func TestEnvironment(t *testing.T) {
	key1 := "random1"
	key2 := "random2"
	obj1 := object.Integer{Value: 1}
	obj2 := object.Integer{Value: 2}
	parent := New(nil)
	parent.Set(key1, obj1)
	parent.Set(key2, obj1)
	environment := New(parent)
	environment.Set(key2, obj2)
	if gotten := environment.Get(key2); gotten != obj2 {
		t.Fatalf("Expecting object %s, got %s instead", obj2, gotten)
	}
	if gotten := environment.Get(key1); gotten != obj1 {
		t.Fatalf("Expecting object %s from parent, got %s instead", obj1, gotten)
	}
}

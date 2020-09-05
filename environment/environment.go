package environment

import (
	"github.com/bingo-lang/bingo/object"
)

func New(parent *Environment) *Environment {
	store := make(map[string]object.Object)
	return &Environment{store: store, parent: parent}
}

type Environment struct {
	store  map[string]object.Object
	parent *Environment
}

func (env *Environment) Set(key string, value object.Object) {
	env.store[key] = value
}

func (env *Environment) Get(key string) object.Object {
	if value, found := env.store[key]; found {
		return value
	}
	if env.parent != nil {
		return env.parent.Get(key)
	}
	return nil
}

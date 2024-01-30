package container

import (
	"github.com/samber/do"
	"reflect"
)

//a container to contain service or functions or any other

//based on https://github.com/samber/do, features:
//Service registration
//Service invocation
//Service health check
//Service shutdown
//Service lifecycle hooks
//Named or anonymous services
//Eagerly or lazily loaded services
//Dependency graph resolution
//Default injector
//Injector cloning
//Service override
//Lightweight, no dependencies
//No code generation

type Container[T any] struct {
	Injector   *do.Injector
	registries map[string]do.Provider[T]
}

func NewContainer[T any]() Container[T] {
	c := Container[T]{
		Injector:   do.New(),
		registries: map[string]do.Provider[T]{},
	}
	return c
}

//https://blogtitle.github.io/go-generics/

func (c *Container[T]) Register(spell do.Provider[T]) {
	do.Provide(c.Injector, spell)
	c.registries[reflect.TypeOf(spell).Name()] = spell
}

//TODO: Autowire/Get/Object Lifecycle

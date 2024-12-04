package examples

import (
	"fmt"
	"testing"

	"github.com/skjdfhkskjds/depinject"
	"github.com/skjdfhkskjds/depinject/internal/testutils"
)

// This example demonstrates how to use the dependency injection
// framework to inject and resolve types that contain generic types.

type FooBarGenerics[FooT any] struct {
	foo FooT
	bar *Bar
}

func NewFooBarGenerics[FooT any]() *FooBarGenerics[FooT] {
	return &FooBarGenerics[FooT]{}
}

func (fb *FooBarGenerics[FooT]) Print() {
	fmt.Println("Hello from FooBar!")
}

func TestWithGenerics(t *testing.T) {
	container := depinject.NewContainer()

	// Supply a value into the container directly.
	testutils.RequireNoError(t, container.Supply(&Foo{}))

	// Provide a set of constructors into the container.
	testutils.RequireNoError(t, container.Provide(
		NewBar,
		NewFooBarGenerics[*Foo],
	))

	// Invoke a function with the dependencies injected
	// to retrieve the FooBarGenerics instance.
	var fooBarGenerics *FooBarGenerics[*Foo]
	testutils.RequireNoError(t, container.Invoke(&fooBarGenerics))
	fooBarGenerics.Print()
}

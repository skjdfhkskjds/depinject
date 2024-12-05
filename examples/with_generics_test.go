package examples

import (
	"testing"

	"github.com/skjdfhkskjds/depinject"
	"github.com/skjdfhkskjds/depinject/internal/testutils"
)

// This example demonstrates how to use the dependency injection
// framework to inject and resolve types that contain generic types.

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
	testutils.RequireNotNil(t, fooBarGenerics)
	fooBarGenerics.Print()
}

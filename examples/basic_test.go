package examples

import (
	"testing"

	"github.com/skjdfhkskjds/depinject"
	"github.com/stretchr/testify/require"
)

// This example demonstrates how to use the dependency injection
// framework to inject and resolve types that are hard types which
// match exactly.
//
// In this case, Foo, Bar and FooBar are all hard types which match
// have constructors requesting those types exactly in the arguments
// lists.
//
// We supply the value of Foo, and the constructors for Bar and FooBar
// into the container and request an instance of FooBar.

func TestBasic(t *testing.T) {
	container := depinject.NewContainer()

	// Supply a value into the container directly.
	require.NoError(t, container.Supply(&Foo{}))

	// Provide a set of constructors into the container.
	require.NoError(t, container.Provide(
		NewBar,
		NewFooBar,
	))

	// Invoke a function with the dependencies injected
	// to retrieve the FooBar instance.
	var fooBar *FooBar
	require.NoError(t, container.Invoke(&fooBar))
	fooBar.Print()
}

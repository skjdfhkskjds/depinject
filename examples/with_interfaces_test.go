package examples

import (
	"testing"

	"github.com/skjdfhkskjds/depinject"
	"github.com/stretchr/testify/require"
)

// This example demonstrates how to use the dependency injection
// framework to inject types that implement interfaces into
// other types.
//
// In this case, Foo is a hard type, Bar is a hard type implementing BarI
// and FooBar has a constructor that requires both Foo and BarI.
// Additionally, FooBarI is an interface that FooBar implements.
//
// We supply the constructors for Foo, Bar and FooBar into the container
// and request an instance of FooBarI.

func TestWithInterfaces(t *testing.T) {
	container := depinject.NewContainer()

	// Provide a set of constructors into the container.
	require.NoError(t, container.Provide(
		NewFoo,
		NewBar,
		NewFooBarWithBarI,
	))

	// Invoke a function with the dependencies injected
	// to retrieve the FooBar instance.
	var fooBar FooBarI
	require.NoError(t, container.Invoke(&fooBar))
	fooBar.Print()
}

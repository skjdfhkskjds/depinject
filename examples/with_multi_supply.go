package examples

import (
	"testing"

	"github.com/skjdfhkskjds/depinject"
	"github.com/stretchr/testify/require"
)

// This example is identical to the basic example but supplies
// multiple values into the container instead.

func TestWithMultiSupply(t *testing.T) {
	container := depinject.NewContainer()

	// Supply a value into the container directly.
	require.NoError(t, container.Supply(&Foo{}, &Bar{}))

	// Provide a set of constructors into the container.
	require.NoError(t, container.Provide(
		NewFooBar,
	))

	// Invoke a function with the dependencies injected
	// to retrieve the FooBar instance.
	var fooBar *FooBar
	require.NoError(t, container.Invoke(&fooBar))
	fooBar.Print()
}

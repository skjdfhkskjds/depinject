package examples

import (
	"testing"

	"github.com/skjdfhkskjds/depinject"
	"github.com/stretchr/testify/require"
)

// This is the same as the basic example, except that we do not
// create an explicit instance of the container.

func TestBasicNoInstance(t *testing.T) {
	// Supply a value into the container directly.
	require.NoError(t, depinject.Supply(&Foo{}))

	// Provide a set of constructors into the container.
	require.NoError(t, depinject.Provide(
		NewBar,
		NewFooBar,
	))

	// Invoke a function with the dependencies injected
	// to retrieve the FooBar instance.
	var fooBar *FooBar
	require.NoError(t, depinject.Invoke(&fooBar))
	fooBar.Print()
}

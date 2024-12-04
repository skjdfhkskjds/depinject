package examples

import (
	"testing"

	"github.com/skjdfhkskjds/depinject"
	"github.com/skjdfhkskjds/depinject/internal/testutils"
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
	testutils.RequireNoError(t, container.Supply(&Foo{}))

	// Provide a set of constructors into the container.
	testutils.RequireNoError(t, container.Provide(
		NewBar,
		NewFooBar,
	))

	// Invoke a function with the dependencies injected
	// to retrieve the FooBar instance.
	var fooBar *FooBar
	testutils.RequireNoError(t, container.Invoke(&fooBar))
	fooBar.Print()
}

// This is the same as the basic example, except that we do not
// create an explicit instance of the container.
func TestBasicNoInstance(t *testing.T) {
	// Supply a value into the container directly.
	testutils.RequireNoError(t, depinject.Supply(&Foo{}))

	// Provide a set of constructors into the container.
	testutils.RequireNoError(t, depinject.Provide(
		NewBar,
		NewFooBar,
	))

	// Invoke a function with the dependencies injected
	// to retrieve the FooBar instance.
	var fooBar *FooBar
	testutils.RequireNoError(t, depinject.Invoke(&fooBar))
	fooBar.Print()
}

// This example is identical to the basic example but supplies
// multiple values into the container instead.

func TestBasicMultiSupply(t *testing.T) {
	container := depinject.NewContainer()

	// Supply a value into the container directly.
	testutils.RequireNoError(t, container.Supply(&Foo{}, &Bar{}))

	// Provide a set of constructors into the container.
	testutils.RequireNoError(t, container.Provide(
		NewFooBar,
	))

	// Invoke a function with the dependencies injected
	// to retrieve the FooBar instance.
	var fooBar *FooBar
	testutils.RequireNoError(t, container.Invoke(&fooBar))
	fooBar.Print()
}

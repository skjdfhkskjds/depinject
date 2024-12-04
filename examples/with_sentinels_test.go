package examples

import (
	"testing"

	"github.com/skjdfhkskjds/depinject"
	"github.com/skjdfhkskjds/depinject/internal/test_utils"
	"github.com/stretchr/testify/require"
)

// This example demonstrates how to use the dependency injection
// framework to inject types into a constructor using the input
// sentinel struct pattern.

type FooBarWithIn struct {
	depinject.In

	Foo *Foo
	Bar *Bar
}

func NewFooBarWithIn(in FooBarWithIn) *FooBar {
	return NewFooBar(in.Foo, in.Bar)
}

type FooBarWithOut struct {
	depinject.Out

	Foo *Foo
	Bar *Bar
}

func NewFooBarWithOut() FooBarWithOut {
	return FooBarWithOut{
		Foo: &Foo{},
		Bar: &Bar{},
	}
}

func TestWithInSentinels(t *testing.T) {
	container := depinject.NewContainer(
		depinject.WithInSentinel(),
	)

	// Supply a value into the container directly.
	require.NoError(t, container.Supply(&Foo{}))

	// Provide a set of constructors into the container.
	require.NoError(t, container.Provide(
		NewBar,
		NewFooBarWithIn,
	))

	// Invoke a function with the dependencies injected
	// to retrieve the FooBar instance.
	var fooBar *FooBar
	require.NoError(t, container.Invoke(&fooBar))
	fooBar.Print()
}

func TestWithInSentinelsMultiple(t *testing.T) {
	test_utils.RunMultiWithoutSTDOUT(t, TestWithInSentinels, 100)
}

func TestWithOutSentinels(t *testing.T) {
	container := depinject.NewContainer(
		depinject.WithOutSentinel(),
	)

	// Provide a set of constructors into the container.
	require.NoError(t, container.Provide(
		NewFooBarWithOut,
		NewFooBar,
	))

	// Invoke a function with the dependencies injected
	// to retrieve the FooBar instance.
	var fooBar *FooBar
	require.NoError(t, container.Invoke(&fooBar))
	fooBar.Print()
	require.Nil(t, nil)
}

func TestWithOutSentinelsMultiple(t *testing.T) {
	test_utils.RunMultiWithoutSTDOUT(t, TestWithOutSentinels, 100)
}

func TestWithInAndOutSentinels(t *testing.T) {
	container := depinject.NewContainer(
		depinject.WithInSentinel(),
		depinject.WithOutSentinel(),
	)

	// Provide a set of constructors into the container.
	require.NoError(t, container.Provide(
		NewFooBarWithOut,
		NewFooBarWithIn,
	))

	// Invoke a function with the dependencies injected
	// to retrieve the FooBar instance.
	var fooBar *FooBar
	require.NoError(t, container.Invoke(&fooBar))
	fooBar.Print()
}

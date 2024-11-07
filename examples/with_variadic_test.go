package examples

import (
	"testing"

	"github.com/skjdfhkskjds/depinject"
	"github.com/stretchr/testify/require"
)

// This example demonstrates how to use the dependency injection
// framework to inject variadic arguments into a function.

func NewFooBarVariadic(foo *Foo, bars ...*Bar) *FooBar {
	return &FooBar{}
}

func NewMultiBar() []*Bar {
	return []*Bar{{}, {}}
}

func TestWithVariadicNoArgs(t *testing.T) {
	container := depinject.NewContainer()

	require.NoError(t, container.Supply(&Foo{}))
	require.NoError(t, container.Provide(NewFooBarVariadic))

	var fooBar *FooBar
	require.NoError(t, container.Invoke(&fooBar))
	fooBar.Print()
}

func TestWithVariadicOneArg(t *testing.T) {
	container := depinject.NewContainer()

	require.NoError(t, container.Supply(&Foo{}, &Bar{}))
	require.NoError(t, container.Provide(NewFooBarVariadic))
}

func TestWithVariadicMultipleArgs(t *testing.T) {
	container := depinject.NewContainer()

	require.NoError(t, container.Supply(&Foo{}))
	require.NoError(t, container.Provide(NewMultiBar, NewFooBarVariadic))
}

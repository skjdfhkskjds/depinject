package examples

import (
	"testing"

	"github.com/skjdfhkskjds/depinject"
	"github.com/stretchr/testify/require"
)

func NewFooBarVariadic(foo *Foo, bars ...*Bar) *FooBar {
	return &FooBar{}
}

func TestWithVariadic(t *testing.T) {
	container := depinject.NewContainer()

	require.NoError(t, container.Supply(&Foo{}))
	require.NoError(t, container.Provide(NewFooBarVariadic))

	var fooBar *FooBar
	require.NoError(t, container.Invoke(&fooBar))
	fooBar.Print()
}

package examples

// import (
// 	"testing"

// 	"github.com/skjdfhkskjds/depinject"
// 	"github.com/stretchr/testify/require"
// )

// // This example demonstrates how to use the dependency injection
// // framework to inject types into a constructor using the input
// // sentinel struct pattern.

// type FooBarSentinels struct {
// 	depinject.In

// 	Foo *Foo
// 	Bar *Bar
// }

// func NewFooBarWithSentinels(in FooBarSentinels) *FooBar {
// 	return NewFooBar(in.Foo, in.Bar)
// }

// func TestWithSentinels(t *testing.T) {
// 	container := depinject.NewContainer()

// 	// Supply a value into the container directly.
// 	require.NoError(t, container.Supply(&Foo{}))

// 	// Provide a set of constructors into the container.
// 	require.NoError(t, container.Provide(
// 		NewBar,
// 		NewFooBarWithSentinels,
// 	))

// 	// Invoke a function with the dependencies injected
// 	// to retrieve the FooBar instance.
// 	var fooBar *FooBar
// 	require.NoError(t, container.Invoke(&fooBar))
// 	fooBar.Print()
// }

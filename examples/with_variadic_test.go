package examples

import (
	"testing"

	"github.com/skjdfhkskjds/depinject"
	"github.com/skjdfhkskjds/depinject/internal/testutils"
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

	testutils.RequireNoError(t, container.Supply(&Foo{}))
	testutils.RequireNoError(t, container.Provide(NewFooBarVariadic))

	var fooBar *FooBar
	testutils.RequireNoError(t, container.Invoke(&fooBar))
	fooBar.Print()
}

func TestWithVariadicOneArg(t *testing.T) {
	container := depinject.NewContainer()

	testutils.RequireNoError(t, container.Supply(&Foo{}, &Bar{}))
	testutils.RequireNoError(t, container.Provide(NewFooBarVariadic))
}

func TestWithVariadicMultipleArgs(t *testing.T) {
	container := depinject.NewContainer()

	testutils.RequireNoError(t, container.Supply(&Foo{}))
	testutils.RequireNoError(t, container.Provide(NewMultiBar, NewFooBarVariadic))
}

func TestWithVariadicInferredListInSupply(t *testing.T) {
	container := depinject.NewContainer(depinject.WithListInference())

	testutils.RequireNoError(t, container.Supply(
		depinject.WithInterfaceInference(),
		depinject.WithListInference(),
		depinject.WithInSentinel(),
	))
	testutils.RequireNoError(t, container.Provide(depinject.NewContainer))

	var container2 *depinject.Container
	testutils.RequireNoError(t, container.Invoke(&container2))
	testutils.RequireNotNil(t, container2)
}

func TestWithVariadicInferredListInProvide(t *testing.T) {
	container := depinject.NewContainer(depinject.WithListInference())

	testutils.RequireNoError(t, container.Provide(
		depinject.NewContainer,
		depinject.WithInterfaceInference(),
		depinject.WithListInference(),
		depinject.WithInSentinel(),
	))

	var container2 *depinject.Container
	testutils.RequireNoError(t, container.Invoke(&container2))
	testutils.RequireNotNil(t, container2)
}

package main

import (
	"github.com/skjdfhkskjds/depinject"
	"github.com/skjdfhkskjds/depinject/examples"
)

// This example is identical to the basic example but supplies
// multiple values into the container instead.

func main() {
	container := depinject.NewContainer()

	// Supply a value into the container directly.
	if err := container.Supply(&examples.Foo{}, &examples.Bar{}); err != nil {
		panic(err)
	}

	// Provide a set of constructors into the container.
	if err := container.Provide(
		examples.NewFooBar,
	); err != nil {
		panic(err)
	}

	// Invoke a function with the dependencies injected
	// to retrieve the FooBar instance.
	var fooBar *examples.FooBar
	if err := container.Invoke(&fooBar); err != nil {
		panic(err)
	}
	fooBar.Print()
}

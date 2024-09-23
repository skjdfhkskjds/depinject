package main

import (
	"github.com/skjdfhkskjds/depinject"
	"github.com/skjdfhkskjds/depinject/examples"
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

func main() {
	container := depinject.NewContainer()

	// Supply a value into the container directly.
	if err := container.Supply(&examples.Foo{}); err != nil {
		panic(err)
	}

	// Provide a set of constructors into the container.
	if err := container.Provide(
		examples.NewBar,
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

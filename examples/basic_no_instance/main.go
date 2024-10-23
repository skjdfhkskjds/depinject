package main

import (
	"github.com/skjdfhkskjds/depinject"
	"github.com/skjdfhkskjds/depinject/examples"
)

// This is the same as the basic example, except that we do not
// create an explicit instance of the container.

func main() {
	// Supply a value into the container directly.
	if err := depinject.Supply(&examples.Foo{}); err != nil {
		panic(err)
	}

	// Provide a set of constructors into the container.
	if err := depinject.Provide(
		examples.NewBar,
		examples.NewFooBar,
	); err != nil {
		panic(err)
	}

	// Invoke a function with the dependencies injected
	// to retrieve the FooBar instance.
	var fooBar *examples.FooBar
	if err := depinject.Invoke(&fooBar); err != nil {
		panic(err)
	}
	fooBar.Print()
}

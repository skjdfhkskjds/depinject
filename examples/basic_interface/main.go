package main

import (
	"github.com/skjdfhkskjds/depinject"
	"github.com/skjdfhkskjds/depinject/examples"
)

// This example demonstrates how to use the dependency injection
// framework to inject types that implement interfaces into
// other types.
//
// In this case, Foo is a hard type, Bar is a hard type implementing BarI
// and FooBar has a constructor that requires both Foo and BarI.
// Additionally, FooBarI is an interface that FooBar implements.
//
// We supply the constructors for Foo, Bar and FooBar into the container
// and request an instance of FooBarI.

func main() {
	container := depinject.NewContainer()

	// Provide a set of constructors into the container.
	if err := container.Provide(
		examples.NewFoo,
		examples.NewBar,
		examples.NewFooBar,
	); err != nil {
		panic(err)
	}

	// Invoke a function with the dependencies injected
	// to retrieve the FooBar instance.
	var fooBar examples.FooBarI
	if err := container.Invoke(&fooBar); err != nil {
		panic(err)
	}
	fooBar.Print()
}

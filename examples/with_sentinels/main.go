package main

import (
	"github.com/skjdfhkskjds/depinject"
	"github.com/skjdfhkskjds/depinject/examples"
)

type FooBarSentinels struct {
	depinject.In

	Foo *examples.Foo
	Bar *examples.Bar
}

func NewFooBarWithSentinels(in FooBarSentinels) *examples.FooBar {
	return examples.NewFooBar(in.Foo, in.Bar)
}

func main() {
	container := depinject.NewContainer()

	// Supply a value into the container directly.
	if err := container.Supply(&examples.Foo{}); err != nil {
		panic(err)
	}

	// Provide a set of constructors into the container.
	if err := container.Provide(
		examples.NewBar,
		NewFooBarWithSentinels,
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

package main

import (
	"fmt"

	"github.com/skjdfhkskjds/depinject"
	"github.com/skjdfhkskjds/depinject/examples"
)

type FooBar[FooT any] struct {
	foo FooT
	bar *examples.Bar
}

func NewFooBar[FooT any]() *FooBar[FooT] {
	return &FooBar[FooT]{}
}

func (fb *FooBar[FooT]) Print() {
	fmt.Println("Hello from FooBar!")
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
		NewFooBar[*examples.Foo],
	); err != nil {
		panic(err)
	}

	// Invoke a function with the dependencies injected
	// to retrieve the FooBar instance.
	var fooBar *FooBar[*examples.Foo]
	if err := container.Invoke(&fooBar); err != nil {
		panic(err)
	}
	fooBar.Print()
}

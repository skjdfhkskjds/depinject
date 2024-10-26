package main

import (
	"errors"

	"github.com/skjdfhkskjds/depinject"
	"github.com/skjdfhkskjds/depinject/examples"
)

// This examples builds a valid container but forces an
// error on the invoke.

type Bar struct{}

func NewBar() (*Bar, error) {
	return &Bar{}, nil
}

type FooBar struct{}

func NewFooBar(foo *examples.Foo, bar *Bar) (*FooBar, error) {
	return &FooBar{}, errors.New(":(")
}

func main() {
	container := depinject.NewContainer()
	if err := container.Provide(
		examples.NewFoo,
		NewBar,
		NewFooBar,
	); err != nil {
		panic(err)
	}

	var fooBar *FooBar
	if err := container.Invoke(&fooBar); err == nil {
		panic("expected error")
	}
}

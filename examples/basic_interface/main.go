package main

import (
	"fmt"

	"github.com/skjdfhkskjds/depinject"
)

type Foo struct{}

func NewFoo() *Foo {
	return &Foo{}
}

type Bar struct{}

var _ BarI = (*Bar)(nil)

func NewBar(_ *Foo) *Bar {
	return &Bar{}
}

func (b *Bar) Bar() {}

type BarI interface {
	Bar()
}

type FooBar struct{}

type FooBarI interface {
	Print()
}

func NewFooBar(_ *Foo, _ BarI) *FooBar {
	return &FooBar{}
}

func (f *FooBar) Print() {
	fmt.Println("Hello from FooBar!")
}

func main() {
	container := depinject.NewContainer()

	// Supply a value into the container directly.
	if err := container.Supply(&Foo{}); err != nil {
		panic(err)
	}

	// Provide a set of constructors into the container.
	if err := container.Provide(
		NewBar,
		NewFooBar,
	); err != nil {
		panic(err)
	}

	// Invoke a function with the dependencies injected
	// to retrieve the FooBar instance.
	var fooBar FooBarI
	if err := container.Invoke(&fooBar); err != nil {
		panic(err)
	}
	fooBar.Print()
}

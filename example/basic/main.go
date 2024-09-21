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

func NewBar(_ *Foo) *Bar {
	return &Bar{}
}

type FooBar struct{}

func NewFooBar(_ *Foo, _ *Bar) *FooBar {
	return &FooBar{}
}

func (f *FooBar) Print() {
	fmt.Println("Hello from FooBar!")
}

func main() {
	container := depinject.NewContainer()

	if err := container.Provide(
		NewFoo,
		NewBar,
		NewFooBar,
	); err != nil {
		panic(err)
	}

	var fooBar *FooBar
	if err := container.Invoke(&fooBar); err != nil {
		panic(err)
	}
	fooBar.Print()
}

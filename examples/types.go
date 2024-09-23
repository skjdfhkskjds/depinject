package examples

import "fmt"

/* -------------------------------------------------------------------------- */
/*                                Hard Types                                  */
/* -------------------------------------------------------------------------- */

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

/* -------------------------------------------------------------------------- */
/*                                Interfaces                                  */
/* -------------------------------------------------------------------------- */

type BarI interface {
	Bar()
}

type FooBarI interface {
	Print()
}

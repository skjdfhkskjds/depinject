package examples

import (
	"errors"
	"fmt"
)

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

type FooBarGenerics[FooT any] struct {
	foo FooT
	bar *Bar
}

type FooBarGenericsOption[FooT any] func(*FooBarGenerics[FooT])

func WithFoo[FooT any](foo FooT) FooBarGenericsOption[FooT] {
	return func(fb *FooBarGenerics[FooT]) {
		fb.foo = foo
	}
}

func NewFooBarGenerics[FooT any]() *FooBarGenerics[FooT] {
	return &FooBarGenerics[FooT]{}
}

func NewFooBarGenericsWithOptions[FooT any](
	opts ...FooBarGenericsOption[FooT],
) (*FooBarGenerics[FooT], error) {
	if len(opts) == 0 {
		return nil, errors.New("no options provided")
	}
	fb := &FooBarGenerics[FooT]{}
	for _, opt := range opts {
		opt(fb)
	}
	return fb, nil
}

func (fb *FooBarGenerics[FooT]) Print() {
	fmt.Println("Hello from FooBar!")
}

/* -------------------------------------------------------------------------- */
/*                                Interfaces                                  */
/* -------------------------------------------------------------------------- */

func NewFooBarWithBarI(_ *Foo, _ BarI) *FooBar {
	return &FooBar{}
}

type BarI interface {
	Bar()
}

var _ BarI = (*Bar)(nil)

func (b *Bar) Bar() {}

type FooBarI interface {
	Print()
}

package examples

import (
	"errors"
	"testing"

	"github.com/skjdfhkskjds/depinject"
	"github.com/stretchr/testify/require"
)

// This examples builds a valid container but forces an
// error on the invoke.

type BarError struct{}

func NewBarError() (*BarError, error) {
	return &BarError{}, nil
}

type FooBarError struct{}

func NewFooBarError(foo *Foo, bar *BarError) (*FooBarError, error) {
	return &FooBarError{}, errors.New(":(")
}

func TestWithError(t *testing.T) {
	container := depinject.NewContainer()
	if err := container.Provide(
		NewFoo,
		NewBarError,
		NewFooBarError,
	); err != nil {
		panic(err)
	}

	var fooBar *FooBarError
	require.Error(t, container.Invoke(&fooBar))
}

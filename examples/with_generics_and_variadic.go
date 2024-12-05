package examples

import (
	"testing"

	"github.com/skjdfhkskjds/depinject"
	"github.com/skjdfhkskjds/depinject/internal/testutils"
)

func TestVariadicOptionsWithGenerics(t *testing.T) {
	container := depinject.NewContainer(
		depinject.WithListInference(),
	)

	testutils.RequireNoError(t, container.Supply(WithFoo(NewFoo())))
	testutils.RequireNoError(t, container.Provide(
		NewFooBarGenericsWithOptions[*Foo],
	))

	var fooBarGenerics *FooBarGenerics[*Foo]
	testutils.RequireNoError(t, container.Invoke(&fooBarGenerics))
	testutils.RequireNotNil(t, fooBarGenerics)
	fooBarGenerics.Print()
}

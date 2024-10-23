package depinject

import (
	"github.com/skjdfhkskjds/depinject/internal/depinject"
	"github.com/skjdfhkskjds/depinject/internal/depinject/types/sentinels"
)

type (
	// Container is the main entrypoint of this depinject library.
	// Its usage should be as follows:
	//
	//	container := NewContainer()
	//	container.Provide(constructor1, constructor2, ...)
	//	container.Invoke(func(dep1, dep2, ...) {
	//		// do something with the dependencies
	//	})
	Container = depinject.Container

	// In is a sentinel type used to indicate that a struct is
	// actually a container for various types that should be included
	// in the constructor's argument list.
	In = sentinels.In
)

// NewContainer returns a new, valid container.
var NewContainer = depinject.NewContainer

// Global container instance for users who would rather not
// manage their own container instances.
var c = NewContainer()

// Invoke invokes the given functions with the dependencies injected
// from the global container instance.
func Invoke(outputs ...any) error {
	return c.Invoke(outputs...)
}

// Provide provides the given constructors into the global container instance.
func Provide(constructors ...any) error {
	return c.Provide(constructors...)
}

// Supply supplies the given values into the global container instance.
func Supply(values ...any) error {
	return c.Supply(values...)
}

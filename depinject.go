package depinject

import "github.com/skjdfhkskjds/depinject/internal/depinject"

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
	In = depinject.In
)

// NewContainer returns a new, valid container.
var NewContainer = depinject.NewContainer

// Global container instance for users who would rather not
// manage their own container instances.
var c = NewContainer()

func Invoke(outputs ...any) error {
	return c.Invoke(outputs...)
}

func Provide(constructors ...any) error {
	return c.Provide(constructors...)
}

func Supply(values ...any) error {
	return c.Supply(values...)
}

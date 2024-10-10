package depinject

import "github.com/skjdfhkskjds/depinject/internal/depinject"

// Container is the main entrypoint of this depinject library.
// Its usage should be as follows:
//
//	container := NewContainer()
//	container.Provide(constructor1, constructor2, ...)
//	container.Invoke(func(dep1, dep2, ...) {
//		// do something with the dependencies
//	})
type Container = depinject.Container

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

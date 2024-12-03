package depinject

import (
	"github.com/skjdfhkskjds/depinject/internal/depinject/types/sentinels"
	depinject "github.com/skjdfhkskjds/depinject/internal/depinject3"
)

// Available types from this package.
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

// Available functions from this package.
var (
	// NewContainer returns a new, valid container.
	NewContainer = depinject.NewContainer

	// Instructs the container to enable the use of sentinel
	// structs in constructor arguments and parses the struct's
	// fields as constructor arguments.
	UseInSentinel = depinject.UseInSentinel

	// Instructs the container to enable the use of sentinel
	// structs in constructor outputs and parses the struct's
	// fields as constructor outputs.
	// TODO: Not implemented yet.
	UseOutSentinel = depinject.UseOutSentinel

	// Allows the container to match dependencies that are interfaces
	// to types which are implementations of those interfaces.
	InferInterfaces = depinject.InferInterfaces

	// Allows the container to have multiple constructors with the same
	// output type, and will process them as lists (slices or arrays).
	InferLists = depinject.InferLists
)

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

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

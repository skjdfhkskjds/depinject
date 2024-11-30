package depinject

import (
	"github.com/skjdfhkskjds/depinject/internal/depinject2/types"
)

const provideErrorName = "provide"

// Provide adds a set of constructors to the container.
// It returns an error if any of the constructors are invalid,
// or if adding them results in an invalid graph.
//
// Note: Constructors are added to the container in the order they are provided.
func (c *Container) Provide(constructors ...any) error {
	var err error
	for _, constructor := range constructors {
		if err = c.provide(constructor); err != nil {
			return err
		}
	}
	return nil
}

// provide adds a constructor to the container.
func (c *Container) provide(fn any) error {
	constructor, err := types.NewConstructor(fn)
	if err != nil {
		return err
	}

	// TODO: Handle all constructor modifiers
	return c.register(constructor, provideErrorName)
}

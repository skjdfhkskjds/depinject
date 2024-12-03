package depinject

import "github.com/skjdfhkskjds/depinject/internal/depinject/types"

const provideErrorName = "provide"

// Provide is a public function that allows for the injection of
// constructors into the container. Constructors are functions
// that return a value of some type.
func (c *Container) Provide(constructors ...any) error {
	for _, constructor := range constructors {
		if err := c.provide(constructor); err != nil {
			return c.interceptError(err)
		}
	}
	return nil
}

func (c *Container) provide(constructor any) error {
	node, err := types.NewNode(constructor)
	if err != nil {
		return newContainerError(err, provideErrorName, node.ID())
	}

	if err = c.register(node, provideErrorName); err != nil {
		return newContainerError(err, provideErrorName, node.ID())
	}

	return nil
}

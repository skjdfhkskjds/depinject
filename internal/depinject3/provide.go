package depinject

import "github.com/skjdfhkskjds/depinject/internal/depinject3/types"

const provideErrorName = "provide"

func (c *Container) Provide(constructors ...any) error {
	for _, constructor := range constructors {
		if err := c.provide(constructor); err != nil {
			return err
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

package depinject

import "github.com/skjdfhkskjds/depinject/internal/depinject3/types"

func (c *Container) register(
	node *types.Node,
	callerErrorName string,
) error {
	var err error
	if err = c.graph.AddVertex(node); err != nil {
		return newContainerError(err, callerErrorName, node.ID())
	}
	if err = c.registry.Register(node); err != nil {
		return newContainerError(err, callerErrorName, node.ID())
	}

	// When a new provider is registered, the container is no longer
	// invokable because we need to check if the new provider
	// introduces any circular dependencies.
	c.invokable = false
	return nil
}

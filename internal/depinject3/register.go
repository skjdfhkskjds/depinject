package depinject

import (
	"github.com/skjdfhkskjds/depinject/internal/depinject3/types"
)

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

// parseSentinels parses the sentinel structs in the node's constructor
// inputs and outputs and modifies the node to use the sentinel's fields
// instead of the sentinel itself.
func parseSentinels(node *types.Node) error {
	return nil
}

// parseInSentinel handles the in sentinel if applicable.
func parseInSentinel(node *types.Node) error {
	structs, ok := embedsInSentinel(node)
	if !ok {
		return nil
	}

	for _, _ = range structs {
		// constructor := structType.Constructor()
	}

	return nil
}

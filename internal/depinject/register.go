package depinject

import (
	"github.com/skjdfhkskjds/depinject/internal/depinject/types"
)

func (c *Container) register(
	node *types.Node,
	callerErrorName string,
) error {
	var err error
	if err = c.registerSentinelsForNode(node, callerErrorName); err != nil {
		return newContainerError(err, callerErrorName, node.ID())
	}

	// Register the node itself.
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

func (c *Container) registerSentinelsForNode(
	node *types.Node, callerErrorName string,
) error {
	// If in sentinels are enabled, register all applicable arguments
	// as nodes.
	if c.useInSentinel {
		sentinelNodes, err := parseInSentinels(node)
		if err != nil {
			return err
		}
		for _, n := range sentinelNodes {
			if err = c.register(n, callerErrorName); err != nil {
				return err
			}
		}
	}

	// If out sentinels are enabled, register all applicable outputs
	// as nodes.
	if c.useOutSentinel {
		sentinelNodes, err := parseOutSentinels(node)
		if err != nil {
			return err
		}
		for _, n := range sentinelNodes {
			if err = c.register(n, callerErrorName); err != nil {
				return err
			}
		}
	}

	return nil
}

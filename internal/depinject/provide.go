package depinject

import (
	"github.com/skjdfhkskjds/depinject/internal/depinject/types/node"
)

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
func (c *Container) provide(constructor any) error {
	node, err := node.New(constructor)
	if err != nil {
		return err
	}

	// Handle sentinels for the node
	if err = c.handleSentinelsForNode(node); err != nil {
		return err
	}

	// Add the node to the graph
	if err = c.graph.AddVertex(node); err != nil {
		return err
	}
	return c.registry.Register(node)
}

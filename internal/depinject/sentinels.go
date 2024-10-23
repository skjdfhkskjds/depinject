package depinject

import (
	"github.com/skjdfhkskjds/depinject/internal/depinject/types/node"
	"github.com/skjdfhkskjds/depinject/internal/depinject/types/sentinels"
	"github.com/skjdfhkskjds/depinject/internal/reflect"
)

// handleSentinelsForNode takes a given node, and injects the node's
// sentinels as new nodes into the container.
func (c *Container) handleSentinelsForNode(n *node.Node) error {
	for _, dep := range n.Dependencies() {
		if !sentinels.IsIn(dep) {
			continue
		}
		if err := c.handleIn(dep); err != nil {
			return err
		}
	}

	return nil
}

// handleIn takes a type that is an input sentinel, and
// generates a new struct and constructor for it, which is
// then provided to the container.
//
// input sentinels are treated as additional nodes whose
// constructors are a maximal list of the struct's fields
func (c *Container) handleIn(t reflect.Type) error {
	// If the container does not yet have the dummy sentinel,
	// supply it.
	if !c.hasSentinels {
		if err := c.supply(sentinels.InOut); err != nil {
			return err
		}
		c.hasSentinels = true
	}

	s, err := reflect.NewStruct(t)
	if err != nil {
		return err
	}

	return c.Provide(s.Constructor())
}

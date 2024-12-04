package depinject

import (
	"github.com/skjdfhkskjds/depinject/internal/depinject/types"
	"github.com/skjdfhkskjds/depinject/internal/errors"
	"github.com/skjdfhkskjds/depinject/internal/reflect"
)

const buildErrorName = "build"

func (c *Container) build() error {
	// iterate through every node in the graph and create incoming
	// edges for each node's dependencies
	for _, node := range c.graph.Vertices() {
		for _, dep := range node.Dependencies() {
			if err := c.buildDependencyForNode(node, dep); err != nil {
				return newContainerError(err, buildErrorName, node.ID())
			}
		}
	}

	nodes, err := c.graph.TopologicalSort()
	if err != nil {
		return err
	}
	c.sortedNodes = nodes
	return nil
}

// buildDependencyForNode creates edges from all providers of a particular
// node's dependency.
func (c *Container) buildDependencyForNode(
	node *types.Node,
	dep *reflect.Arg,
) error {
	// Search the registry for the dependency
	providers, err := c.registry.Lookup(dep.Type, dep.IsVariadic)
	if err != nil {
		return err
	}

	// If the container does not support array inferencing,
	// there should be at most one provider.
	if (!c.inferLists || !(dep.IsArray || dep.IsSlice)) && len(providers) > 1 {
		return errors.Newf(expected1ProviderErrMsg, len(providers))
	}

	for _, provider := range providers {
		// Don't create an edge from a node to itself.
		if provider == node {
			continue
		}
		if err := c.graph.AddEdge(provider, node); err != nil {
			return err
		}
	}

	return nil
}

package depinject

import (
	"reflect"

	"github.com/skjdfhkskjds/depinject/depinject/types"
	"github.com/skjdfhkskjds/depinject/internal/graph"
)

// A Container is a dependency injection container.
// Container usage should be as follows:
//
//	container := New()
//	container.Provide(constructor1, constructor2, ...)
//	container.Invoke(func(dep1, dep2, ...) {
//		// do something with the dependencies
//	})
type Container struct {
	graph *graph.Graph[*types.Node]

	// nodes is a map of a particular output type in
	// the container to the node which produces it.
	nodes map[reflect.Type]*types.Node
}

// New creates a new container.
func New() *Container {
	return &Container{
		graph: graph.New[*types.Node](),
		nodes: make(map[reflect.Type]*types.Node),
	}
}

// build builds the container by iterating through every
// node, and creating edges in the internal graph representation
// based on the dependencies and outputs of each node.
func (c *Container) build() error {
	for _, node := range c.nodes {
		for _, dep := range node.Dependencies() {
			source, ok := c.nodes[dep]
			if !ok {
				return types.NewError(
					ErrMissingDependency,
					node.ID(),
					dep.Name(),
				)
			}

			// Add the edge to the graph. If the edge violates
			// the acyclicity constraint, return an error.
			if err := c.graph.AddEdge(source, node); err != nil {
				return types.NewError(
					err,
					node.ID(),
					dep.Name(),
				)
			}
		}
	}

	return nil
}

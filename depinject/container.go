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

// NewContainer creates a new container.
func NewContainer() *Container {
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

// resolve resolves the container by iterating through every
// node in the container and executing them in a topological order.
func (c *Container) resolve() error {
	order, err := c.graph.TopologicalSort()
	if err != nil {
		return err
	}

	for _, node := range order {
		// Get the dependencies of the node.
		depTypes := node.Dependencies()
		deps := make([]any, 0, len(depTypes))
		for _, dep := range depTypes {
			value, err := c.nodes[dep].ValueOf(dep)
			if err != nil {
				return err
			}
			deps = append(deps, value)
		}

		// Execute the node with the dependencies.
		if err := node.Execute(deps...); err != nil {
			return err
		}
	}

	return nil
}

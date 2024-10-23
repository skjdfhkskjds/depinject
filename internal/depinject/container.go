package depinject

import (
	"github.com/skjdfhkskjds/depinject/internal/depinject/types/errors"
	"github.com/skjdfhkskjds/depinject/internal/depinject/types/node"
	"github.com/skjdfhkskjds/depinject/internal/graph"
)

// A Container is a dependency injection container.
type Container struct {
	// The internal graph representation of the container.
	graph *graph.Graph[*node.Node]

	// The node registry of the container.
	registry *node.Registry

	// Whether the container requires sentinels.
	hasSentinels bool
}

// NewContainer creates a new container.
func NewContainer() *Container {
	return &Container{
		graph:        graph.New[*node.Node](),
		registry:     node.NewRegistry(),
		hasSentinels: false,
	}
}

// build builds the container by iterating through every
// node, and creating edges in the internal graph representation
// based on the dependencies and outputs of each node.
func (c *Container) build() error {
	for _, node := range c.registry.Nodes() {
		for _, dep := range node.Dependencies() {
			source, err := c.registry.Get(dep)
			if err != nil {
				return errors.New(err, node.ID(), dep.Name())
			}

			// Add the edge to the graph. If the edge violates
			// the acyclicity constraint, return an error.
			if err := c.graph.AddEdge(source, node); err != nil {
				return errors.New(
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
			source, err := c.registry.Get(dep)
			if err != nil {
				return errors.New(err, node.ID(), dep.Name())
			}

			value, err := source.ValueOf(dep)
			if err != nil {
				return errors.New(err, node.ID(), dep.Name())
			}

			// Append the underlying casted value to deps
			deps = append(deps, value.Interface())
		}

		// Execute the node with the dependencies.
		if err := node.Execute(deps...); err != nil {
			return err
		}
	}

	return nil
}

package depinject

import (
	"github.com/skjdfhkskjds/depinject/internal/depinject/types/errors"
	"github.com/skjdfhkskjds/depinject/internal/depinject/types/node"
	"github.com/skjdfhkskjds/depinject/internal/depinject/types/sentinels"
	"github.com/skjdfhkskjds/depinject/internal/graph"
)

const (
	buildErrorName   = "build"
	resolveErrorName = "resolve"
)

// A Container is a dependency injection container.
type Container struct {
	// The internal graph representation of the container.
	graph *graph.DAG[*node.Node]

	// The node registry of the container.
	registry *node.Registry

	// Whether the container requires sentinels.
	hasIn  bool
	hasOut bool
}

// NewContainer creates a new container.
func NewContainer() *Container {
	return &Container{
		graph:    graph.NewDAG[*node.Node](),
		registry: node.NewRegistry(),
		hasIn:    false,
		hasOut:   false,
	}
}

// build builds the container by iterating through every
// node, and creating edges in the internal graph representation
// based on the dependencies and outputs of each node.
func (c *Container) build() error {
	// Before building, supply the sentinels.
	if c.hasIn {
		if err := c.supply(sentinels.In{}); err != nil {
			return err
		}
	}
	if c.hasOut {
		if err := c.supply(sentinels.Out{}); err != nil {
			return err
		}
	}

	for _, node := range c.registry.Nodes() {
		for _, dep := range node.Dependencies() {
			source, err := c.registry.Get(dep)
			if err != nil {
				return errors.New(err, buildErrorName, node.ID(), dep.Name())
			}

			// Add the edge to the graph. If the edge violates
			// the acyclicity constraint, return an error.
			if err := c.graph.AddEdge(source, node); err != nil {
				return errors.New(err, buildErrorName, node.ID(), dep.Name())
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
				return errors.New(err, resolveErrorName, node.ID(), dep.Name())
			}

			value, err := source.ValueOf(dep)
			if err != nil {
				return errors.New(err, resolveErrorName, node.ID(), dep.Name())
			}

			// Append the underlying casted value to deps
			deps = append(deps, value.Interface())
		}

		// Execute the node with the dependencies.
		if err := node.Execute(deps...); err != nil {
			return errors.New(err, resolveErrorName, node.ID())
		}
	}

	return nil
}

// addNode adds a node to the container.
func (c *Container) addNode(node *node.Node) error {
	if err := c.graph.AddVertex(node); err != nil {
		return err
	}

	return c.registry.Register(node)
}

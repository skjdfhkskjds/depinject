package depinject

import (
	"fmt"

	"github.com/skjdfhkskjds/depinject/internal/depinject3/types"
	"github.com/skjdfhkskjds/depinject/internal/graph"
)

type Container struct {
	graph    *graph.DAG[*types.Node]
	registry *types.Registry

	// Whether the container is ready to be invoked.
	invokable bool

	// Sorted nodes in topological order.
	sortedNodes []*types.Node

	// Options
	// Instructs the container to enable the use of sentinel
	// structs in constructor arguments and parses the struct's
	// fields as constructor arguments.
	useInSentinel bool

	// Instructs the container to enable the use of sentinel
	// structs in constructor outputs and parses the struct's
	// fields as constructor outputs.
	// TODO: Not implemented yet.
	useOutSentinel bool

	// Allows the container to match dependencies that are interfaces
	// to types which are implementations of those interfaces.
	inferInterfaces bool

	// Allows the container to have multiple constructors with the same
	// output type, and will process them as lists (slices or arrays).
	inferLists bool
}

func NewContainer(opts ...Option) *Container {
	c := &Container{
		invokable:       false,
		sortedNodes:     nil,
		useInSentinel:   false,
		useOutSentinel:  false,
		inferInterfaces: false,
		inferLists:      false,
	}
	for _, opt := range opts {
		c = opt(c)
	}

	c.graph = graph.NewDAG[*types.Node]()
	c.registry = types.NewRegistry(c.inferLists, c.inferInterfaces)
	return c
}

func (c *Container) Dump() {
	fmt.Println(c.registry.Dump())
}

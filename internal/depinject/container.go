package depinject

import (
	"io"
	"os"

	"github.com/skjdfhkskjds/depinject/internal/depinject/types"
	"github.com/skjdfhkskjds/depinject/internal/graph"
)

type Container struct {
	graph    *graph.DAG[*types.Node]
	registry *types.Registry

	// The writer to dump the container's info to.
	writer io.Writer

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
	useOutSentinel bool

	// Allows the container to match dependencies that are interfaces
	// to types which are implementations of those interfaces.
	inferInterfaces bool

	// Allows the container to have multiple constructors with the same
	// output type, and will process them as lists (slices or arrays).
	inferLists bool
}

// DefaultContainer returns a new container with the default options.
func DefaultContainer() *Container {
	return &Container{
		writer:          os.Stdout,
		invokable:       false,
		sortedNodes:     nil,
		useInSentinel:   false,
		useOutSentinel:  false,
		inferInterfaces: false,
		inferLists:      false,
	}
}

// NewContainer returns a new container with the given options.
func NewContainer(opts ...Option) *Container {
	c := DefaultContainer()
	for _, opt := range opts {
		opt(c)
	}

	c.graph = graph.NewDAG[*types.Node]()
	c.registry = types.NewRegistry(c.inferLists, c.inferInterfaces)
	return c
}

// Dump dumps the container's registry into the writer.
func (c *Container) Dump() {
	c.writer.Write([]byte(c.registry.Dump()))
}

// Destroy destroys the container and frees its memory.
func (c *Container) Destroy() {
	c.graph = nil
	c.registry = nil
	c.sortedNodes = nil
	c = nil
}

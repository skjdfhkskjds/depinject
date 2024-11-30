package depinject

import (
	"github.com/skjdfhkskjds/depinject/internal/depinject2/types"
	"github.com/skjdfhkskjds/depinject/internal/graph"
)

type Container struct {
	graph    *graph.DAG[*types.Constructor]
	registry *types.Registry

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

	// Allows the container to have multiple constructors with the
	// same output type, and will process them as slices.
	inferSlices bool
}

func New(opts ...Option) *Container {
	c := &Container{
		graph: graph.NewDAG[*types.Constructor](),
	}

	for _, opt := range opts {
		c = opt(c)
	}

	return c
}

func (c *Container) register(
	constructor *types.Constructor,
	callerErrorName string,
) error {
	var err error
	if err = c.graph.AddVertex(constructor); err != nil {
		return newContainerError(err, callerErrorName, constructor.ID())
	}
	if err = c.registry.Register(constructor); err != nil {
		return newContainerError(err, callerErrorName, constructor.ID())
	}

	return nil
}

package depinject

import (
	"fmt"
	"log"
	"strings"

	"github.com/skjdfhkskjds/depinject/internal/depinject/types"
	"github.com/skjdfhkskjds/depinject/internal/graph"
)

type Container struct {
	graph    *graph.DAG[*types.Node]
	registry *types.Registry

	// The logger used handle the container's error info.
	logger *log.Logger

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
		logger:          log.Default(),
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

	c.graph = graph.NewDAG[*types.Node](c.inferLists)
	c.registry = types.NewRegistry(c.inferLists, c.inferInterfaces)
	return c
}

// Destroy destroys the container and frees its memory.
func (c *Container) Destroy() {
	c.graph = nil
	c.registry = nil
	c.sortedNodes = nil
	c = nil
}

const (
	errHeaderText      = "Depinject Error"
	registryHeaderText = "Registry Contents"
)

// interceptError intercepts an error and performs the container's
// error handling routine. It then continues to propagate the error.
func (c *Container) interceptError(receivedErr error) error {
	if receivedErr == nil {
		return nil
	}

	// Create header with dynamic width
	errStr := receivedErr.Error()
	regDump := c.registry.Dump()
	maxLen := 0
	// Find max line length across all lines
	for _, line := range strings.Split(errStr, "\n") {
		maxLen = max(maxLen, len(line))
	}
	for _, line := range strings.Split(regDump, "\n") {
		maxLen = max(maxLen, len(line))
	}

	// Add padding for header text to center them
	errPadding := strings.Repeat(" ", (maxLen-len(errHeaderText))/2)
	regPadding := strings.Repeat(" ", (maxLen-len(registryHeaderText))/2)
	centeredErrHeader := errPadding + errHeaderText
	centeredRegHeader := regPadding + registryHeaderText

	border := strings.Repeat("=", maxLen)
	output := fmt.Sprintf(
		"\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s",
		border, centeredErrHeader, border, errStr,
		border, centeredRegHeader, border, regDump,
	)
	c.logger.Println(output)

	// Return the error to be handled by the caller.
	return receivedErr
}

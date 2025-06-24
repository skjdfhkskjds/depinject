package depinject

import (
	"log"
)

// Sets the logger to dump the container's info to.
func (c *Container) WithLogger(l *log.Logger) *Container {
	c.logger = l
	return c
}

// Instructs the container to enable the use of sentinel
// structs in constructor arguments and parses the struct's
// fields as constructor arguments.
func (c *Container) WithInSentinel() *Container {
	c.useInSentinel = true
	return c
}

// Instructs the container to enable the use of sentinel
// structs in constructor outputs and parses the struct's
// fields as constructor outputs.
func (c *Container) WithOutSentinel() *Container {
	c.useOutSentinel = true
	return c
}

// Allows the container to match dependencies that are interfaces
// to types which are implementations of those interfaces.
func (c *Container) WithInterfaceInference() *Container {
	c.inferInterfaces = true
	c.registry.WithInferInterfaces()
	return c
}

// Allows the container to have multiple constructors with the same
// output type, and will process them as lists (slices or arrays).
func (c *Container) WithListInference() *Container {
	c.inferLists = true
	c.graph.WithNonUniqueVertices()
	c.registry.WithInferLists()
	return c
}

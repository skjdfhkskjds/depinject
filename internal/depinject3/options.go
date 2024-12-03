package depinject

type Option func(*Container) *Container

// Instructs the container to enable the use of sentinel
// structs in constructor arguments and parses the struct's
// fields as constructor arguments.
func UseInSentinel(c *Container) *Container {
	c.useInSentinel = true
	return c
}

// Instructs the container to enable the use of sentinel
// structs in constructor outputs and parses the struct's
// fields as constructor outputs.
// TODO: Not implemented yet.
func UseOutSentinel(c *Container) *Container {
	c.useOutSentinel = true
	return c
}

// Allows the container to match dependencies that are interfaces
// to types which are implementations of those interfaces.
func InferInterfaces(c *Container) *Container {
	c.inferInterfaces = true
	return c
}

// Allows the container to have multiple constructors with the same
// output type, and will process them as lists (slices or arrays).
func InferLists(c *Container) *Container {
	c.inferLists = true
	return c
}

package depinject

type Option func(*Container) *Container

// Option to enable the use of sentinel structs in constructor.
func UseInSentinel(c *Container) *Container {
	c.useInSentinel = true
	return c
}

// Option to enable the use of sentinel structs in constructor outputs.
func UseOutSentinel(c *Container) *Container {
	c.useOutSentinel = true
	return c
}

// Option to enable the container to infer interfaces.
func InferInterfaces(c *Container) *Container {
	c.inferInterfaces = true
	return c
}

// Option to enable the container to infer slices.
func InferSlices(c *Container) *Container {
	c.inferSlices = true
	return c
}

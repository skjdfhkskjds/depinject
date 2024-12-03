package depinject

import "log"

type Option func(*Container)

// Sets the logger to dump the container's info to.
func WithLogger(l *log.Logger) Option {
	return func(c *Container) {
		c.logger = l
	}
}

// Instructs the container to enable the use of sentinel
// structs in constructor arguments and parses the struct's
// fields as constructor arguments.
func WithInSentinel() Option {
	return func(c *Container) {
		c.useInSentinel = true
	}
}

// Instructs the container to enable the use of sentinel
// structs in constructor outputs and parses the struct's
// fields as constructor outputs.
// TODO: Not implemented yet.
func WithOutSentinel() Option {
	return func(c *Container) {
		c.useOutSentinel = true
	}
}

// Allows the container to match dependencies that are interfaces
// to types which are implementations of those interfaces.
func WithInterfaceInference() Option {
	return func(c *Container) {
		c.inferInterfaces = true
	}
}

// Allows the container to have multiple constructors with the same
// output type, and will process them as lists (slices or arrays).
func WithListInference() Option {
	return func(c *Container) {
		c.inferLists = true
	}
}

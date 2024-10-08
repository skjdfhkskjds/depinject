package node

import "errors"

var (
	// ErrValueNotFound is returned when a value does not exist in the node.
	ErrValueNotFound = errors.New("value not found")

	// ErrMissingDependency is returned when a dependency is not found in the node.
	ErrMissingDependency = errors.New("missing dependency")

	// ErrMultipleImplementations is returned when multiple implementations
	// of an interface are found.
	ErrMultipleImplementations = errors.New("multiple implementations found")
)

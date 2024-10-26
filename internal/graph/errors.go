package graph

import "errors"

var (
	// ErrVertexNotFound is returned when a vertex is not found in the graph.
	ErrVertexNotFound = errors.New("vertex not found")

	// ErrVertexAlreadyExists is returned when a vertex already exists in the graph.
	ErrVertexAlreadyExists = errors.New("vertex already exists")

	// ErrAcyclicConstraintViolation is returned when adding an edge
	// would violate the acyclic constraint of the graph.
	ErrAcyclicConstraintViolation = errors.New("adding this edge would create a cycle")
)

package graph

// Vertex represents a vertex in a graph.
type Vertex[T any] interface {
	// ID returns the unique identifier for the vertex.
	ID() string
	// Incoming returns a list of vertices that have an
	// incoming edge to this vertex.
	Incoming() []T
}

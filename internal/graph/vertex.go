package graph

// Vertex represents a vertex in a graph.
type Vertex[T any] interface {
	// ID returns the unique identifier for the vertex.
	ID() string
}

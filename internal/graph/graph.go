package graph

import "github.com/dominikbraun/graph"

// Graph is a directed acyclic graph.
type Graph[VertexT Vertex[VertexT]] struct {
	graph graph.Graph[string, VertexT]
}

// New creates a new graph.
func New[VertexT Vertex[VertexT]]() *Graph[VertexT] {
	return &Graph[VertexT]{
		graph: graph.New(
			graph.Hash[string, VertexT](
				func(v VertexT) string {
					return v.ID()
				},
			),
			graph.Directed(),
			graph.Acyclic(),
		),
	}
}

// AddVertex adds a vertex to the graph.
func (g *Graph[VertexT]) AddVertex(v VertexT) error {
	return g.graph.AddVertex(v)
}

// AddEdge adds a directed edge of from -> to to the graph.
func (g *Graph[VertexT]) AddEdge(from, to VertexT) error {
	return g.graph.AddEdge(from.ID(), to.ID())
}

// TopologicalSort returns a list of vertices in topological order.
func (g *Graph[VertexT]) TopologicalSort() ([]VertexT, error) {
	order, err := graph.TopologicalSort(g.graph)
	if err != nil {
		return nil, err
	}

	vertices := make([]VertexT, len(order))
	for i, id := range order {
		vertices[i], err = g.graph.Vertex(id)
		if err != nil {
			return nil, err
		}
	}

	return vertices, nil
}

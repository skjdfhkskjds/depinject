package graph

// Graph structure
type Graph[VertexT Vertex] struct {
	vertices      map[string]VertexT
	adjacencyList map[string][]VertexT
}

// NewGraph creates a new Graph instance
func New[VertexT Vertex]() *Graph[VertexT] {
	return &Graph[VertexT]{
		vertices:      make(map[string]VertexT),
		adjacencyList: make(map[string][]VertexT),
	}
}

// AddVertex adds a vertex to the graph
func (g *Graph[VertexT]) AddVertex(v VertexT) error {
	if !g.hasVertex(v) {
		g.vertices[v.ID()] = v
		g.adjacencyList[v.ID()] = []VertexT{}
		return nil
	}

	return ErrVertexAlreadyExists
}

// AddEdge adds a directed edge from src to dest
func (g *Graph[VertexT]) AddEdge(src, dest VertexT) error {
	if !g.hasVertex(src) || !g.hasVertex(dest) {
		return ErrVertexNotFound
	}

	// Check for cycles before adding the edge
	if g.hasPath(dest.ID(), src.ID()) {
		return ErrAcyclicConstraintViolation
	}

	g.adjacencyList[src.ID()] = append(g.adjacencyList[src.ID()], dest)
	return nil
}

// hasPath checks if there's a path from start to end
func (g *Graph[VertexT]) hasPath(start, end string) bool {
	visited := make(map[string]bool)
	var dfs func(current string) bool
	dfs = func(current string) bool {
		if current == end {
			return true
		}
		visited[current] = true
		for _, neighbor := range g.adjacencyList[current] {
			if !visited[neighbor.ID()] {
				if dfs(neighbor.ID()) {
					return true
				}
			}
		}
		return false
	}
	return dfs(start)
}

// hasVertex checks if a vertex exists in the graph.
func (g *Graph[VertexT]) hasVertex(v VertexT) bool {
	_, exists := g.vertices[v.ID()]
	return exists
}

// TopologicalSort returns a topological ordering of the vertices in the graph.
// If the graph contains a cycle, it returns an error.
func (g *Graph[VertexT]) TopologicalSort() ([]VertexT, error) {
	status := make(map[string]int) // 0: unseen, 1: visiting, 2: visited
	stack := make([]VertexT, 0)

	var visit func(v string) error
	visit = func(v string) error {
		status[v] = 1

		for _, w := range g.adjacencyList[v] {
			if status[w.ID()] == 0 {
				visit(w.ID())
			} else if status[w.ID()] == 1 {
				return ErrAcyclicConstraintViolation
			}
		}

		status[v] = 2
		stack = append(stack, g.vertices[v])
		return nil
	}

	for v := range g.adjacencyList {
		if status[v] == 0 {
			if err := visit(v); err != nil {
				return nil, err
			}
		}
	}

	return stack, nil
}

// getVertex is a helper function to get the VertexT object from its ID
func (g *Graph[VertexT]) getVertex(id string) (VertexT, error) {
	v, ok := g.vertices[id]
	if !ok {
		return v, ErrVertexNotFound
	}

	return v, nil
}

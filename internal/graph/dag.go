package graph

import "github.com/skjdfhkskjds/depinject/internal/utils"

type DAG[VertexT Vertex] struct {
	// vertices is a map of vertex IDs to vertices.
	// Under the uniqueness constraint, each vertex ID maps to a single
	// vertex. Without uniqueness, the directed graph is calculated with
	// only a single vertex, but returns the whole set of vertices in order
	// when topologically sorted.
	vertices *utils.OrderedMap[string, []VertexT]
	edges    map[string][]VertexT
	indegree map[string]int

	totalVertices int

	enforceUniqueVertices bool
}

// NewDAG creates a new empty DAG.
func NewDAG[VertexT Vertex](enforceUniqueVertices bool) *DAG[VertexT] {
	return &DAG[VertexT]{
		vertices:              utils.NewOrderedMap[string, []VertexT](),
		edges:                 make(map[string][]VertexT),
		indegree:              make(map[string]int),
		enforceUniqueVertices: enforceUniqueVertices,
	}
}

// Vertices returns all vertices in the DAG.
func (g *DAG[VertexT]) Vertices() []VertexT {
	vertices := make([]VertexT, g.totalVertices)
	i := 0
	for _, v := range g.vertices.Keys() {
		verticesForKey, _ := g.vertices.Get(v)
		for _, vertex := range verticesForKey {
			vertices[i] = vertex
			i++
		}
	}
	return vertices
}

// AddVertex adds a new vertex to the DAG.
func (g *DAG[VertexT]) AddVertex(v VertexT) error {
	if g.hasVertex(v) && g.enforceUniqueVertices {
		return ErrVertexAlreadyExists
	}

	vertices, ok := g.vertices.Get(v.ID())
	if !ok {
		g.vertices.Set(v.ID(), []VertexT{})
	}

	g.vertices.Set(v.ID(), append(vertices, v))
	g.indegree[v.ID()] = 0
	g.totalVertices++
	return nil
}

// AddEdge adds a directed edge from vertex 'from' to vertex 'to'.
// Returns an error if adding the edge would create a cycle.
func (g *DAG[VertexT]) AddEdge(from, to VertexT) error {
	// Ensure both vertices exist
	if !g.hasVertex(from) || !g.hasVertex(to) {
		return ErrVertexNotFound
	}

	// Check if adding the edge would create a cycle
	if g.hasCycle(from, to) {
		return ErrAcyclicConstraintViolation
	}

	// Add the edge and update indegree of the destination vertex
	g.edges[from.ID()] = append(g.edges[from.ID()], to)
	g.indegree[to.ID()]++
	return nil
}

// TopologicalSort performs a topological sort on the DAG and
// returns a slice of vertices in topologically sorted order.
func (g *DAG[VertexT]) TopologicalSort() ([]VertexT, error) {
	// Kahn's algorithm for topological sorting
	var sorted []VertexT
	queue := []string{}

	// Enqueue vertices with zero indegree
	for _, vertex := range g.vertices.Keys() {
		if g.indegree[vertex] == 0 {
			queue = append(queue, vertex)
		}
	}

	// Process vertices in the queue
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		verticesForKey, _ := g.vertices.Get(v)
		sorted = append(sorted, verticesForKey...)

		// For each outgoing edge from 'v', reduce indegree and
		// enqueue if it becomes zero
		for _, neighbor := range g.edges[v] {
			g.indegree[neighbor.ID()]--
			if g.indegree[neighbor.ID()] == 0 {
				queue = append(queue, neighbor.ID())
			}
		}
	}

	// Check if we could process all vertices (DAG should have no cycles)
	if len(sorted) != g.totalVertices {
		return nil, ErrAcyclicConstraintViolation
	}
	return sorted, nil
}

// Helper function to check if adding an edge would create a cycle using DFS.
func (g *DAG[VertexT]) hasCycle(from, to VertexT) bool {
	visited := make(map[string]bool)
	return g.detectCycle(to, from, visited)
}

// detectCycle is a helper function which returns if there is a cycle
// via DFS.
func (g *DAG[VertexT]) detectCycle(
	v, target VertexT, visited map[string]bool,
) bool {
	if v.ID() == target.ID() {
		return true
	}
	visited[v.ID()] = true
	for _, neighbor := range g.edges[v.ID()] {
		if !visited[neighbor.ID()] {
			if g.detectCycle(neighbor, target, visited) {
				return true
			}
		}
	}
	visited[v.ID()] = false
	return false
}

// hasVertex returns whether the given vertex exists in the DAG.
func (g *DAG[VertexT]) hasVertex(v VertexT) bool {
	_, exists := g.vertices.Get(v.ID())
	return exists
}

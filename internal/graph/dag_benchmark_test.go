package graph

import (
	"strconv"
	"testing"
)

type testVertex struct {
	id string
}

func (v testVertex) ID() string {
	return v.id
}

// BenchmarkAddVertex benchmarks the AddVertex method.
func BenchmarkAddVertex(b *testing.B) {
	dag := NewDAG[testVertex]()
	numVertices := 100000
	vertices := make([]testVertex, numVertices)

	for i := 0; i < numVertices; i++ {
		vertices[i] = testVertex{id: strconv.Itoa(i)}
	}

	b.ResetTimer()
	for _, v := range vertices {
		_ = dag.AddVertex(v)
	}
}

// BenchmarkAddEdgeLinear benchmarks the AddEdge method with a prebuilt,
// linearly growing graph structure.
func BenchmarkAddEdgeLinear(b *testing.B) {
	dag := NewDAG[testVertex]()

	// pre-create a list of vertices
	numVertices := 100000
	vertices := make([]testVertex, numVertices)
	for i := 0; i < numVertices; i++ {
		vertices[i] = testVertex{id: strconv.Itoa(i)}
		dag.AddVertex(vertices[i])
	}

	// pre-create a linear edge structure
	var edges [][2]testVertex
	for i := 1; i < numVertices; i++ {
		edges = append(edges, [2]testVertex{vertices[i-1], vertices[i]})
	}

	b.ResetTimer()
	for _, edge := range edges {
		_ = dag.AddEdge(edge[0], edge[1])
	}
}

// BenchmarkAddEdgeComplex benchmarks the AddEdge method with a prebuilt,
// branching graph structure.
func BenchmarkAddEdgeComplex(b *testing.B) {
	dag := NewDAG[testVertex]()
	initialVertex := testVertex{id: "0"}
	dag.AddVertex(initialVertex)

	numVertices := 10000

	// pre-create a list of vertices and edges to avoid setup in the main benchmark loop
	vertices := make([]testVertex, numVertices)
	for i := 1; i <= numVertices; i++ {
		newVertex := testVertex{id: strconv.Itoa(i)}
		dag.AddVertex(newVertex)
		vertices = append(vertices, newVertex)
	}

	// pre-create a list of edges to add at the end
	var edges [][2]testVertex
	for i := 1; i < len(vertices); i++ {
		// linear edge
		edges = append(edges, [2]testVertex{vertices[i-1], vertices[i]})

		if i > 1 {
			// edge from two steps back
			edges = append(edges, [2]testVertex{vertices[i-2], vertices[i]})
		}
		if i > 2 && i%2 == 0 {
			// branch by adding edge every other iteration
			edges = append(edges, [2]testVertex{vertices[i-3], vertices[i]})
		}
	}

	b.ResetTimer()
	for _, edge := range edges {
		_ = dag.AddEdge(edge[0], edge[1])
	}
}

// BenchmarkTopologicalSortBasic benchmarks the TopologicalSort method with a
// basic graph.
func BenchmarkTopologicalSortBasic(b *testing.B) {
	dag := NewDAG[testVertex]()
	vertices := []testVertex{
		{id: "1"}, {id: "2"}, {id: "3"},
		{id: "4"}, {id: "5"}, {id: "6"},
	}

	// build basic graph
	for _, v := range vertices {
		dag.AddVertex(v)
	}
	dag.AddEdge(vertices[0], vertices[1])
	dag.AddEdge(vertices[1], vertices[2])
	dag.AddEdge(vertices[2], vertices[3])
	dag.AddEdge(vertices[3], vertices[4])
	dag.AddEdge(vertices[4], vertices[5])

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = dag.TopologicalSort()
	}
}

// BenchmarkTopologicalSort benchmarks the TopologicalSort method with a more
// complex graph.
func BenchmarkTopologicalSortComplex(b *testing.B) {
	dag := NewDAG[testVertex]()

	// large set of vertices
	numVertices := 100
	vertices := make([]testVertex, numVertices)
	for i := 0; i < numVertices; i++ {
		vertices[i] = testVertex{id: strconv.Itoa(i)}
		dag.AddVertex(vertices[i])
	}

	// create a graph with multiple paths and branches
	for i := 0; i < numVertices-1; i++ {
		// initial linear chain
		dag.AddEdge(vertices[i], vertices[i+1])

		// branching edges
		if i+2 < numVertices {
			dag.AddEdge(vertices[i], vertices[i+2])
		}
		if i+3 < numVertices && i%2 == 0 {
			dag.AddEdge(vertices[i], vertices[i+3])
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = dag.TopologicalSort()
	}
}

package graph_test

import (
	"strconv"
	"testing"

	"github.com/skjdfhkskjds/depinject/internal/graph"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testVertex struct {
	id string
}

func (v testVertex) ID() string {
	return v.id
}

func TestDAG(t *testing.T) {
	t.Run("NewDAG", func(t *testing.T) {
		dag := graph.NewDAG[testVertex](false)
		assert.NotNil(t, dag)
	})

	t.Run("Vertices", func(t *testing.T) {
		dag := graph.NewDAG[testVertex](false)
		assert.Empty(t, dag.Vertices())

		v1 := testVertex{id: "1"}
		v2 := testVertex{id: "2"}
		dag.AddVertex(v1)
		dag.AddVertex(v2)
		assert.Equal(t, []testVertex{v1, v2}, dag.Vertices())
	})

	t.Run("AddVertex", func(t *testing.T) {
		dag := graph.NewDAG[testVertex](true)
		v := testVertex{id: "1"}
		err := dag.AddVertex(v)
		assert.NoError(t, err)

		// Adding the same vertex again should return an error
		err = dag.AddVertex(v)
		assert.ErrorIs(t, err, graph.ErrVertexAlreadyExists)
	})

	t.Run("AddEdge", func(t *testing.T) {
		dag := graph.NewDAG[testVertex](false)
		v1 := testVertex{id: "1"}
		v2 := testVertex{id: "2"}

		err := dag.AddVertex(v1)
		assert.NoError(t, err)
		err = dag.AddVertex(v2)
		assert.NoError(t, err)

		err = dag.AddEdge(v1, v2)
		assert.NoError(t, err)

		// Adding an edge that would create a cycle should return an error
		err = dag.AddEdge(v2, v1)
		assert.ErrorIs(t, err, graph.ErrAcyclicConstraintViolation)
	})

	t.Run("TopologicalSort", func(t *testing.T) {
		dag := graph.NewDAG[testVertex](false)
		v1 := testVertex{id: "1"}
		v2 := testVertex{id: "2"}
		v3 := testVertex{id: "3"}

		require.NoError(t, dag.AddVertex(v1))
		require.NoError(t, dag.AddVertex(v2))
		require.NoError(t, dag.AddVertex(v3))

		require.NoError(t, dag.AddEdge(v1, v2))
		require.NoError(t, dag.AddEdge(v2, v3))

		sorted, err := dag.TopologicalSort()
		assert.NoError(t, err)
		assert.Equal(t, []testVertex{v1, v2, v3}, sorted)
	})
}

// ----------------------------------------------------------------------------
//                                 Benchmarks
// ----------------------------------------------------------------------------

// BenchmarkAddVertex benchmarks the AddVertex method.
func BenchmarkAddVertex(b *testing.B) {
	dag := graph.NewDAG[testVertex](false)
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
	dag := graph.NewDAG[testVertex](false)

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
	dag := graph.NewDAG[testVertex](false)
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
	dag := graph.NewDAG[testVertex](false)
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
	dag := graph.NewDAG[testVertex](false)

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

package graph_test

import (
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
		dag := graph.NewDAG[testVertex]()
		assert.NotNil(t, dag)
	})

	t.Run("AddVertex", func(t *testing.T) {
		dag := graph.NewDAG[testVertex]()
		v := testVertex{id: "1"}
		err := dag.AddVertex(v)
		assert.NoError(t, err)

		// Adding the same vertex again should return an error
		err = dag.AddVertex(v)
		assert.ErrorIs(t, err, graph.ErrVertexAlreadyExists)
	})

	t.Run("AddEdge", func(t *testing.T) {
		dag := graph.NewDAG[testVertex]()
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
		dag := graph.NewDAG[testVertex]()
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

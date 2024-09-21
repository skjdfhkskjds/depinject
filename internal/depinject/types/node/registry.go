package node

import (
	stderr "errors"

	"github.com/skjdfhkskjds/depinject/internal/depinject/types/errors"
	"github.com/skjdfhkskjds/depinject/internal/reflect"
)

var ErrDuplicateOutput = stderr.New("duplicate output")

// A Registry is a registry of nodes which stores node
// data by inferring type associations to the node's outputs
// and allows for quick lookups of nodes by their outputs.
type Registry struct {
	nodes map[reflect.Type]*Node
}

// NewRegistry creates a new NodeRegistry.
func NewRegistry() *Registry {
	return &Registry{
		nodes: make(map[reflect.Type]*Node),
	}
}

// Register adds a node to the registry.
func (r *Registry) Register(node *Node) error {
	for _, output := range node.Outputs() {
		if _, ok := r.nodes[output]; ok {
			return errors.New(ErrDuplicateOutput, output.String())
		}
		r.nodes[output] = node
	}

	return nil
}

// Get retrieves a node from the registry by its output type.
func (r *Registry) Get(t reflect.Type) (*Node, bool) {
	node, ok := r.nodes[t]
	return node, ok
}

// Nodes returns all nodes in the registry.
func (r *Registry) Nodes() []*Node {
	nodes := make([]*Node, 0, len(r.nodes))
	for _, node := range r.nodes {
		nodes = append(nodes, node)
	}
	return nodes
}
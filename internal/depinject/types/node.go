package types

import (
	"errors"

	"github.com/skjdfhkskjds/depinject/internal/graph"
	"github.com/skjdfhkskjds/depinject/internal/reflect"
)

var _ graph.Vertex[*Node] = (*Node)(nil)

// Node is a Node in the graph.
type Node struct {
	id string

	outputs     map[reflect.Type]reflect.Value
	constructor *reflect.Func
}

// NewNode creates a new node.
func NewNode(constructor any) (*Node, error) {
	fn, err := reflect.NewFunc(constructor)
	if err != nil {
		return nil, err
	}

	return &Node{
		id:          fn.Name,
		outputs:     make(map[reflect.Type]reflect.Value),
		constructor: fn,
	}, nil
}

// Dependencies returns the dependencies of the node.
func (n *Node) Dependencies() []reflect.Type {
	return n.constructor.Args
}

// Outputs returns the outputs of the node.
func (n *Node) Outputs() []reflect.Type {
	return n.constructor.Ret
}

// Execute executes the node.
func (n *Node) Execute(args ...any) error {
	values, err := n.constructor.Call(args...)
	if err != nil {
		return err
	}

	for _, value := range values {
		n.outputs[reflect.TypeOf(value)] = value
	}

	return nil
}

// ValueOf returns the value of the node for the given type.
func (n *Node) ValueOf(t reflect.Type) (reflect.Value, error) {
	if v, ok := n.outputs[t]; ok {
		return v, nil
	}

	return reflect.Value{}, errors.New("value not found")
}

// ID returns the ID of the node.
func (n *Node) ID() string {
	return n.id
}

package types

import (
	"errors"

	"github.com/skjdfhkskjds/depinject/internal/graph"
	"github.com/skjdfhkskjds/depinject/internal/reflect"
)

var _ graph.Vertex[*Node] = (*Node)(nil)

// Node is a Node in the graph.
type Node struct {
	id       string
	incoming []*Node

	outputs     map[reflect.Type]reflect.Value
	constructor *reflect.Func
}

// NewNode creates a new node.
func NewNode(constructor *reflect.Func) (*Node, error) {
	return &Node{
		outputs:     make(map[reflect.Type]reflect.Value),
		constructor: constructor,
	}, nil
}

// Init initializes the node with the given incoming nodes.
func (n *Node) Init(incoming []*Node) {
	n.incoming = append(n.incoming, incoming...)
}

// Dependencies returns the dependencies of the node.
func (n *Node) Dependencies() []reflect.Type {
	return n.constructor.Args
}

// Execute executes the node.
func (n *Node) Execute() error {
	values, err := n.constructor.Call()
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

// Incoming returns the incoming nodes of the node.
func (n *Node) Incoming() []*Node {
	return n.incoming
}

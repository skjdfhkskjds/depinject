package node

import (
	stdreflect "reflect"

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

// New creates a new node.
func New(constructor any) (*Node, error) {
	fn, err := reflect.WrapFunc(constructor)
	if err != nil {
		return nil, err
	}
	fn.Ret = filterError(fn.Ret)

	return &Node{
		id:          fn.Name,
		outputs:     make(map[reflect.Type]reflect.Value),
		constructor: fn,
	}, nil
}

// NewFromFunc creates a new node from a reflect.Func.
func NewFromFunc(fn *reflect.Func) *Node {
	fn.Ret = filterError(fn.Ret)
	return &Node{
		id:          fn.Name,
		outputs:     make(map[reflect.Type]reflect.Value),
		constructor: fn,
	}
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
		n.outputs[value.Type()] = value
	}

	return nil
}

// ValueOf returns the value of the node for the given type.
// If the type is an interface, it will attempt to infer the
// implementation from the node's outputs.
func (n *Node) ValueOf(t reflect.Type) (reflect.Value, error) {
	if v, ok := n.outputs[t]; ok {
		return v, nil
	}

	if t.Kind() == stdreflect.Interface {
		var impl reflect.Type
		for _, output := range n.Outputs() {
			if output.Implements(t) {
				if impl != nil {
					return reflect.Value{}, ErrMultipleImplementations(t)
				}
				impl = output
			}
		}
		if impl != nil {
			return n.outputs[impl], nil
		}
	}

	return reflect.Value{}, ErrValueNotFound(t)
}

// ID returns the ID of the node.
func (n *Node) ID() string {
	return n.id
}

func filterError(types []reflect.Type) []reflect.Type {
	var ret []reflect.Type
	for _, t := range types {
		if t != reflect.TypeOf((*error)(nil)).Elem() {
			ret = append(ret, t)
		}
	}
	return ret
}

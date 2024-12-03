package types

import (
	"fmt"

	"github.com/skjdfhkskjds/depinject/internal/reflect"
)

type Node struct {
	id string

	// The wrapped function.
	constructor *reflect.Func
}

func NewNode(constructor any) (*Node, error) {
	fn, err := reflect.WrapFunc(constructor)
	if err != nil {
		return nil, err
	}
	return NewNodeFromFunc(fn), nil
}

func NewNodeFromFunc(fn *reflect.Func) *Node {
	return &Node{
		id:          fn.Name,
		constructor: fn,
	}
}

// ============================================================================
//                                   Getters
// ============================================================================

func (n *Node) ID() string {
	return n.id
}

func (n *Node) Dependencies() []*reflect.Arg {
	return n.constructor.Args
}

func (n *Node) ValueOf(t reflect.Type) (reflect.Value, error) {
	value, ok := n.constructor.Ret[t]
	if !ok {
		return reflect.Value{}, fmt.Errorf("no value for type %v", t)
	}
	return value, nil
}

func (n *Node) Outputs() []reflect.Type {
	types := make([]reflect.Type, len(n.constructor.Ret))
	i := 0
	for t := range n.constructor.Ret {
		types[i] = t
		i++
	}
	return types
}

// ============================================================================
//                                    Misc
// ============================================================================

// Execute calls the constructor with the given arguments.
func (n *Node) Execute(inferInterfaces bool, args ...any) error {
	return n.constructor.Call(inferInterfaces, args...)
}

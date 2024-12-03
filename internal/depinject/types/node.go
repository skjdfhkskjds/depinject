package types

import (
	"github.com/skjdfhkskjds/depinject/internal/errors"
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

func (n *Node) ValueOf(t reflect.Type, inferInterfaces bool) (reflect.Value, error) {
	if value, ok := n.constructor.Ret[t]; ok {
		return value, nil
	}

	if !inferInterfaces {
		return reflect.Value{}, errors.Newf(noValueForTypeErrMsg, t)
	}

	// If we are inferring interfaces, we search for the first type
	// that is assignable to the requested type.
	for returnType, value := range n.constructor.Ret {
		if returnType.AssignableTo(t) {
			return value, nil
		}
	}
	return reflect.Value{}, errors.Newf(noValueForTypeErrMsg, t)
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

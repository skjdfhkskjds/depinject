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

// Returns the first value of the constructor that matches the given type.
// Checks in the order of:
//   - exact match
//   - element type of slice/array exact match (if matchElement is true)
//   - assignable match (if inferInterfaces is true)
//   - element type of slice/array assignable match (if matchElement is true)
//
// Requires:
//   - matchElement is true if and only if t is a slice or array type and we
//     are allowing list inference for an element type match.
func (n *Node) ValueOf(
	t reflect.Type, matchElement, inferInterfaces bool,
) (reflect.Value, error) {
	if value, ok := n.constructor.Ret[t]; ok && value.IsValid() {
		return value, nil
	}
	if matchElement {
		if value, ok := n.constructor.Ret[t.Elem()]; ok && value.IsValid() {
			return value, nil
		} else if ok {
			return reflect.Value{}, errors.Newf(
				noValueForTypeErrMsg, t.Elem(), n.ID(),
			)
		}
	}

	if !inferInterfaces {
		return reflect.Value{}, errors.Newf(
			noValueForTypeErrMsg, t, n.ID(),
		)
	}

	// If we are inferring interfaces, we search for the first type
	// that is assignable to the requested type.
	for returnType, value := range n.constructor.Ret {
		if returnType.AssignableTo(t) && value.IsValid() {
			return value, nil
		}
		if matchElement && returnType.AssignableTo(t.Elem()) && value.IsValid() {
			return value, nil
		} else if !value.IsValid() {
			return reflect.Value{}, errors.Newf(
				noValueForTypeErrMsg, t.Elem(), n.ID(),
			)
		}
	}
	return reflect.Value{}, errors.Newf(noValueForTypeErrMsg, t, n.ID())
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

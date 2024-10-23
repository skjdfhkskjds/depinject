package node

import (
	"fmt"

	"reflect"
)

// errMissingDependency is returned when a dependency is not found in the node.
type errMissingDependency struct {
	Type reflect.Type
}

func (e *errMissingDependency) Error() string {
	return fmt.Sprintf("missing dependency for type: %s", e.Type)
}

func ErrMissingDependency(t reflect.Type) error {
	return &errMissingDependency{Type: t}
}

// errMultipleImplementations is returned when multiple implementations
// of an interface are found.
type errMultipleImplementations struct {
	Type reflect.Type
}

func (e *errMultipleImplementations) Error() string {
	return fmt.Sprintf("multiple implementations found for type: %s", e.Type)
}

func ErrMultipleImplementations(t reflect.Type) error {
	return &errMultipleImplementations{Type: t}
}

// errValueNotFound is returned when a value does not exist in the node.
type errValueNotFound struct {
	Type reflect.Type
}

func (e *errValueNotFound) Error() string {
	return fmt.Sprintf("value not found for type: %s", e.Type)
}

func ErrValueNotFound(t reflect.Type) error {
	return &errValueNotFound{Type: t}
}

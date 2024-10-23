package depinject

import (
	"github.com/skjdfhkskjds/depinject/internal/reflect"
)

type (
	In  interface{ in() }
	Out interface{ out() }
)

// handleInputSentinel takes a type, and injects the type's
// constructor as a new node into the container if the type
// is an input sentinel, otherwise it is a no-op.
//
// input sentinels are treated as additional nodes whose
// constructors are a maximal list of the struct's fields
func (c *Container) handleInputSentinel(t reflect.Type) error {
	// If t does not implement In, skip.
	if !t.Implements(reflect.TypeOf((*In)(nil)).Elem()) {
		return nil
	}

	s, err := reflect.NewStruct(t)
	if err != nil {
		return err
	}

	return c.Provide(s.Constructor())
}

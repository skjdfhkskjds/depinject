package sentinels

import "reflect"

var InOut = &S{}

type (
	In  interface{ in() }
	Out interface{ out() }

	// internal implementation to use for container resolution
	S struct{}
)

func (*S) in()  {}
func (*S) out() {}

// IsIn returns true if the given reflect.Type implements the In interface.
func IsIn(t reflect.Type) bool {
	if t == nil {
		return false
	}
	return t.Implements(reflect.TypeOf((*In)(nil)).Elem())
}

// IsOut returns true if the given reflect.Type implements the Out interface.
func IsOut(t reflect.Type) bool {
	if t == nil {
		return false
	}
	return t.Implements(reflect.TypeOf((*Out)(nil)).Elem())
}

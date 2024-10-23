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

// ImplementsIn returns true if the given reflect.Type implements the In interface.
// It returns false if the type is nil or exactly the In interface itself.
func ImplementsIn(t reflect.Type) bool {
	if t == nil || t == reflect.TypeOf((*In)(nil)).Elem() {
		return false
	}
	return t.Implements(reflect.TypeOf((*In)(nil)).Elem())
}

// ImplementsOut returns true if the given reflect.Type implements the Out interface.
// It returns false if the type is nil or exactly the Out interface itself.
func ImplementsOut(t reflect.Type) bool {
	if t == nil || t == reflect.TypeOf((*Out)(nil)).Elem() {
		return false
	}
	return t.Implements(reflect.TypeOf((*Out)(nil)).Elem())
}

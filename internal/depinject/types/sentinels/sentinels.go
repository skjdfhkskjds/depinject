package sentinels

import "reflect"

type (
	// Note: Embedding the sentinels needs to be at the top level
	// of the struct, we do not check recursively.
	In  struct{ _ sentinel }
	Out struct{ _ sentinel }

	// internal implementation to use for container resolution
	sentinel struct{}
)

// EmbedsIn returns true if the given reflect.Type implements the In interface.
// It returns false if the type is nil or exactly the In type itself.
func EmbedsIn(t reflect.Type) bool {
	if t == nil || t == reflect.TypeOf(In{}) {
		return false
	}

	// Check if the type is a struct
	if t.Kind() != reflect.Struct {
		return false
	}

	// Iterate through all fields of the struct
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type == reflect.TypeOf(In{}) {
			return true
		}
	}

	return false
}

// EmbedsOut returns true if the given reflect.Type embeds the Out type.
// It returns false if the type is nil or exactly the Out type itself.
func EmbedsOut(t reflect.Type) bool {
	if t == nil || t == reflect.TypeOf(Out{}) {
		return false
	}

	// Check if the type is a struct
	if t.Kind() != reflect.Struct {
		return false
	}

	// Iterate through all fields of the struct
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type == reflect.TypeOf(Out{}) {
			return true
		}
	}

	return false
}

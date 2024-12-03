package depinject

import (
	"github.com/skjdfhkskjds/depinject/internal/depinject3/types"
	"github.com/skjdfhkskjds/depinject/internal/reflect"
)

type (
	// Note: Embedding the sentinels needs to be at the top level
	// of the struct, we do not check recursively.
	In  struct{ _ sentinel }
	Out struct{ _ sentinel }

	// internal implementation to use for container resolution
	sentinel struct{}
)

// embedsInSentinel returns true if any of the given node's parameters
// embeds the In type.
// Returns a list of struct types that embed the In type.
func embedsInSentinel(node *types.Node) ([]*reflect.StructType, bool) {
	var structs []*reflect.StructType
	for _, dep := range node.Dependencies() {
		if typeEmbedsSentinel(dep.Type, reflect.TypeOf(In{})) {
			s, err := reflect.NewStruct(dep.Type)
			if err != nil {
				return nil, false
			}
			structs = append(structs, s)
		}
	}

	return structs, len(structs) > 0
}

// embedsOutSentinel returns true if any of the given node's outputs
// embeds the Out type.
// Returns a list of struct types that embed the Out type.
func embedsOutSentinel(node *types.Node) ([]*reflect.StructType, bool) {
	var structs []*reflect.StructType
	for _, out := range node.Outputs() {
		if typeEmbedsSentinel(out, reflect.TypeOf(Out{})) {
			s, err := reflect.NewStruct(out)
			if err != nil {
				return nil, false
			}
			structs = append(structs, s)
		}
	}

	return structs, len(structs) > 0
}

// typeEmbedsSentinel returns true if the given type embeds the sentinel type.
// It returns false otherwise, including if the type is nil or exactly
// the sentinel type itself.
func typeEmbedsSentinel(t reflect.Type, sentinel reflect.Type) bool {
	if t == nil || t == sentinel {
		return false
	}

	// Check if the type is a struct
	if t.Kind() != reflect.Struct {
		return false
	}

	// Iterate through all fields of the struct
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type == sentinel {
			return true
		}
	}
	return false
}

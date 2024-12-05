package depinject

import (
	"github.com/skjdfhkskjds/depinject/internal/depinject/types"
	"github.com/skjdfhkskjds/depinject/internal/reflect"
	"github.com/skjdfhkskjds/depinject/internal/utils"
)

type (
	// Note: Embedding the sentinels needs to be at the top level
	// of the struct, we do not check recursively.
	In  struct{ _ sentinel }
	Out struct{ _ sentinel }

	// internal implementation to use for container resolution
	sentinel struct{}
)

// parseInSentinels parses the sentinel structs in the node's constructor
// inputs and returns a new set of nodes that execute the constructor of
// each of the sentinel structs.
func parseInSentinels(node *types.Node) ([]*types.Node, error) {
	sentinelNodes := make([]*types.Node, 0)
	// Handle the In sentinel structs
	inSentinelStructs, err := structsForSentinelByType(
		utils.MapSlice(
			node.Dependencies(),
			func(arg *reflect.Arg) reflect.Type { return arg.Type },
		),
		reflect.TypeOf(In{}),
	)
	if err != nil {
		return nil, err
	}
	for _, s := range inSentinelStructs {
		sentinelNodes = append(
			sentinelNodes,
			types.NewNodeFromFunc(s.Constructor()),
		)
	}
	return sentinelNodes, nil
}

// parseOutSentinels parses the sentinel structs in the node's constructor
// outputs and returns a new set of nodes that execute the provider of
// each of the sentinel structs.
func parseOutSentinels(node *types.Node) ([]*types.Node, error) {
	sentinelNodes := []*types.Node{}

	// Handle the Out sentinel structs
	outSentinelStructs, err := structsForSentinelByType(
		node.Outputs(),
		reflect.TypeOf(Out{}),
	)
	if err != nil {
		return nil, err
	}
	for _, s := range outSentinelStructs {
		sentinelNodes = append(
			sentinelNodes,
			types.NewNodeFromFunc(s.Provider()),
		)
	}

	return sentinelNodes, nil
}

// structsForSentinelByType returns a list of structs that embed the
// sentinel type. It also filters out the sentinel type as a field from
// the structs.
func structsForSentinelByType(
	sourceTypes []reflect.Type, sentinelType reflect.Type,
) ([]*reflect.StructType, error) {
	structs := make([]*reflect.StructType, 0)
	for _, sourceType := range sourceTypes {
		if embedsSentinel(sourceType, sentinelType) {
			s, err := reflect.NewStruct(sourceType)
			if err != nil {
				return nil, err
			}
			// Filter out the sentinel struct as a field.
			isNotSentinel := func(t reflect.Type) bool {
				return t != sentinelType
			}
			s.Fields = s.Fields.Filter(isNotSentinel)
			structs = append(structs, s)
		}
	}

	return structs, nil
}

// embedsSentinel returns true if the given type embeds the sentinel type.
// It returns false otherwise, including if the type is nil or exactly
// the sentinel type itself.
func embedsSentinel(t reflect.Type, sentinel reflect.Type) bool {
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

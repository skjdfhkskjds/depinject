package types

import (
	"strings"

	"github.com/skjdfhkskjds/depinject/internal/errors"
	"github.com/skjdfhkskjds/depinject/internal/reflect"
)

// The registry is responsible for managing the relationship
// between types and the nodes that provide them.
type Registry struct {
	// providers maps a particular argument to all the nodes
	// which provide that argument.
	providers map[reflect.Type][]*Node

	inferLists      bool
	inferInterfaces bool
}

func NewRegistry(inferLists, inferInterfaces bool) *Registry {
	return &Registry{
		providers:       make(map[reflect.Type][]*Node),
		inferLists:      inferLists,
		inferInterfaces: inferInterfaces,
	}
}

// Register registers a node in the registry.
// Contract:
//   - if inferLists is true, the registry will permit multiple
//     providers being registered for the same type.
func (r *Registry) Register(node *Node) error {
	for _, t := range node.Outputs() {
		// Skip errors, they are handled separately
		if reflect.IsError(t) {
			continue
		}
		if _, exists := r.providers[t]; exists && !r.inferLists {
			return errors.Newf(multipleProvidersErrMsg, t)
		} else if !exists {
			r.providers[t] = make([]*Node, 0)
		}
		r.providers[t] = append(r.providers[t], node)
	}
	return nil
}

// Lookup returns all the nodes which provide the given type.
// Contract:
//   - if inferInterfaces is true, this node will be registered as
//     a provider for ALL registered types which are assignable by
//     an output of this node.
func (r *Registry) Lookup(requested reflect.Type, optional bool) ([]*Node, error) {
	allProviders := make([]*Node, 0)
	for _, t := range r.allMatchingTypes(requested) {
		providers, ok := r.providers[t]
		if ok {
			allProviders = append(allProviders, providers...)
		}
	}
	if !optional && len(allProviders) == 0 {
		return nil, errors.Newf(noProvidersErrMsg, requested)
	}
	return allProviders, nil
}

func (r *Registry) Dump() string {
	var dump strings.Builder
	for t, nodes := range r.providers {
		dump.WriteString(t.String() + ":\n")
		for _, node := range nodes {
			dump.WriteString("\t" + node.ID() + "\n")
		}
	}
	return dump.String()
}

// allMatchingTypes returns all the types in the registry that are
// "matched by" the given type.
// That is, if r.inferInterfaces is true, and there exists type A in
// the registry that implements t, then [t, A] is returned.
func (r *Registry) allMatchingTypes(t reflect.Type) []reflect.Type {
	types := []reflect.Type{t}

	var internalType reflect.Type
	if r.inferLists && (t.Kind() == reflect.Slice || t.Kind() == reflect.Array) {
		internalType = t.Elem()
	}

	for existingType := range r.providers {
		if t == existingType {
			continue
		} else if internalType != nil && internalType == existingType {
			types = append(types, existingType)
		}

		// Only check assignability if we are inferring interfaces.
		if !r.inferInterfaces {
			continue
		}
		if existingType.AssignableTo(t) {
			types = append(types, existingType)
		}
		if internalType != nil && existingType.AssignableTo(internalType) {
			types = append(types, existingType)
		}
	}

	return types
}

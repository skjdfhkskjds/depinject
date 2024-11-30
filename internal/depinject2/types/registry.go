package types

import (
	"fmt"

	"github.com/skjdfhkskjds/depinject/internal/reflect"
)

// TODO: use ocean's error stuff.

const (
	// errDuplicateConstructor is the formattable error message
	// for a duplicate constructor.
	errDuplicateConstructor = "duplicate constructor: %s"

	// errNoConstructorsForType is the formattable error message
	// for when there are no constructors for a given type.
	errNoConstructorsForType = "no constructors for type: %s"

	// errMultipleConstructorsForType is the formattable error message
	// for when there are multiple constructors for a given type.
	errMultipleConstructorsForType = "multiple constructors for type: %s"

	// errMultipleImplementations is the formattable error message
	// for when there are multiple implementations for a given interface.
	errMultipleImplementations = "multiple implementations for type: %s"
)

// Registry is a collection of types and which constructors
// are responsible for creating them.
// Note: this registry does NOT enforce a 1-to-1 relationship
// between constructors and outputs, that is a single type
// can be created by multiple constructors.
type Registry struct {
	constructors map[*Constructor]struct{}
	typeRegistry map[reflect.Type][]*Constructor

	inferInterfaces bool
}

func NewRegistry(inferInterfaces bool) *Registry {
	return &Registry{
		constructors:    make(map[*Constructor]struct{}),
		typeRegistry:    make(map[reflect.Type][]*Constructor),
		inferInterfaces: inferInterfaces,
	}
}

// Register adds an entire constructor's output list to the
// registry.
// Returns an error if there already exists a constructor with
// the same ID in the registry.
func (r *Registry) Register(c *Constructor) error {
	if _, ok := r.constructors[c]; ok {
		return fmt.Errorf(errDuplicateConstructor, c.ID())
	}
	r.constructors[c] = struct{}{}

	for output := range c.Outputs() {
		if _, ok := r.typeRegistry[output]; !ok {
			r.typeRegistry[output] = make([]*Constructor, 0)
		}
		r.typeRegistry[output] = append(r.typeRegistry[output], c)
	}

	return nil
}

// Get returns all constructors that are responsible for creating
// the given type.
func (r *Registry) Get(t reflect.Type) ([]*Constructor, error) {
	if constructors, ok := r.typeRegistry[t]; ok {
		return constructors, nil
	}
	return nil, fmt.Errorf(errNoConstructorsForType, t.String())
}

func (r *Registry) GetByArg(a *reflect.Arg) ([]*Constructor, error) {
	var output []*Constructor
	for t, constructors := range r.typeRegistry {
		if a.IsType(t, r.inferInterfaces) {
			output = append(output, constructors...)
		}
	}

	if len(output) == 0 {
		return nil, fmt.Errorf(errNoConstructorsForType, a.Type.String())
	}
	return output, nil
}

// Constructors returns all constructors in the registry.
func (r *Registry) Constructors() []*Constructor {
	constructors := make([]*Constructor, 0, len(r.constructors))

	i := 0
	for c := range r.constructors {
		constructors[i] = c
		i++
	}
	return constructors
}

func (r *Registry) ValueOf(t reflect.Type) (reflect.Value, error) {
	constructors, err := r.Get(t)
	if err != nil {
		return reflect.Value{}, err
	}

	if len(constructors) > 1 {
		return reflect.Value{}, fmt.Errorf(errMultipleConstructorsForType, t.String())
	}
	c := constructors[0]

	// If the value is found directly, return it.
	if v, err := c.ValueOf(t); err == nil {
		return v, nil
	}

	// If the type is an interface, try finding the implementation
	// and return its value.
	if t.Kind() == reflect.Interface {
		var impl reflect.Type
		for output := range c.Outputs() {
			if output.Implements(t) {
				if impl != nil {
					return reflect.Value{}, fmt.Errorf(errMultipleImplementations, t.String())
				}
				impl = output
			}
		}
		if impl != nil {
			return c.outputs[impl], nil
		}
	}

	return reflect.Value{}, fmt.Errorf(errValueNotFound, t.String())
}

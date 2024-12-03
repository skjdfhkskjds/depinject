package depinject

import (
	"fmt"

	"github.com/skjdfhkskjds/depinject/internal/depinject/types"
	"github.com/skjdfhkskjds/depinject/internal/reflect"
)

const resolveErrorName = "resolve"

func (c *Container) resolve() error {
	for _, node := range c.sortedNodes {
		if err := c.resolveNode(node); err != nil {
			return newContainerError(err, resolveErrorName, node.ID())
		}
	}

	return nil
}

func (c *Container) resolveNode(node *types.Node) error {
	dependencies := node.Dependencies()
	values := make([]any, 0)
	for _, dep := range dependencies {
		// Get all the providers for each dependency.
		providers, err := c.registry.Lookup(dep.Type, dep.IsVariadic)
		if err != nil {
			return err
		}
		var value reflect.Value

		// If the dependency is an array or slice, create a slice of the
		// appropriate size and set the values from the providers.
		if c.inferLists && ((dep.IsArray && len(providers) == dep.ArraySize) || dep.IsSlice) {
			value, err = newSliceOfDep(dep, providers, c.inferInterfaces)
			if err != nil {
				return err
			}
			continue
		} else if len(providers) == 0 && dep.IsVariadic {
			continue
		} else if len(providers) != 1 {
			return fmt.Errorf("expected 1 provider, got %d", len(providers))
		}

		value, err = providers[0].ValueOf(dep.Type, c.inferInterfaces)
		if err != nil {
			return err
		}

		values = append(values, value.Interface())
	}

	if err := node.Execute(c.inferInterfaces, values...); err != nil {
		return err
	}

	return nil
}

// newSliceOfDep creates a slice of the given dependency type with the
// appropriate size and sets the values from the providers.
func newSliceOfDep(
	dep *reflect.Arg, providers []*types.Node, inferInterfaces bool,
) (reflect.Value, error) {
	slice := reflect.MakeSlice(
		dep.Type,
		len(providers),
		max(len(providers), dep.ArraySize),
	)
	for j, provider := range providers {
		providerValue, err := provider.ValueOf(dep.Type, inferInterfaces)
		if err != nil {
			return reflect.Value{}, err
		}
		slice.Index(j).Set(providerValue)
	}
	return slice, nil
}

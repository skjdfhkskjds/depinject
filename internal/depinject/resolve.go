package depinject

import (
	"github.com/skjdfhkskjds/depinject/internal/depinject/types"
	"github.com/skjdfhkskjds/depinject/internal/errors"
	"github.com/skjdfhkskjds/depinject/internal/reflect"
)

const resolveErrorName = "resolve"

// resolve resolves the container's nodes in order. By the end of this
// routine, every node will have been invoked.
func (c *Container) resolve() error {
	for _, node := range c.sortedNodes {
		if err := c.resolveNode(node); err != nil {
			return newContainerError(err, resolveErrorName, node.ID())
		}
	}

	return nil
}

// resolveNode resolves a single node.
func (c *Container) resolveNode(node *types.Node) error {
	dependencies := node.Dependencies()
	values := make([]any, 0)
	for _, dep := range dependencies {
		// Get all the providers for each dependency.
		providers, err := c.registry.Lookup(dep.Type, dep.IsVariadic)
		if err != nil {
			return err
		}

		// If the dependency is a variadic argument and there are no
		// providers, we can skip the dependency.
		if len(providers) == 0 && dep.IsVariadic {
			continue
		}

		var value reflect.Value

		// If the dependency is an array or slice, create a slice of the
		// appropriate size and set the values from the providers.
		if c.inferLists && (dep.IsArray || dep.IsSlice) {
			// Validate that the number of providers matches the expected size.
			if dep.IsArray && len(providers) != dep.ArraySize {
				return errors.Newf(
					expectedArraySizeErrMsg, dep.ArraySize, len(providers),
				)
			}

			value, err = newSliceOfDep(dep, providers, c.inferInterfaces)
			if err != nil {
				return err
			}
		} else if len(providers) != 1 {
			// If the dependency is not a list or slice and not variadic and
			// there is not exactly one provider, return an error.
			return errors.Newf(expected1ProviderErrMsg, len(providers))
		} else {
			// Otherwise, get the value from the provider. At this point, if the
			// dependency is a list or a slice, it must be provided exactly as is.
			value, err = providers[0].ValueOf(dep.Type, false, c.inferInterfaces)
			if err != nil {
				return err
			}
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
// Note: this function is only even called if c.inferLists is true.
func newSliceOfDep(
	dep *reflect.Arg, providers []*types.Node, inferInterfaces bool,
) (reflect.Value, error) {
	var values []reflect.Value
	for _, provider := range providers {
		providerValue, err := provider.ValueOf(dep.Type, true, inferInterfaces)
		if err != nil {
			return reflect.Value{}, err
		} else if dep.Type.Elem() != providerValue.Type() {
			return reflect.Value{}, errors.Newf(
				sliceElementTypesMismatchErrMsg,
				dep.Type,
				providerValue.Type(),
			)
		}
		values = append(values, providerValue)
	}
	return reflect.MakeInitializedSlice(dep.Type, values...), nil
}

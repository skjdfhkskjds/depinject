package depinject

import (
	"github.com/skjdfhkskjds/depinject/internal/depinject/types"
	"github.com/skjdfhkskjds/depinject/internal/reflect"
)

const supplyErrorName = "supply"

// Supply is a helper function that allows for the injection of
// values into the container. It is useful for injecting values
// that are not created by the container, such as command-line
// arguments or environment variables.
func (c *Container) Supply(values ...any) error {
	for _, value := range values {
		if err := c.supply(value); err != nil {
			return err
		}
	}
	return nil
}

func (c *Container) supply(value any) error {
	// Generate a function that returns the supplied value
	fn := reflect.MakeNamedFunc(
		nil, []reflect.Type{reflect.TypeOf(value)},
		func(args []reflect.Value) []reflect.Value {
			return []reflect.Value{reflect.ValueOf(value)}
		},
		reflect.TypeOf(value).String(),
	)

	if err := c.register(types.NewNodeFromFunc(fn), supplyErrorName); err != nil {
		return newContainerError(err, supplyErrorName, fn.Name)
	}
	return nil
}

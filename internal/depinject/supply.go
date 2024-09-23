package depinject

import (
	"reflect"
)

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
	// Use reflect to get the type and value of the supplied value
	valueType := reflect.TypeOf(value)

	// Generate a function that returns the supplied value
	fn := reflect.MakeFunc(
		reflect.FuncOf(nil, []reflect.Type{valueType}, false),
		func(args []reflect.Value) []reflect.Value {
			return []reflect.Value{reflect.ValueOf(value)}
		},
	)

	return c.Provide(fn.Interface())
}

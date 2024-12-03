package depinject

import (
	"github.com/skjdfhkskjds/depinject/internal/errors"
	"github.com/skjdfhkskjds/depinject/internal/reflect"
)

const invokeErrorName = "invoke"

// Invoke is a public function that allows for the invocation of
// values from the container. This function should be called after
// all required values and providers have been registered.
func (c *Container) Invoke(outputs ...any) error {
	if !c.invokable {
		var err error
		if err = c.build(); err != nil {
			return c.interceptError(err)
		}
		if err = c.resolve(); err != nil {
			return c.interceptError(err)
		}
		c.invokable = true
	}

	for _, output := range outputs {
		if err := c.invoke(output); err != nil {
			return c.interceptError(newContainerError(
				err, invokeErrorName, reflect.TypeOf(output).Elem().String(),
			))
		}
	}
	return nil
}

func (c *Container) invoke(output any) error {
	// Infer the type of the output using reflect
	outputType := reflect.TypeOf(output)

	// If the output is a pointer, get the element type
	if outputType.Kind() == reflect.Ptr {
		outputType = outputType.Elem()
	}

	// Search the registry for any value which matches the type of v
	providers, err := c.registry.Lookup(outputType, false)
	if err != nil {
		return err
	}

	// TODO: add support for array referencing on invoke.
	if len(providers) != 1 {
		return errors.Newf(expected1ProviderErrMsg, len(providers))
	}

	// Assign the value to the output
	value, err := providers[0].ValueOf(outputType, c.inferInterfaces)
	if err != nil {
		return err
	}
	reflect.ValueOf(output).Elem().Set(value)

	return nil
}

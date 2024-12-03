package depinject

import (
	"fmt"

	"github.com/skjdfhkskjds/depinject/internal/reflect"
)

const invokeErrorName = "invoke"

func (c *Container) Invoke(outputs ...any) error {
	if !c.invokable {
		var err error
		if err = c.build(); err != nil {
			return err
		}
		if err = c.resolve(); err != nil {
			return err
		}
		c.invokable = true
	}

	for _, output := range outputs {
		if err := c.invoke(output); err != nil {
			return newContainerError(
				err, invokeErrorName, reflect.TypeOf(output).Elem().String(),
			)
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
	providers, err := c.registry.Lookup(outputType)
	if err != nil {
		return err
	}

	// TODO: should we support array inferencing on invoke?
	if len(providers) != 1 {
		// fmt.Println(providers)
		return fmt.Errorf("expected 1 provider, got %d", len(providers))
	}

	// Assign the value to the output
	fmt.Println("PROVIDER")
	fmt.Println(providers[0].ID())
	fmt.Println(outputType)
	value, err := providers[0].ValueOf(outputType)
	if err != nil {
		return err
	}
	fmt.Println("VALUE")
	fmt.Println(value)
	reflect.ValueOf(output).Elem().Set(value)

	return nil
}

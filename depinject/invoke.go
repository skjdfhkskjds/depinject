package depinject

import (
	"reflect"

	"github.com/skjdfhkskjds/depinject/depinject/types"
)

// Invoke resolves the container and extracts the resulting
// values from the container.
// It returns an error if the container is invalid (not resolvable), or
// if any of the requested output types are missing from the container.
func (c *Container) Invoke(outputs ...any) error {
	var err error

	// Build the directed edges in the graph
	if err = c.build(); err != nil {
		return err
	}

	// Resolve the container
	if err = c.resolve(); err != nil {
		return err
	}

	// Invoke the outputs
	for _, output := range outputs {
		if err = c.invoke(output); err != nil {
			return err
		}
	}
	return nil
}

// invoke resolves a single output from a fully built container.
// It assumes the container is complete and valid and thus returns
// an error if the output is not found.
// TODO: do just-in-time resolution of values in case container
// is superfluous.
func (c *Container) invoke(output any) error {
	// Infer the type of the output using reflect
	outputType := reflect.TypeOf(output)

	// If the output is a pointer, get the element type
	if outputType.Kind() == reflect.Ptr {
		outputType = outputType.Elem()
	}

	// Search for the output type in the container
	node, ok := c.nodes[outputType]
	if !ok {
		return types.NewError(
			ErrMissingOutput,
			outputType.String(),
		)
	}

	// Resolve the output
	value, err := node.ValueOf(outputType)
	if err != nil {
		return types.NewError(
			err,
			node.ID(),
		)
	}

	// Assign the value to the output
	reflect.ValueOf(output).Elem().Set(value)

	return nil
}

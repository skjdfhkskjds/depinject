package depinject

import (
	"fmt"

	"github.com/skjdfhkskjds/depinject/internal/depinject/types/errors"
	"github.com/skjdfhkskjds/depinject/internal/depinject/types/node"
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
	fn, err := reflect.NewFunc(
		nil,
		[]reflect.Value{reflect.ValueOf(value)},
	)
	if err != nil {
		return errors.New(err, supplyErrorName, reflect.TypeOf(value).String())
	}

	node := node.NewFromFunc(fn)
	if err = c.addNode(node); err != nil {
		fmt.Println("NODE NAME", node.ID())
		return errors.New(err, supplyErrorName, reflect.TypeOf(value).String())
	}

	return nil
}

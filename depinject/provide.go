package depinject

import "github.com/skjdfhkskjds/depinject/depinject/types"

// Provide adds a set of constructors to the container.
// It returns an error if any of the constructors are invalid,
// or if adding them results in an invalid graph.
//
// Note: Constructors are added to the container in the order they are provided.
func (c *Container) Provide(constructors ...any) error {
	var err error
	for _, constructor := range constructors {
		if err = c.provide(constructor); err != nil {
			return err
		}
	}
	return nil
}

// provide adds a constructor to the container.
func (c *Container) provide(constructor any) error {
	node, err := types.NewNode(constructor)
	if err != nil {
		return err
	}

	// Grab the outputs of the node and track them.
	outputs := node.Outputs()
	for _, output := range outputs {
		if _, ok := c.nodes[output]; ok {
			return types.NewError(
				ErrDuplicateOutput,
				output.String(),
			)
		}
		c.nodes[output] = node
	}

	return nil
}

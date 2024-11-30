package depinject

import "github.com/skjdfhkskjds/depinject/internal/depinject2/types/sentinels"

const (
	buildErrorName   = "build"
	resolveErrorName = "resolve"
)

func (c *Container) build() error {
	if err := c.supplySentinelsIfNeeded(); err != nil {
		return err
	}

	// TODO: use reflect.Arg instead for dep resolution
	for _, constructor := range c.registry.Constructors() {
		for _, dep := range constructor.Dependencies() {
			sources, err := c.registry.GetByArg(dep)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// resolve resolves the container by iterating through every
// node in the container and executing them in a topological order.
func (c *Container) resolve() error {
	order, err := c.graph.TopologicalSort()
	if err != nil {
		return err
	}

	// TODO: use reflect.Arg instead for dep resolution
	for _, ctor := range order {
		// Get the dependencies of the ctor.
		depTypes := ctor.Dependencies()
		deps := make([]any, 0, len(depTypes))
		for i, dep := range depTypes {
			source, err := c.registry.Get(dep.Type)
			// If the last dependency is not found, and the constructor is variadic,
			// we continue without error.
			if err != nil && dep.IsVariadic && i == len(ctor.Dependencies())-1 {
				continue
			} else if err != nil {
				return newContainerError(err, resolveErrorName, ctor.ID(), dep.Type.Name())
			}

			value, err := source.ValueOf(dep.Type)
			if err != nil {
				return newContainerError(err, resolveErrorName, ctor.ID(), dep.Type.Name())
			}

			// Append the underlying casted value to deps
			deps = append(deps, value.Interface())
		}

		// Execute the constructor with the dependencies.
		if err := ctor.Execute(deps...); err != nil {
			return newContainerError(err, resolveErrorName, ctor.ID())
		}
	}

	return nil
}

// supplySentinelsIfNeeded supplies the sentinels to the container
// if they are needed as specified by the container's config.
func (c *Container) supplySentinelsIfNeeded() error {
	if c.useInSentinel {
		if err := c.supply(sentinels.In{}); err != nil {
			return err
		}
	}
	if c.useOutSentinel {
		if err := c.supply(sentinels.Out{}); err != nil {
			return err
		}
	}

	return nil
}

package depinject

import (
	"fmt"
)

var _ error = (*containerError)(nil)

// containerError is a wrapper around an error which reports on some
// error which occurred during the lifecycle of a container.
type containerError struct {
	root error

	sourceName    string
	resolvingType string
	args          []any
}

// newContainerError creates a new container error.
func newContainerError(
	root error, sourceName, resolvingType string, args ...any,
) *containerError {
	return &containerError{
		root:          root,
		sourceName:    sourceName,
		resolvingType: resolvingType,
		args:          args,
	}
}

func (e *containerError) Error() string {
	var msg = fmt.Sprintf(
		"error in %s, on %s: %s",
		e.sourceName,
		e.resolvingType,
		e.root.Error(),
	)
	if len(e.args) > 0 {
		msg += fmt.Sprintf(" (%v)", e.args...)
	}

	return msg
}

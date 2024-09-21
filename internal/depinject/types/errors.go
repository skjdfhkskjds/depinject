package types

import (
	"fmt"
)

var _ error = (*Error)(nil)

// Error is a wrapper around an error which reports on some
// error which occurred during the lifecycle of a container.
type Error struct {
	root error

	resolvingType string
	args          []any
}

// NewError creates a new error.
func NewError(root error, resolvingType string, args ...any) *Error {
	return &Error{
		root:          root,
		resolvingType: resolvingType,
		args:          args,
	}
}

func (e *Error) Error() string {
	var msg = fmt.Sprintf(
		"error while resolving %s: %s",
		e.resolvingType,
		e.root.Error(),
	)
	if len(e.args) > 0 {
		msg += fmt.Sprintf(" (%v)", e.args)
	}

	return msg
}

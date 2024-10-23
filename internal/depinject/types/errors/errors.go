package errors

import (
	"fmt"
)

var _ error = (*Error)(nil)

// Error is a wrapper around an error which reports on some
// error which occurred during the lifecycle of a container.
type Error struct {
	root error

	sourceName    string
	resolvingType string
	args          []any
}

// New creates a new error.
func New(
	root error, sourceName, resolvingType string, args ...any,
) *Error {
	return &Error{
		root:          root,
		sourceName:    sourceName,
		resolvingType: resolvingType,
		args:          args,
	}
}

func (e *Error) Error() string {
	var msg = fmt.Sprintf(
		"error in %s: failed to resolve %s: %s",
		e.sourceName,
		e.resolvingType,
		e.root.Error(),
	)
	if len(e.args) > 0 {
		msg += fmt.Sprintf(" (%v)", e.args...)
	}

	return msg
}

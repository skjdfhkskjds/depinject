package depinject

import (
	"fmt"
)

const (
	// expected1ProviderErrMsg is the error message for when the
	// expected number of providers does not match the actual number
	// of providers.
	expected1ProviderErrMsg = "expected 1 provider, got %d"

	// expectedArraySizeErrMsg is the error message for when the expected
	// array size does not match the actual array size.
	expectedArraySizeErrMsg = "expected array size %d, got %d"

	// sliceElementTypesMismatchErrMsg is the error message for when the
	// element types of a slice do not match the expected type.
	sliceElementTypesMismatchErrMsg = "slice element types mismatch: %s != %s"
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
		"Error in %s: \n\t on %s \n\t got: %s \n\t\t",
		e.sourceName,
		e.resolvingType,
		e.root.Error(),
	)
	if len(e.args) > 0 {
		msg += fmt.Sprintf("(%v)", e.args...)
	}

	return msg
}

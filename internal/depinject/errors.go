package depinject

import "errors"

var (
	// ErrMissingOutput is returned when an output is not found in the container.
	ErrMissingOutput = errors.New("missing output")

	// ErrDuplicateOutput is returned when trying to add a type which already exists in the container.
	ErrDuplicateOutput = errors.New("duplicate output")
)

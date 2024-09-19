package reflect

import "errors"

var (
	// ErrNotAFunction is returned when the type is not a function.
	ErrNotAFunction = errors.New("type is not a function")

	// ErrWrongNumArgs is returned when the number of arguments is wrong.
	ErrWrongNumArgs = errors.New("wrong number of arguments")
)

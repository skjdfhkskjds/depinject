package reflect

import "errors"

var (
	// ErrNotAFunction is returned when the type is not a function.
	ErrNotAFunction = errors.New("type is not a function")

	// ErrNotAStruct is returned when the type is not a struct.
	ErrNotAStruct = errors.New("type is not a struct")

	// ErrWrongNumArgs is returned when the number of arguments is wrong.
	ErrWrongNumArgs = errors.New("wrong number of arguments")

	// ErrInvalidArgType is returned when an argument type is invalid.
	ErrInvalidArgType = errors.New("invalid argument type")
)

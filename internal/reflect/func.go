package reflect

import (
	"errors"
	"reflect"
	"runtime"
)

// A Func is a wrapper around a reflect function value that
// provides convenience functions to get metadata and execute
// a function.
type Func struct {
	// Name is the name of the function.
	// It is formatted as "package.functionName".
	Name string

	// Args is the argument types of the function.
	Args []Type

	// Ret is the return types of the function.
	Ret []Type

	// fn is the executable Value of the function.
	fn Value
}

// NewFunc creates a new Func instance from the given function.
// It returns an error if the TypeOf(f) is not a function.
func NewFunc(f any) (*Func, error) {
	if f == nil {
		return nil, ErrNotAFunction
	}

	// Check if f is a function
	funcType := TypeOf(f)

	// Check if funcType is not nil and its kind is Func
	if funcType == nil || funcType.Kind() != reflect.Func {
		return nil, errors.Join(ErrNotAFunction, errors.New(funcType.String()))
	}

	// Create a new Func instance
	fn := &Func{
		Name: runtime.FuncForPC(ValueOf(f).Pointer()).Name(),
		Args: make([]Type, funcType.NumIn()),
		Ret:  make([]Type, funcType.NumOut()),
		fn:   ValueOf(f),
	}

	// Extract argument types
	for i := 0; i < funcType.NumIn(); i++ {
		fn.Args[i] = funcType.In(i)
	}

	// Extract return value types
	for i := 0; i < funcType.NumOut(); i++ {
		fn.Ret[i] = funcType.Out(i)
	}

	return fn, nil
}

// Call calls the original function with the given arguments.
func (f *Func) Call(args ...any) ([]Value, error) {
	if len(args) != len(f.Args) {
		return nil, ErrWrongNumArgs
	}

	// Get the arguments as Values
	in := make([]Value, len(args))
	for i, arg := range args {
		in[i] = ValueOf(arg)
	}

	// Call the function
	res := f.fn.Call(in)

	// Check if the last return value is an error
	// TODO: should we make this assumption? this is in
	// accordance to best practices in Go but not guaranteed
	if len(res) == 0 {
		return nil, nil
	}
	lastReturnValue := res[len(res)-1]
	if lastReturnValue.Type().Implements(TypeOf((*error)(nil)).Elem()) {
		if !lastReturnValue.IsNil() {
			return nil, lastReturnValue.Interface().(error)
		}
	}

	return res, nil
}

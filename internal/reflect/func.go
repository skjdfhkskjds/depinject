package reflect

import (
	"reflect"
)

// A Func is a wrapper around a reflect function value that
// provides convenience functions to get metadata and execute
// a function.
type Func struct {
	Name string

	Args []reflect.Type
	Ret  []reflect.Type

	fn reflect.Value
}

// NewFunc creates a new Func instance from the given function.
// It returns an error if the reflect.TypeOf(f) is not a function.
func NewFunc(f any) (*Func, error) {
	// Check if f is a function
	funcType := reflect.TypeOf(f)

	// Check if funcType is not nil and its kind is reflect.Func
	// TODO: make this error more specific
	if funcType == nil || funcType.Kind() != reflect.Func {
		return nil, ErrNotAFunction
	}

	// Create a new Func instance
	fn := &Func{
		Name: funcType.Name(),
		Args: make([]reflect.Type, funcType.NumIn()),
		Ret:  make([]reflect.Type, funcType.NumOut()),
		fn:   reflect.ValueOf(f),
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
func (f *Func) Call(args ...any) ([]reflect.Value, error) {
	if len(args) != len(f.Args) {
		return nil, ErrWrongNumArgs
	}

	// Get the arguments as reflect.Values
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
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
	if lastReturnValue.Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) {
		if !lastReturnValue.IsNil() {
			return nil, lastReturnValue.Interface().(error)
		}
	}

	return res, nil
}

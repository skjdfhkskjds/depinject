package reflect

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/skjdfhkskjds/depinject/internal/errors"
)

const (
	generatedFuncNamePrefix     = "GeneratedFunc"
	generatedFuncNameArgsPrefix = "Args"
	generatedFuncNameRetPrefix  = "Returns"
)

// A Func is a wrapper around a reflect function value that
// provides convenience functions to get metadata and execute
// a function.
type Func struct {
	// Name is the name of the function.
	// It is formatted as "package.functionName".
	Name string

	// Args is the argument types of the function.
	Args []*Arg

	// Ret is a mapping of return types to values of the function.
	Ret map[Type]Value

	// IsVariadic is true if the function is variadic.
	IsVariadic bool

	// HasError is true if the function returns an error.
	HasError bool

	// fn is the executable Value of the function.
	fn Value
}

// MakeNamedFunc creates a new Func instance from the given argument and return
// values.
// It generates a function which when called consumes the specified args
// and returns the given return values. It assigns this function a name
// which is formatted as "GeneratedFuncArgs{argTypes...}Ret{retTypes...}".
func MakeNamedFunc(args []Type, ret []Type, fn func([]Value) []Value, prefix string) *Func {
	generatedFn := reflect.MakeFunc(reflect.FuncOf(args, ret, false), fn).Interface()
	wrappedFunc, _ := WrapFunc(generatedFn)

	// Generate a name for the function.
	name := fmt.Sprintf("%s%s(%s%s)",
		prefix,
		generatedFuncNamePrefix,
		formatList(generatedFuncNameArgsPrefix, args),
		formatList(generatedFuncNameRetPrefix, ret),
	)
	wrappedFunc.Name = name

	return wrappedFunc
}

// WrapFunc wraps an existing go function into a Func instance.
// It returns an error if the reflect.TypeOf(f) is not a function.
func WrapFunc(f any) (*Func, error) {
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
		Name:       GetFunctionName(f),
		Args:       make([]*Arg, funcType.NumIn()),
		Ret:        make(map[Type]Value, funcType.NumOut()),
		IsVariadic: funcType.IsVariadic(),
		fn:         ValueOf(f),
	}

	// Extract argument types
	for i := 0; i < funcType.NumIn(); i++ {
		fn.Args[i] = NewArg(funcType.In(i), fn.argIsVariadic(i))
	}

	hasError := false
	// Extract return value types
	for i := 0; i < funcType.NumOut(); i++ {
		if IsError(funcType.Out(i)) {
			hasError = true
		}
		fn.Ret[funcType.Out(i)] = Value{}
	}
	fn.HasError = hasError

	return fn, nil
}

// Call calls the original function with the given arguments.
func (f *Func) Call(inferInterfaces bool, args ...any) error {
	in, err := buildAndValidateCallArgs(
		args, f.Args, f.IsVariadic, inferInterfaces,
	)
	if err != nil {
		return err
	}

	// Call the function
	var res []Value
	// If the function is variadic and the variadic argument is not
	// empty, call the function with the variadic argument as a slice.
	if f.IsVariadic && len(in) >= len(f.Args) {
		res = f.fn.CallSlice(in)
	} else {
		res = f.fn.Call(in)
	}
	if len(res) == 0 {
		return nil
	}

	if f.HasError && !res[len(res)-1].IsNil() {
		return res[len(res)-1].Interface().(error)
	}

	// Set the return values in the Func
	for _, value := range res {
		f.Ret[value.Type()] = value
	}

	return nil
}

// GetFunctionName returns the name of the function.
func GetFunctionName(f any) string {
	// Check if f is a function
	funcType := TypeOf(f)

	// Check if funcType is not nil and its kind is Func
	if funcType == nil || funcType.Kind() != reflect.Func {
		return ""
	}

	return runtime.FuncForPC(ValueOf(f).Pointer()).Name()
}

// buildAndValidateCallArgs validates the arguments against the expected types
// and returns a list of built arguments.
func buildAndValidateCallArgs(
	args []any, expected []*Arg, isVariadic, inferInterfaces bool,
) ([]Value, error) {
	lastIndex := len(expected) - 1

	// If the number of arguments is not equal to the number of expected
	// arguments, and the function is not variadic, or the function is
	// variadic but the number of arguments is less than the number of
	// expected arguments minus one, return an error.
	if len(args) != len(expected) && !(isVariadic && len(args) >= lastIndex) {
		return nil, ErrWrongNumArgs
	}

	// Build the arguments as Values
	callArgValues := make([]Value, len(args))

	// Check argument type matching.
	for i, e := range expected {
		// If we are checking the variadic argument and the remaining
		// number of arguments is not 1, we need to check the special
		// cases n = 0 and n > 1.
		if (isVariadic && i == lastIndex) && (len(args)-i != 1) {
			// If the function is variadic, the last argument
			// is a slice of the remaining arguments.
			// Variadic case with no inputs is valid.
			if len(args) == i {
				continue
			}

			// Check the argument types from the remaining arguments.
			for _, arg := range args[i:] {
				if !e.IsType(TypeOf(arg), inferInterfaces) {
					return nil, errors.Newf(
						"%w: got %s, expected %s",
						ErrInvalidArgType,
						TypeOf(arg).String(),
						e.String(),
					)
				}
			}
		}

		argValue := ValueOf(args[i])
		if !argValue.IsValid() {
			return nil, errors.Newf(ArgValueIsZeroErrMsg, e.String())
		}

		// Check if the argument matches the expected type.
		if !e.IsType(TypeOf(args[i]), inferInterfaces) {
			return nil, errors.Newf(
				"%w: got %s, expected %s",
				ErrInvalidArgType,
				TypeOf(args[i]).String(),
				e.String(),
			)
		}

		// If the argument matches the expected type, add the argument
		// as a Value to the list of arguments.
		callArgValues[i] = argValue
	}

	return callArgValues, nil
}

// argIsVariadic returns whether the argument at the given index is variadic.
func (f *Func) argIsVariadic(index int) bool {
	return f.IsVariadic && index == len(f.Args)-1
}

// formatList formats a list of types as a string.
func formatList(prefix string, list []reflect.Type) string {
	if len(list) == 0 {
		return ""
	}

	types := make([]string, len(list))
	for i, t := range list {
		types[i] = t.String()
	}
	return prefix + "{" + strings.Join(types, ", ") + "}"
}

package reflect_test

import (
	"errors"
	"testing"

	"github.com/skjdfhkskjds/depinject/internal/reflect"
	"github.com/skjdfhkskjds/depinject/internal/testutils"
)

const pkgPath = "github.com/skjdfhkskjds/depinject/internal/reflect_test."

// Mock functions for testing
func add1(a int) int {
	return a + 1
}

func addMulti(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

var errDivisionByZero = errors.New("division by zero")

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errDivisionByZero
	}
	return a / b, nil
}

// TestMakeNamedFunc tests the creation of a new Func instance using MakeNamedFunc.
func TestMakeNamedFunc(t *testing.T) {
	tests := []struct {
		name         string
		args         []reflect.Type
		ret          []reflect.Type
		fn           func([]reflect.Value) []reflect.Value
		wantHasError bool
		wantNumIn    int
		wantNumOut   int
		wantInTypes  []*reflect.Arg
		wantOutTypes map[reflect.Type]reflect.Value
		wantName     string
	}{
		{
			name: "valid function with one input and one output",
			args: []reflect.Type{reflect.TypeOf(0)},
			ret:  []reflect.Type{reflect.TypeOf(1)},
			fn: func(args []reflect.Value) []reflect.Value {
				return []reflect.Value{reflect.ValueOf(1)}
			},
			wantHasError: false,
			wantNumIn:    1,
			wantNumOut:   1,
			wantInTypes:  []*reflect.Arg{reflect.NewArg(reflect.TypeOf(0), false)},
			wantOutTypes: map[reflect.Type]reflect.Value{
				reflect.TypeOf(1): {},
			},
			wantName: "GeneratedFunc(Args{int}Returns{int})",
		},
		{
			name: "valid function with two inputs and two outputs",
			args: []reflect.Type{reflect.TypeOf(0), reflect.TypeOf(0)},
			ret:  []reflect.Type{reflect.TypeOf(1), reflect.TypeOf(errors.New(""))},
			fn: func(args []reflect.Value) []reflect.Value {
				return []reflect.Value{
					reflect.ValueOf(1),
					reflect.ValueOf(errors.New("")),
				}
			},
			wantHasError: true,
			wantNumIn:    2,
			wantNumOut:   2,
			wantInTypes: []*reflect.Arg{
				reflect.NewArg(reflect.TypeOf(0), false),
				reflect.NewArg(reflect.TypeOf(0), false),
			},
			wantOutTypes: map[reflect.Type]reflect.Value{
				reflect.TypeOf(0):              {},
				reflect.TypeOf(errors.New("")): {},
			},
			wantName: "GeneratedFunc(Args{int, int}Returns{int, *errors.errorString})",
		},
		{
			name:         "no inputs and no outputs",
			args:         []reflect.Type{},
			ret:          []reflect.Type{},
			fn:           nil,
			wantHasError: false,
			wantNumIn:    0,
			wantNumOut:   0,
			wantInTypes:  []*reflect.Arg{},
			wantOutTypes: map[reflect.Type]reflect.Value{},
			wantName:     "GeneratedFunc(Args{}Returns{})",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := reflect.MakeNamedFunc(tt.args, tt.ret, tt.fn, "")
			testutils.RequireNotNil(t, fn)

			testutils.RequireEquals(t, tt.wantNumIn, len(fn.Args))
			testutils.RequireEquals(t, tt.wantNumOut, len(fn.Ret))
			testutils.RequireEquals(t, tt.wantInTypes, fn.Args)
			testutils.RequireEquals(t, tt.wantOutTypes, fn.Ret)
			testutils.RequireEquals(t, fn.Name, tt.wantName)
			testutils.RequireEquals(t, fn.HasError, tt.wantHasError)
		})
	}
}

// TestWrapFunc tests the creation of a new Func instance.
func TestWrapFunc(t *testing.T) {
	tests := []struct {
		name           string
		input          any
		err            error
		wantHasError   bool
		wantNumIn      int
		wantNumOut     int
		wantInTypes    []*reflect.Arg
		wantOutTypes   map[reflect.Type]reflect.Value
		wantName       string
		wantIsVariadic bool
	}{
		{
			name:         "valid function with one input and one output",
			input:        add1,
			err:          nil,
			wantHasError: false,
			wantNumIn:    1,
			wantNumOut:   1,
			wantInTypes:  []*reflect.Arg{reflect.NewArg(reflect.TypeOf(0), false)},
			wantOutTypes: map[reflect.Type]reflect.Value{
				reflect.TypeOf(0): {},
			},
			wantName:       pkgPath + "add1",
			wantIsVariadic: false,
		},
		{
			name:         "valid function with two inputs and two outputs",
			input:        divide,
			err:          nil,
			wantHasError: true,
			wantNumIn:    2,
			wantNumOut:   2,
			wantInTypes: []*reflect.Arg{
				reflect.NewArg(reflect.TypeOf(0), false),
				reflect.NewArg(reflect.TypeOf(0), false),
			},
			wantOutTypes: map[reflect.Type]reflect.Value{
				reflect.TypeOf(0):                    {},
				reflect.TypeOf((*error)(nil)).Elem(): {},
			},
			wantName:       pkgPath + "divide",
			wantIsVariadic: false,
		},
		{
			name:           "nil input",
			input:          nil,
			err:            reflect.ErrNotAFunction,
			wantHasError:   true,
			wantNumIn:      0,
			wantNumOut:     0,
			wantInTypes:    nil,
			wantOutTypes:   nil,
			wantName:       "",
			wantIsVariadic: false,
		},
		{
			name:           "non-function input",
			input:          42,
			err:            reflect.ErrNotAFunction,
			wantHasError:   true,
			wantNumIn:      0,
			wantNumOut:     0,
			wantInTypes:    nil,
			wantOutTypes:   nil,
			wantName:       "",
			wantIsVariadic: false,
		},
		{
			name:         "variadic function",
			input:        addMulti,
			err:          nil,
			wantHasError: false,
			wantNumIn:    1,
			wantNumOut:   1,
			wantInTypes:  []*reflect.Arg{reflect.NewArg(reflect.TypeOf([]int{}), true)},
			wantOutTypes: map[reflect.Type]reflect.Value{
				reflect.TypeOf(0): {},
			},
			wantName:       pkgPath + "addMulti",
			wantIsVariadic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn, err := reflect.WrapFunc(tt.input)
			testutils.RequireErrorIs(t, err, tt.err)
			if err != nil {
				return
			}

			testutils.RequireEquals(t, tt.wantName, fn.Name)
			testutils.RequireEquals(t, tt.wantHasError, fn.HasError)
			testutils.RequireEquals(t, len(fn.Args), tt.wantNumIn)
			testutils.RequireEquals(t, len(fn.Ret), tt.wantNumOut)
			for i, arg := range fn.Args {
				testutils.RequireEquals(t, tt.wantInTypes[i], arg)
			}
			for outputType, ret := range fn.Ret {
				testutils.RequireEquals(t, tt.wantOutTypes[outputType], ret)
			}
		})
	}
}

// TestFunc_Call tests the Call method of the Func type.
func TestFunc_Call(t *testing.T) {
	addFn, _ := reflect.WrapFunc(add1)
	divideFn, _ := reflect.WrapFunc(divide)
	tests := []struct {
		name   string
		f      *reflect.Func
		args   []any
		output []any
		err    error
	}{
		{
			name:   "valid arguments",
			f:      addFn,
			args:   []any{5},
			output: []any{6},
			err:    nil,
		},
		{
			name:   "incorrect number of arguments",
			f:      addFn,
			args:   []any{5, 6},
			output: nil,
			err:    reflect.ErrWrongNumArgs,
		},
		{
			name:   "function execution error",
			f:      divideFn,
			args:   []any{5, 0},
			output: nil,
			err:    errDivisionByZero,
		},
		{
			name:   "valid division",
			f:      divideFn,
			args:   []any{10, 2},
			output: []any{5, (error)(nil)},
			err:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutils.RequireErrorIs(t, tt.f.Call(true, tt.args...), tt.err)

			// Only check the output if the function has no error.
			if tt.err == nil {
				got := tt.f.Ret
				testutils.RequireEquals(t, len(got), len(tt.output))

				i := 0
				for _, v := range got {
					testutils.RequireEquals(t, tt.output[i], v.Interface())
					i++
				}
			}
		})
	}
}

package reflect_test

import (
	"errors"
	"fmt"
	stdreflect "reflect"
	"testing"

	"github.com/skjdfhkskjds/depinject/internal/reflect"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const pkgPath = "github.com/skjdfhkskjds/depinject/internal/reflect_test."

// Mock functions for testing
func add1(a int) int {
	return a + 1
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
		wantNumIn    int
		wantNumOut   int
		wantInTypes  []reflect.Type
		wantOutTypes []reflect.Type
		wantName     string
	}{
		{
			name: "valid function with one input and one output",
			args: []reflect.Type{reflect.TypeOf(0)},
			ret:  []reflect.Type{reflect.TypeOf(1)},
			fn: func(args []reflect.Value) []reflect.Value {
				return []reflect.Value{reflect.ValueOf(1)}
			},
			wantNumIn:    1,
			wantNumOut:   1,
			wantInTypes:  []reflect.Type{reflect.TypeOf(0)},
			wantOutTypes: []reflect.Type{reflect.TypeOf(1)},
			wantName:     "GeneratedFunc(Args{int}Returns{int})",
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
			wantNumIn:   2,
			wantNumOut:  2,
			wantInTypes: []reflect.Type{reflect.TypeOf(0), reflect.TypeOf(0)},
			wantOutTypes: []reflect.Type{
				reflect.TypeOf(0),
				reflect.TypeOf(errors.New("")),
			},
			wantName: "GeneratedFunc(Args{int, int}Returns{int, *errors.errorString})",
		},
		{
			name:         "no inputs and no outputs",
			args:         []reflect.Type{},
			ret:          []reflect.Type{},
			fn:           nil,
			wantNumIn:    0,
			wantNumOut:   0,
			wantInTypes:  []reflect.Type{},
			wantOutTypes: []reflect.Type{},
			wantName:     "GeneratedFunc(Args{}Returns{})",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := reflect.MakeNamedFunc(tt.args, tt.ret, tt.fn)
			assert.NotNil(t, fn)

			assert.Equal(t, tt.wantNumIn, len(fn.Args))
			assert.Equal(t, tt.wantNumOut, len(fn.Ret))
			assert.Equal(t, tt.wantInTypes, fn.Args)
			assert.Equal(t, tt.wantOutTypes, fn.Ret)
			fmt.Println("fn.Name", fn.Name)
			assert.Equal(t, fn.Name, tt.wantName)
		})
	}
}

// TestWrapFunc tests the creation of a new Func instance.
func TestWrapFunc(t *testing.T) {
	tests := []struct {
		name         string
		input        any
		err          error
		wantNumIn    int
		wantNumOut   int
		wantInTypes  []reflect.Type
		wantOutTypes []reflect.Type
		wantName     string
	}{
		{
			name:         "valid function with one input and one output",
			input:        add1,
			err:          nil,
			wantNumIn:    1,
			wantNumOut:   1,
			wantInTypes:  []reflect.Type{reflect.TypeOf(0)},
			wantOutTypes: []reflect.Type{reflect.TypeOf(0)},
			wantName:     pkgPath + "add1",
		},
		{
			name:       "valid function with two inputs and two outputs",
			input:      divide,
			err:        nil,
			wantNumIn:  2,
			wantNumOut: 2,
			wantInTypes: []reflect.Type{
				reflect.TypeOf(0),
				reflect.TypeOf(0),
			},
			wantOutTypes: []reflect.Type{
				reflect.TypeOf(0),
				reflect.TypeOf((*error)(nil)).Elem(),
			},
			wantName: pkgPath + "divide",
		},
		{
			name:         "nil input",
			input:        nil,
			err:          reflect.ErrNotAFunction,
			wantNumIn:    0,
			wantNumOut:   0,
			wantInTypes:  nil,
			wantOutTypes: nil,
			wantName:     "",
		},
		{
			name:         "non-function input",
			input:        42,
			err:          reflect.ErrNotAFunction,
			wantNumIn:    0,
			wantNumOut:   0,
			wantInTypes:  nil,
			wantOutTypes: nil,
			wantName:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn, err := reflect.WrapFunc(tt.input)
			require.ErrorIs(
				t, err, tt.err,
				"NewFunc() error = %v, wantErr %v", err, tt.err,
			)
			if err != nil {
				return
			}
			require.Equal(
				t, tt.wantName, fn.Name,
				"NewFunc() function name = %s, want %s", fn.Name, tt.wantName,
			)
			require.Equal(
				t, len(fn.Args), tt.wantNumIn,
				"NewFunc() number of inputs = %d, want %d", len(fn.Args), tt.wantNumIn,
			)
			require.Equal(
				t, len(fn.Ret), tt.wantNumOut,
				"NewFunc() number of outputs = %d, want %d", len(fn.Ret), tt.wantNumOut,
			)
			for i, arg := range fn.Args {
				require.Equal(
					t, tt.wantInTypes[i], arg,
					"Input type mismatch at index %d", i,
				)
			}
			for i, ret := range fn.Ret {
				require.Equal(
					t, tt.wantOutTypes[i], ret,
					"Output type mismatch at index %d", i,
				)
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
			output: []any{5, nil},
			err:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.Call(tt.args...)
			require.ErrorIs(
				t, err, tt.err,
				"Func.Call() error = %v, wantErr %v", err, tt.err,
			)
			require.Equal(
				t, len(got), len(tt.output),
				"Func.Call() got %d return values, want %d", len(got), len(tt.output),
			)
			for i, v := range tt.output {
				if !stdreflect.DeepEqual(got[i].Interface(), v) {
					t.Errorf("Func.Call() got[%d] = %v, want %v", i, got[i].Interface(), v)
				}
			}
		})
	}
}

package reflect_test

import (
	"errors"
	stdreflect "reflect"
	"testing"

	"github.com/skjdfhkskjds/depinject/internal/reflect"
	"github.com/stretchr/testify/assert"
)

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

// TestNewFunc tests the creation of a new Func instance.
func TestNewFunc(t *testing.T) {
	tests := []struct {
		name       string
		input      any
		err        error
		wantNumIn  int
		wantNumOut int
	}{
		{
			name:       "valid function with one input and one output",
			input:      add1,
			err:        nil,
			wantNumIn:  1,
			wantNumOut: 1,
		},
		{
			name:       "valid function with two inputs and two outputs",
			input:      divide,
			err:        nil,
			wantNumIn:  2,
			wantNumOut: 2,
		},
		{
			name:       "nil input",
			input:      nil,
			err:        reflect.ErrNotAFunction,
			wantNumIn:  0,
			wantNumOut: 0,
		},
		{
			name:       "non-function input",
			input:      42,
			err:        reflect.ErrNotAFunction,
			wantNumIn:  0,
			wantNumOut: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn, err := reflect.NewFunc(tt.input)
			assert.ErrorIs(
				t, err, tt.err,
				"NewFunc() error = %v, wantErr %v", err, tt.err,
			)
			if err != nil {
				return
			}
			assert.Equal(
				t, len(fn.Args), tt.wantNumIn,
				"NewFunc() number of inputs = %d, want %d", len(fn.Args), tt.wantNumIn,
			)
			assert.Equal(
				t, len(fn.Ret), tt.wantNumOut,
				"NewFunc() number of outputs = %d, want %d", len(fn.Ret), tt.wantNumOut,
			)
		})
	}
}

// TestFunc_Call tests the Call method of the Func type.
func TestFunc_Call(t *testing.T) {
	addFn, _ := reflect.NewFunc(add1)
	divideFn, _ := reflect.NewFunc(divide)
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
			assert.ErrorIs(
				t, err, tt.err,
				"Func.Call() error = %v, wantErr %v", err, tt.err,
			)
			assert.Equal(
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

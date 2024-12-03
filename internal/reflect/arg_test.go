package reflect_test

import (
	"testing"

	"github.com/skjdfhkskjds/depinject/internal/reflect"
)

func TestNewArg(t *testing.T) {
	tests := []struct {
		name            string
		input           reflect.Type
		wantIsArray     bool
		wantIsSlice     bool
		wantIsPointer   bool
		wantIsInterface bool
		wantArraySize   int
	}{
		{
			name:            "fixed length array",
			input:           reflect.TypeOf([3]int{}),
			wantIsArray:     true,
			wantArraySize:   3,
			wantIsSlice:     false,
			wantIsPointer:   false,
			wantIsInterface: false,
		},
		{
			name:            "slice",
			input:           reflect.TypeOf([]string{}),
			wantIsArray:     false,
			wantArraySize:   0,
			wantIsSlice:     true,
			wantIsPointer:   false,
			wantIsInterface: false,
		},
		{
			name:            "pointer",
			input:           reflect.TypeOf(&struct{}{}),
			wantIsArray:     false,
			wantArraySize:   0,
			wantIsSlice:     false,
			wantIsPointer:   true,
			wantIsInterface: false,
		},
		{
			name:            "interface",
			input:           reflect.TypeOf(func(interface{}) {}).In(0),
			wantIsArray:     false,
			wantArraySize:   0,
			wantIsSlice:     false,
			wantIsPointer:   false,
			wantIsInterface: true,
		},
		{
			name:            "standard type",
			input:           reflect.TypeOf(42),
			wantIsArray:     false,
			wantArraySize:   0,
			wantIsSlice:     false,
			wantIsPointer:   false,
			wantIsInterface: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arg := reflect.NewArg(tt.input, false)

			if arg.IsArray != tt.wantIsArray {
				t.Errorf("IsArray = %v, want %v", arg.IsArray, tt.wantIsArray)
			}
			if arg.IsSlice != tt.wantIsSlice {
				t.Errorf("IsSlice = %v, want %v", arg.IsSlice, tt.wantIsSlice)
			}
			if arg.IsInterface != tt.wantIsInterface {
				t.Errorf("IsInterface = %v, want %v", arg.IsInterface, tt.wantIsInterface)
			}
			if arg.IsPointer != tt.wantIsPointer {
				t.Errorf("IsPointer = %v, want %v", arg.IsPointer, tt.wantIsPointer)
			}
			if arg.ArraySize != tt.wantArraySize {
				t.Errorf("ArraySize = %v, want %v", arg.ArraySize, tt.wantArraySize)
			}
		})
	}
}

package reflect_test

import (
	"testing"

	"github.com/skjdfhkskjds/depinject/internal/reflect"
	"github.com/skjdfhkskjds/depinject/internal/testutils"
)

func TestNewArg(t *testing.T) {
	tests := []struct {
		name          string
		input         reflect.Type
		wantIsArray   bool
		wantArraySize int
		wantIsSlice   bool
	}{
		{
			name:          "fixed length array",
			input:         reflect.TypeOf([3]int{}),
			wantIsArray:   true,
			wantArraySize: 3,
			wantIsSlice:   false,
		},
		{
			name:          "slice",
			input:         reflect.TypeOf([]string{}),
			wantIsArray:   false,
			wantArraySize: 0,
			wantIsSlice:   true,
		},
		{
			name:          "pointer",
			input:         reflect.TypeOf(&struct{}{}),
			wantIsArray:   false,
			wantArraySize: 0,
			wantIsSlice:   false,
		},
		{
			name:          "interface",
			input:         reflect.TypeOf(func(interface{}) {}).In(0),
			wantIsArray:   false,
			wantArraySize: 0,
			wantIsSlice:   false,
		},
		{
			name:          "standard type",
			input:         reflect.TypeOf(42),
			wantIsArray:   false,
			wantArraySize: 0,
			wantIsSlice:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arg := reflect.NewArg(tt.input, false)
			testutils.RequireEquals(t, arg.IsArray, tt.wantIsArray)
			testutils.RequireEquals(t, arg.ArraySize, tt.wantArraySize)
			testutils.RequireEquals(t, arg.IsSlice, tt.wantIsSlice)
		})
	}
}

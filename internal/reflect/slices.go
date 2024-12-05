package reflect

import "reflect"

// MakeInitializedSlice creates a slice of the given type with the given
// values initialized.
func MakeInitializedSlice(
	sliceType reflect.Type, values ...reflect.Value,
) reflect.Value {
	out := reflect.MakeSlice(sliceType, len(values), len(values))
	for i, v := range values {
		out.Index(i).Set(v)
	}
	return out
}

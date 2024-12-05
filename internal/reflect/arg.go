package reflect

import "reflect"

type Arg struct {
	Type

	// Whether the argument is variadic.
	IsVariadic bool

	// Whether the argument is a slice.
	IsSlice bool

	// Whether the argument is an array.
	IsArray   bool
	ArraySize int
}

func NewArg(t Type, isVariadic bool) *Arg {
	arg := &Arg{Type: t, IsVariadic: isVariadic}

	if t.Kind() == reflect.Array {
		arg.IsArray = true
		arg.ArraySize = t.Len()
	} else if t.Kind() == reflect.Slice || isVariadic {
		arg.IsSlice = true
	}

	return arg
}

// IsType returns whether the type matches the argument type.
func (a *Arg) IsType(t Type, inferInterfaces bool) bool {
	if a.IsArray && t.Kind() == a.Kind() && t.Len() == a.ArraySize {
		return matchesType(t.Elem(), a.Type.Elem(), inferInterfaces)
	} else if a.IsSlice && t.Kind() == a.Kind() {
		return matchesType(t.Elem(), a.Type.Elem(), inferInterfaces)
	}

	return matchesType(t, a.Type, inferInterfaces)
}

// matchesType returns whether toCheck is exactly expected,
// or assignable to expected.
func matchesType(toCheck, expected Type, inferInterfaces bool) bool {
	if inferInterfaces {
		return toCheck == expected || toCheck.AssignableTo(expected)
	}
	return toCheck == expected
}

package reflect

import "reflect"

type Arg struct {
	Type

	// The underlying type of a pointer, slice or array.
	UnderlyingType Type

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

	switch t.Kind() {
	case reflect.Array:
		arg.IsArray = true
		arg.UnderlyingType = t.Elem()
		arg.ArraySize = t.Len()
	case reflect.Slice:
		arg.IsSlice = true
		arg.UnderlyingType = t.Elem()
	}

	return arg
}

// IsType returns whether the type matches the argument type.
func (a *Arg) IsType(t Type, inferInterfaces bool) bool {
	if a.IsArray && t.Kind() == a.Kind() && t.Len() == a.ArraySize {
		return matchesType(t.Elem(), a.UnderlyingType, inferInterfaces)
	} else if a.IsSlice && t.Kind() == a.Kind() {
		return matchesType(t.Elem(), a.UnderlyingType, inferInterfaces)
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

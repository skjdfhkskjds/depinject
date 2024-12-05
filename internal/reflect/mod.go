package reflect

import "reflect"

type (
	Type  = reflect.Type
	Value = reflect.Value
)

var (
	TypeOf    = reflect.TypeOf
	ValueOf   = reflect.ValueOf
	MakeFunc  = reflect.MakeFunc
	MakeSlice = reflect.MakeSlice

	Interface = reflect.Interface
	Ptr       = reflect.Ptr
	Struct    = reflect.Struct
	Slice     = reflect.Slice
	Array     = reflect.Array
)

// IsError returns true if the given type is an error.
func IsError(t Type) bool {
	return t.AssignableTo(TypeOf((*error)(nil)).Elem())
}

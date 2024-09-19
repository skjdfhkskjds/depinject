package reflect

import "reflect"

type (
	Type  = reflect.Type
	Value = reflect.Value
)

var (
	TypeOf  = reflect.TypeOf
	ValueOf = reflect.ValueOf
)

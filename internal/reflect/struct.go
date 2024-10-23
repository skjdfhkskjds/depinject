package reflect

import (
	"reflect"
)

type Struct struct {
	Name string

	StructType  Type
	NamedFields []Type
}

func NewStruct(s any) (*Struct, error) {
	t, ok := s.(Type)
	if !ok || t.Kind() != reflect.Struct {
		return nil, ErrNotAStruct
	}

	// Loop through each field
	var fields []Type
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fields = append(fields, field.Type)
	}

	return &Struct{
		Name:        t.Name(),
		StructType:  t,
		NamedFields: fields,
	}, nil
}

// Constructor returns a function that constructs a new instance of the struct.
func (s *Struct) Constructor() *Func {
	// return MakeFunc(
	// 	reflect.FuncOf(s.NamedFields, []Type{s.StructType}, false),
	// 	func(args []Value) []Value {
	// 		structValue := reflect.New(s.StructType).Elem()
	// 		for i := range s.NamedFields {
	// 			structValue.Field(i).Set(args[i])
	// 		}
	// 		return []Value{structValue}
	// 	},
	// )

	
}
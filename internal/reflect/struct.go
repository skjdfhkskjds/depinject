package reflect

import (
	"reflect"
)

type Struct struct {
	Name string

	StructType Type
	Fields     []Type
}

func NewStruct(s any) (*Struct, error) {
	t, ok := s.(Type)
	if !ok || t.Kind() != reflect.Struct {
		return nil, ErrNotAStruct
	}

	// Loop through each field
	var fields []Type
	for i := 0; i < t.NumField(); i++ {
		fields = append(fields, t.Field(i).Type)
	}

	return &Struct{
		Name:       t.Name(),
		StructType: t,
		Fields:     fields,
	}, nil
}

// Constructor returns a function that constructs a new instance of the struct.
func (s *Struct) Constructor() *Func {
	return MakeNamedFunc(
		s.Fields,
		[]Type{s.StructType},
		func(args []Value) []Value {
			structValue := reflect.New(s.StructType).Elem()
			for i := range s.Fields {
				structValue.Field(i).Set(args[i])
			}
			return []Value{structValue}
		},
	)
}

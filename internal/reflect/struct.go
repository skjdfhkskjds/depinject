package reflect

import (
	"reflect"
)

type StructType struct {
	Name string

	Type   Type
	Fields []Type
}

func NewStruct(s any) (*StructType, error) {
	t, ok := s.(Type)
	if !ok || t.Kind() != Struct {
		return nil, ErrNotAStruct
	}

	// Loop through each field
	var fields []Type
	for i := 0; i < t.NumField(); i++ {
		fields = append(fields, t.Field(i).Type)
	}

	return &StructType{
		Name:   t.Name(),
		Type:   t,
		Fields: fields,
	}, nil
}

// Constructor returns a function that constructs a new instance of the struct.
func (s *StructType) Constructor() *Func {
	return MakeNamedFunc(
		s.Fields,
		[]Type{s.Type},
		func(args []Value) []Value {
			structValue := reflect.New(s.Type).Elem()
			for i := range s.Fields {
				structValue.Field(i).Set(args[i])
			}
			return []Value{structValue}
		},
		s.Name,
	)
}

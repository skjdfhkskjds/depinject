package reflect

import (
	"reflect"
)

type StructType struct {
	Name string

	Type Type

	// Note: We need to keep this duplicate mapping in order to
	// maintaining consistent iterative field order while also
	// decoupling the fields list from the comprehensive field list.
	Fields        []Type
	FieldsToIndex map[Type]int
}

func NewStruct(s any) (*StructType, error) {
	t, ok := s.(Type)
	if !ok || t.Kind() != Struct {
		return nil, ErrNotAStruct
	}

	// Loop through each field
	fields := make([]Type, t.NumField())
	fieldsToIndex := make(map[Type]int)
	for i := 0; i < t.NumField(); i++ {
		fields[i] = t.Field(i).Type
		fieldsToIndex[t.Field(i).Type] = i
	}

	return &StructType{
		Name:          t.Name(),
		Type:          t,
		Fields:        fields,
		FieldsToIndex: fieldsToIndex,
	}, nil
}

// Constructor returns a function that constructs a new instance of the struct.
func (s *StructType) Constructor() *Func {
	return MakeNamedFunc(
		s.Fields,
		[]Type{s.Type},
		func(args []Value) []Value {
			structValue := reflect.New(s.Type).Elem()
			for _, arg := range args {
				structValue.Field(s.FieldsToIndex[arg.Type()]).Set(arg)
			}
			return []Value{structValue}
		},
		s.Name,
	)
}

// Provider returns a function that takes in an instance of the struct
// and returns the value of each field as output.
func (s *StructType) Provider() *Func {
	return MakeNamedFunc(
		[]Type{s.Type},
		s.Fields,
		func(args []Value) []Value {
			structValue := args[0]
			outputs := make([]Value, 0)
			for _, t := range s.Fields {
				outputs = append(outputs, structValue.Field(s.FieldsToIndex[t]))
			}
			return outputs
		},
		s.Name,
	)
}

package reflect

import (
	"reflect"

	"github.com/skjdfhkskjds/depinject/internal/utils"
)

type StructType struct {
	Name string

	Type Type

	// Fields is a mapping of field types to their indices.
	Fields *utils.OrderedMap[Type, int]
}

func NewStruct(s any) (*StructType, error) {
	t, ok := s.(Type)
	if !ok || t.Kind() != Struct {
		return nil, ErrNotAStruct
	}

	// Loop through each field
	fields := utils.NewOrderedMap[Type, int]()
	for i := 0; i < t.NumField(); i++ {
		fields.Set(t.Field(i).Type, i)
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
		s.Fields.Keys(),
		[]Type{s.Type},
		func(args []Value) []Value {
			structValue := reflect.New(s.Type).Elem()
			for _, arg := range args {
				index, _ := s.Fields.Get(arg.Type())
				structValue.Field(index).Set(arg)
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
		s.Fields.Keys(),
		func(args []Value) []Value {
			structValue := args[0]
			outputs := make([]Value, 0)
			for _, t := range s.Fields.Keys() {
				index, _ := s.Fields.Get(t)
				outputs = append(outputs, structValue.Field(index))
			}
			return outputs
		},
		s.Name,
	)
}

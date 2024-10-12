package reflect

import (
	"reflect"
)

type Struct struct {
	Name string

	Fields []Type
}

func NewStruct(s any) (*Struct, error) {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Struct {
		return nil, ErrNotAStruct
	}

	// Loop through each field
	fields := make([]Type, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fields[i] = TypeOf(field.Type)
	}

	return &Struct{
		Name:   t.Name(),
		Fields: fields,
	}, nil
}

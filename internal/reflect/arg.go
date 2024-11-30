package reflect

type Arg struct {
	Type Type

	IsVariadic bool
}

func NewArg(t Type, isVariadic bool) *Arg {
	return &Arg{
		Type:       t,
		IsVariadic: isVariadic,
	}
}

// IsType returns whether the type t matches the argument type.
func (a *Arg) IsType(t Type, inferInterfaces bool) bool {
	if inferInterfaces {
		return t == a.Type || t.AssignableTo(a.Type) || t.ConvertibleTo(a.Type)
	}
	return t == a.Type
}

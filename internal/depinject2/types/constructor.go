package types

import (
	"fmt"

	"github.com/skjdfhkskjds/depinject/internal/reflect"
)

const (
	errValueNotFound = "value not found: %s"
)

// A constructor represents a single node in the container.
// It's responsibility is to tell a container how to interact
// with the underlying function, that is:
// - how to supply its dependencies
// - executing the function
// - retrieving its outputs
type Constructor struct {
	id string

	args    []*reflect.Arg
	outputs map[reflect.Type]reflect.Value

	// The wrapped function.
	fn *reflect.Func
}

func NewConstructor(fn any) (*Constructor, error) {
	wrappedFn, err := reflect.WrapFunc(fn)
	if err != nil {
		return nil, err
	}

	return NewConstructorFromFunc(wrappedFn), nil
}

func NewConstructorFromFunc(fn *reflect.Func) *Constructor {
	// TODO: Bring this back to ignore error output overlap.
	// fn.Ret = filterError(fn.Ret)
	args := make([]*reflect.Arg, len(fn.Args))
	for i, arg := range fn.Args {
		args[i] = reflect.NewArg(arg, i == len(fn.Args)-1 && fn.IsVariadic)
	}

	return &Constructor{
		id:      fn.Name,
		args:    args,
		outputs: make(map[reflect.Type]reflect.Value, len(fn.Ret)),
		fn:      fn,
	}
}

func (c *Constructor) ID() string {
	return c.id
}

func (c *Constructor) Dependencies() []*reflect.Arg {
	return c.args
}

func (c *Constructor) ValueOf(t reflect.Type) (reflect.Value, error) {
	if val, ok := c.outputs[t]; ok {
		return val, nil
	}
	return reflect.Value{}, fmt.Errorf(errValueNotFound, t.String())
}

func (c *Constructor) Outputs() map[reflect.Type]reflect.Value {
	return c.outputs
}

// Execute executes the constructor.
func (c *Constructor) Execute(args ...any) error {
	values, err := c.fn.Call(args...)
	if err != nil {
		return err
	}

	for _, value := range values {
		c.outputs[value.Type()] = value
	}

	return nil
}

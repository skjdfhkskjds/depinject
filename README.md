# Depinject 

[![Go Reference](https://pkg.go.dev/badge/github.com/skjdfhkskjds/depinject.svg)](https://pkg.go.dev/github.com/skjdfhkskjds/depinject)

This project is a runtime-based dependency injection framework written in Go designed to simplify the management of dependencies in Go applications. It provides a robust and flexible way to construct your application components and resolve their dependencies at runtime.

## Feature Support

Currently the framework supports the following features:

- Constructors which return types **exactly** as requested by another's constructor.
- Constructors which return types which **implement interfaces** requested by another's constructor.
- Supplying values directly into the container.

## Getting Started

### Clone the Project

To get started with this project, clone the repository to your local machine:

```bash
$ git clone https://github.com/skjdfhkskjds/depinject
$ cd depinject
```

### Build and Run

To run the project and see the dependency injection in action:

```bash
$ go run ./example/basic
```

or run

```bash
$ make run example=basic
```

## Example Usage

Here is a simple example demonstrating how to use the dependency injection framework:

```go
package main

import (
    "fmt"
    "github.com/skjdfhkskjds/depinject"
)

func NewFoo() *Foo {
    return &Foo{}
}

func NewBar(_ *Foo) *Bar {
    return &Bar{}
}

func NewFooBar(_ *Foo, _ *Bar) *FooBar {
    return &FooBar{}
}

func main() {
	container := depinject.NewContainer()

	// Supply a value into the container directly.
	if err := container.Supply(&Foo{}); err != nil {
		panic(err)
	}

	// Provide a set of constructors into the container.
	if err := container.Provide(
		NewBar,
		NewFooBar,
	); err != nil {
		panic(err)
	}

	// Invoke a function with the dependencies injected
	// to retrieve the FooBar instance.
	var fooBar *FooBar
	if err := container.Invoke(&fooBar); err != nil {
		panic(err)
	}
	fooBar.Print()
}
```

For more examples, check out the [examples](./examples) directory.

## Documentation

For more detailed documentation, refer to the code comments and the structured documentation within each package.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

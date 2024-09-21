# Depinject 

[![Go Reference](https://pkg.go.dev/badge/github.com/skjdfhkskjds/depinject.svg)](https://pkg.go.dev/github.com/skjdfhkskjds/depinject)

This project is a runtime-based dependency injection framework written in Go designed to simplify the management of dependencies in Go applications. It provides a robust and flexible way to construct your application components and resolve their dependencies at runtime.

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

## Example Usage

Here is a simple example demonstrating how to use the dependency injection framework:

```go
package main

import (
    "fmt"
    "github.com/yourgithubusername/dependency-injection-framework"
)

type Foo struct{}

func NewFoo() *Foo {
    return &Foo{}
}

type Bar struct{}

func NewBar(_ *Foo) *Bar {
    return &Bar{}
}

type FooBar struct{}

func NewFooBar(_ *Foo, _ *Bar) *FooBar {
    return &FooBar{}
}

func (f *FooBar) Print() {
    fmt.Println("Hello from FooBar!")
}

func main() {
    container := depinject.NewContainer()

    if err := container.Provide(
        NewFoo,
        NewBar,
        NewFooBar,
    ); err != nil {
        panic(err)
    }

    var fooBar *FooBar
    if err := container.Invoke(&fooBar); err != nil {
        panic(err)
    }
    fooBar.Print()
}
```

For more examples, check out the [example](./example) directory.

## Documentation

For more detailed documentation, refer to the code comments and the structured documentation within each package.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

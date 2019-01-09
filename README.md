# Tour of Golang

A Tour of Golang - First learning Go Programming Language

[The tour of Golang](https://tour.golang.org)

## Go Playground - Online Editor

[The Go Playground - Available for share](https://play.golang.org/)


## Notes when install and setup Go - IMPORTANTS

### On MaxOS/Linux
#### Installation
Follow the instructions for your platform to install the Go tools: https://golang.org/doc/install#install. It is recommended to use the default installation settings.

On Mac OS X and Linux, by default Go is installed to directory `/usr/local/go/`, and the GOROOT environment variable is set to `/usr/local/go/bin`.

#### Set your GOPATH
Your Go working directory (GOPATH) is where you store your Go code. It can be any path you choose but must be separate from your Go installation directory (GOROOT).

The following instructions describe one way you can set your GOPATH. Refer to the official Go documentation for more details: https://golang.org/doc/code.html.

**Mac OS X and Linux**

Set the `GOPATH` environment variable for your workspace:
```bash
export GOPATH=$HOME/go
```

Also set the `GOPATH/bin` variable, which is used to run compiled Go programs.
```bash
export PATH=$PATH:$GOPATH/bin
```

### On Windows

## Quickstart

### Run single file
```bash
go run [path/to/file].go
```

This command only runs one file with **main()** function declared. It's not working with other files in the same package.

### Build package
```bash
go build
```

This command build all files in package **main** into build file **tour-golang**, which is executable, you can run 

```bash
./tour-golang
```

### Run package at runtime
```bash
go run *.go
```

This command is similar with build package and run the executable file. It will link files `.go` and run them.


## Notes with packages in Golang
### Package declaration
In a directory, one package is declared at the same level. If we want to create another package, create sub-folder and declare with new package for subfolder.

Name of package should be name of directory.

### Export modules (type, struct, interface, function, variable) from pacakge
Golang uses first character with capitalized (from `A-Z`) is exported, that means other package can read.

Example in [directory packages](./packages)

```go
package A
// Foo : exported function, should be comment by go-lint
func Foo() {}
func bar() {
	// this is private in package A, but other file in package A can use bar()
}
// Animal : type animal is exported, shoul be comment by go-lint
type Animal struct {
	a string // that is private
	B int    // that is exported, that mean public
}
```

And in package B

```go
package B

import "github.com/huynhsamha/tour-golang/packages/A"

func todo() {
	A.Foo() // OK

	// A.bar()
	// cannot refer to unexported name A.bar

	// x := A.Animal{"123", 4}
	// implicit assignment of unexported field 'a' in A.Animal literal

	x := A.Animal{B: 123}
	// x.a = "123"
	// x.a undefined (cannot refer to unexported field or method a)

	x.B = 123
}
```


## Golang topics 
+ [Variable - Function declaration](./variable_function_declare/variable_function_declare.go)
+ [For Statement](./for_statement/for_statement.go)
+ [If Statement](./if_statement/if_statement.go)
+ [Switch Statement](./switch/switch.go)
+ [Array](./array/array.go)
+ [Map](./map/map.go)
+ [Map with interface](./map_interface/map_interface.go)
+ [Function as values](./function_values/function_values.go)
+ [Function as closures](./function_closures/function_closures.go)
+ [Interface](./interface/interface.go)
+ [Struct](./struct/struct.go)
+ [Example Fibonaci](./fibonacci/fibonacci.go)
+ [Example Power](./power/power.go)
+ [Spread (...) - Unpacking array to arguments](./spread_operator/spread_operator.go)
+ [Same type in struct declaration](./same_type_in_struct/same_type_in_struct.go)
+ [Exported modules in package](./packages)
+ [Sortings](./sortings)
+ [Error return](./return_error/return_error.go)
+ [Read / Write files](./files)
+ [JWT - jwt-go Golang JSON Web Token](./jwt-go/jwt-go.go)
+ [Crypto (SHA256, MD5, ...)](./crypto)
+ [Random Numbers](./random_numbers/random_numbers.go)
+ [Random Hex, String,... Simple](./random_string/random_string.go)
+ [Go Routine](./goroutines/goroutines.go)

## OOP in Golang

[Detail](./OOP.md)

## DEP - Dependency Pakage Management Go

[Detail](./DEP.md)

## Best Documents or References
+ Go by Examples: https://gobyexample.com/
+ Go Bootcamp: http://www.golangbootcamp.com/book
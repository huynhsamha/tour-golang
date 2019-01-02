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


## Golang topics 
+ [Variable - Function declaration](./var_func.go)
+ [For](./for.go)
+ [If](./if.go)
+ [Array](./array.go)
+ [Map](./map.go)
+ [Map with interface](./map_interface.go)
+ [Function as values](./func_values.go)
+ [Function as closures](./func_closures.go)
+ [Interface](./interface.go)
+ [Struct](./struct.go)
+ [Example Fibonaci](./fibo.go)
+ [Example Power](./pow.go)
+ [Spread (...) - Unpacking array to arguments](./spread.go)
+ [Same type in struct declaration](./same_type_in_struct.go)

## OOP in Golang

[Detail](./OOP.md)

## DEP - Dependency Pakage Management Go

[Detail](./DEP.md)

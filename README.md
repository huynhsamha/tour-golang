# Tour of Golang

A Tour of Golang - First learning Go Programming Language

[The tour of Golang](https://tour.golang.org)

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
+ [Array]("./array.go")
+ [Map]("./map.go")
+ [Map with interface]("./map_interface.go")
+ [Function as values]("./func_values.go")
+ [Function as closures]("./func_closures.go")
+ [Interface]("./interface.go")
+ [Struct]("./struct.go")
+ [Example Fibonaci]("./fibo.go")
+ [Example Power]("./pow.go")

## OOP in Golang

[Detail]("./OOP.md")

## DEP - Dependency Pakage Management Go

[Detail]("./DEP.md")

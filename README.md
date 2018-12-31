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
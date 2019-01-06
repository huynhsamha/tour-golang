package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
)

func main() {
	dirname, err := filepath.Abs("./files")
	data, err := ioutil.ReadFile(path.Join(dirname, "./data.txt"))
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	// data is []byte
	err = ioutil.WriteFile(path.Join(dirname, "./clone.txt"), data, 0755) // -rwxr-xr-x
	if err != nil {
		panic(err)
	}
}

package main

import (
	"errors"
	"fmt"
)

func functionReturnError(a, b int) (int, error) {
	if a+b > 0 {
		return -1, errors.New("Sum > 0")
	}
	return a - b, nil
}

func main() {
	res, err := functionReturnError(-1, 2)
	fmt.Println(res, err)
	// -1 Sum > 0

	res, err = functionReturnError(1, -2)
	fmt.Println(res, err)
	// 3 <nil>
}

package main

import "fmt"

type BasicType struct {
	a string
	b int
	c float32
}

// WTH -> very complicated type when declare
type ComplicatedType struct {
	a BasicType
	b BasicType
	c BasicType
	d BasicType
	e BasicType
	f BasicType
}

// Simplify this type
type SimplyType struct {
	a, b, c, d, e, f BasicType
}

func main() {
	x := BasicType{}

	y := ComplicatedType{x, x, x, x, x, x}
	z := SimplyType{x, x, x, x, x, x}

	fmt.Println(y, z)
}

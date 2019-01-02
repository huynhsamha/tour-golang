package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"time"
)

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func swap(x, y string) (string, string) {
	return y, x
}

func nakedReturn(a, b float32) (x, y float32) {
	x = a + b
	y = x * (a - b)
	return
}

var p, q bool // declare global scope
var g, h int = 1, 2
var o, k = 5, 6

// p := 1 // error syntax, global scope require var

func main() {
	fmt.Println("Hello Golang")
	fmt.Println("Now is  : ", time.Now())
	fmt.Println("Random  : ", rand.Intn(15))
	fmt.Println("SQRT    : ", math.Sqrt(18))
	fmt.Println("PI      : ", math.Pi)

	println("ABC")
	fmt.Println("ABC")

	fmt.Println(add(5, 10))
	fmt.Println(sub(5, 10))

	b, a := swap("A", "B")
	println(a, b)

	i, j := nakedReturn(4.5, 5.6)
	println(i, j)

	var p = true // new var override global declare p
	println(p)
	p = false // re assign
	println(p)
	// p := true // error redeclare
	// var p = true // error redeclare

	var (
		a1 = false
		b1 = 1.5
		c1 complex128
		d1 = "123"
	)

	println(a1, b1, c1)
	fmt.Printf("Type = %T, Value = %v", a1, a1)
	println() // print with format
	fmt.Printf("Type = %T, Value = %v", b1, b1)
	println() // print with format
	fmt.Printf("Type = %T, Value = %v", c1, c1)
	println() // print with format
	fmt.Printf("Type = %T, Value = %v", d1, d1)
	println() // print with format

	/** Type Conversion */
	g := 1
	f := float32(g)
	fmt.Println(reflect.TypeOf(g))
	fmt.Println(reflect.TypeOf(f))

	/** Constant */
	const wth = "世界"
	println(wth)
}

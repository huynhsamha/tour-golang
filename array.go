package main

import (
	"fmt"
	"strings"
)

func LogArray(a []string) {
	fmt.Printf("len = %d, cap = %d\n", len(a), cap(a))

	// loop array by capacity with idx and val
	for idx, val := range a {
		fmt.Println(idx, val) // eg. 0 "a"
	}

	// loop array by capacity with only index
	for idx := range a {
		fmt.Println(idx, a[idx]) // eg. 0 "a"
	}

	// loop array by capacity with only val
	for _, val := range a {
		fmt.Println(val) // eg. "a"
	}
}

func main() {

	// Create array of string
	a := []string{"a", "b", "c", "d", "e", "f"}

	// Create new array with capacity = 8
	b := make([]string, 8)

	LogArray(a) // len 6, cap 6
	LogArray(b) // len 8, cap 8, initial value is ""

	// Slice array (extend length | drop elements),
	// but capacity maybe keep without extended
	LogArray(a[:0])  // len 0, cap 6
	LogArray(a[:5])  // len 5, cap 6
	LogArray(a[1:])  // len 5, cap 5, drop 0th
	LogArray(a[2:5]) // len 3, cap 4, drop 0th, 1st, 5th, keep memory of the last elements

	// Join array of elements to string
	fmt.Println(strings.Join(a, ",")) // "a,b,c,d,e,f"

	// Extend capacity of array
	LogArray(append(a, "g"))           // append "g" to array, raise capacity if not enough
	LogArray(append(a, "g", "h", "i")) // append "g", "h", "i" to array, raise capacity if not enough
}

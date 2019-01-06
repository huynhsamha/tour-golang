package main

import (
	"fmt"
	"sort"
)

func main() {
	// Sort increasing int type
	a := []int{1, 5, 4, 6, 9, 2, 2, 5}
	sort.Ints(a)
	fmt.Println(a)
	fmt.Println(sort.IntsAreSorted(a)) // check sorted

	// Sort increasing string type
	s := []string{"A", "b", "B", "c", "A", "B", "C", "X", "d"}
	sort.Strings(s)
	fmt.Println(s)

	// Sort increasing float type
	f := []float64{1.2, 3.5, 5.1, 2.4, -9.3, -2, 0}
	sort.Float64s(f)
	fmt.Println(f)

	// Sort a slice in array
	b := []int{9, 9, 1, 5, 4, 2, 3, 1, 2, 3}
	sort.Ints(b[2:6])
	fmt.Println(b) // [9 9 1 2 4 5 3 1 2 3]
}

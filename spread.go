package main

import "fmt"

/**
Spread : Unpacking array as arguments

Example: [1,2,3] -> 1,2,3
In JavaScript: ...[1,2,3] -> 1,2,3
*/

func sum(args ...int) int {
	sum := 0
	for _, v := range args {
		sum = sum + v
	}

	return sum
}

func main_spread() {

	arr := []int{2, 4, 6, 8}

	sum := sum(arr...)

	fmt.Println("Sum is ", sum)
}

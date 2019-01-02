package main

import "fmt"

func power(x, n int) int64 {
	if n == 1 {
		return int64(x)
	}
	m := power(x, n/2)
	if n%2 == 0 {
		return m * m
	}
	return m * m * int64(x)
}

func main() {
	fmt.Println(power(5, 10))
	fmt.Println(power(2, 5))
	fmt.Println(power(1, 1e18))
}

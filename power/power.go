package main

import "fmt"

func pow(x, n int) int64 {
	if n == 1 {
		return int64(x)
	}
	m := pow(x, n/2)
	if n%2 == 0 {
		return m * m
	}
	return m * m * int64(x)
}

func main_pow() {
	fmt.Println(pow(5, 10))
	fmt.Println(pow(2, 5))
	fmt.Println(pow(1, 1e18))
}

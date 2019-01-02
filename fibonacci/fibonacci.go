package main

import "fmt"

func fibo_array(n int) []int {
	f := make([]int, n+1)
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f
}

func main() {
	n := 10

	fmt.Println(fibo_array(n))
}

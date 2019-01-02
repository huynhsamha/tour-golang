package main

import "fmt"

func main_for() {
	sum := 0
	for i := 1; i < 10; i++ {
		fmt.Println(i)
		sum += i
	}

	fmt.Println("Sum = ", sum)

	sum = 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}

package main

import "fmt"

func f(from string) {
	for i := 0; i < 5; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	go f("A")
	go f("B")
	go f("C")
	go f("D")
	go f("E")
	go f("F")

	f("X")

	go func(msg string) {
		fmt.Println(msg)
	}("----------------")

	// requires we press a key before the program exits
	fmt.Scanln()

	// Press <Enter> and show this line
	fmt.Println("Done")
}

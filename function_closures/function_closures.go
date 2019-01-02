package main

import "fmt"

/**
Function closures
Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
*/

func getFunction(operator string) func(int, int) int {
	switch operator {
	case "+":
		return func(a, b int) int { return a + b }
	case "-":
		return func(a, b int) int { return a - b }
	case "*":
		return func(a, b int) int { return a * b }
	case "/":
		return func(a, b int) int { return a / b }
	default:
		return nil
	}
}

func LogFunc(a, b int, op string) {
	f := getFunction(op)
	if f == nil {
		fmt.Println("Can't execute operator " + op)
	} else {
		fmt.Printf("%d %s %d = %d\n", a, op, b, f(a, b))
	}
}

func main_func_closures() {
	a := 8
	b := 2
	LogFunc(a, b, "+")
	LogFunc(a, b, "-")
	LogFunc(a, b, "*")
	LogFunc(a, b, "/")
	LogFunc(a, b, "%")
}

/**** Results ****

8 + 2 = 10
8 - 2 = 6
8 * 2 = 16
8 / 2 = 4
Can't execute operator %

*****************/

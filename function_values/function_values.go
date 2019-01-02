package main

import "fmt"

/**
Function values
Functions are values too. They can be passed around just like other values.

Function values may be used as function arguments and return values.
*/

func ajax(url string, method string, cb func(err interface{}, data string)) {
	fmt.Println("AJAX: start")
	message := "Successfully"
	// run callback function
	cb(nil, message)
	fmt.Println("AJAX: end")
}

func main() {

	// Pass alias function to parameter of a function
	ajax("http://ajax.io", "POST", func(err interface{}, data string) {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(data)
		}
	})

	// Declare new function with assign operator :=
	cb := func(err interface{}, data string) {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(data)
		}
	}

	// Pass function to parameter of a function
	ajax("http://ajax.io", "POST", cb)
}

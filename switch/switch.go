package main

import (
	"fmt"
	"time"
)

func main() {
	// Normal
	a := 1
	switch a {
	case 1:
		fmt.Println("Equal 1")
	case 2:
		fmt.Println("Equal 2")
	case 5:
		fmt.Println("Equal 5")
	default:
		fmt.Println("No Match")
	}

	// Multiple match
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// If/else
	t := time.Now()
	switch {
	case t.Hour() > 12:
		fmt.Println("Current is", t.Hour()-12, "h AM")
	default:
		fmt.Println("Current is", t.Hour(), "h AM")
	}

	// Type Assert
	check := func(v interface{}) {
		switch t := v.(type) {
		case int:
			fmt.Println("This is int")
		case bool:
			fmt.Println("This is bool")
		case float32:
			fmt.Println("This is float 32")
		case float64:
			fmt.Println("This is float 64")
		default:
			fmt.Println("What the hell", t)
		}
	}
	check(1)
	check(1 == 1)
	check(float32(1.5))
	check(1.5)
	check("123")
}

package main

import "fmt"

func main() {

	m := map[string]interface{}{
		"Id":       10005,
		"Username": "alice",
		"Password": "tw8493ur9023u029c",
		"DisplayName": map[string]string{
			"FirstName": "Alice",
			"LastName":  "Marco",
		},
	}

	for key, val := range m {
		fmt.Println(key, val)
	}
}

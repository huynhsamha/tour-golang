package main

import "fmt"

func LogMap(m map[string]int) {
	for key, val := range m {
		fmt.Println(key, val)
	}

	// loop map with only value
	for _, val := range m {
		fmt.Println(val)
	}

	// loop map with only key
	for key := range m {
		fmt.Println(key)
	}
}

func TestExist(m map[string]int, key string) {
	// the second param is true/false if key is existen in map
	val, exist := m[key]
	fmt.Println(val, exist)
}

func main_map() {

	// Create new map string -> int
	m := make(map[string]int)
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30

	LogMap(m)

	TestExist(m, "a") // 10, true
	TestExist(m, "t") // 0, false

	// Remove key "a" from map m
	delete(m, "a")
	TestExist(m, "a") // 0, false

	// initialize map string -> int without make()
	m = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(m)
}

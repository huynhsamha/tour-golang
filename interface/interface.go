package main

import "fmt"

// Animal : interface for Dog and Cat
type Animal interface {
	Speak()
	Voice() string // get voice of animal
}

// Dog : implement Animal
type Dog struct {
}

// Cat : implement Animal
type Cat struct {
}

func (dog Dog) Speak() {
	fmt.Println(dog.Voice())
}

func (cat Cat) Speak() {
	fmt.Println(cat.Voice())
}

// Implement Voice for class Dog
func (dog Dog) Voice() string {
	return "Woof"
}

// Implement Voice for class Cat
func (cat Cat) Voice() string {
	return "Meow"
}

func main_interface() {
	myPets := []Animal{
		Dog{},
		Dog{},
		Cat{},
		Dog{},
		Cat{},
		Cat{},
	}

	for _, pet := range myPets {
		pet.Speak()
	}
}

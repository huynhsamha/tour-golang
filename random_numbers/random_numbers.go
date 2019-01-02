package main

import "time"
import "fmt"
import "math/rand"

func main() {

	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Float64())

	// `5.0 <= f' < 10.0`.
	fmt.Println((rand.Float64() * 5) + 5)

	// The default number generator is deterministic, so it'll
	// produce the same sequence of numbers each time by default.
	// To produce varying sequences, give it a seed that changes.
	// Note that this is not safe to use for random numbers you
	// intend to be secret, use `crypto/rand` for those.
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Println(r1.Intn(100))
	fmt.Println(r1.Float64())
	fmt.Println((r1.Float64() * 5) + 5)

	// If you seed a source with the same number, it
	// produces the same sequence of random numbers.
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)

	fmt.Println(r2.Intn(100))
	fmt.Println(r2.Float64())
	fmt.Println((r2.Float64() * 5) + 5)
}

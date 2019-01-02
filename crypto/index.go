package main

import (
	"fmt"
)

func main() {
	s := "Hello Alice, Bob and Trudy"

	fmt.Println(SHA1(s))
	fmt.Println(SHA256(s))
	fmt.Println(MD5(s))

	// b297666ebffedb643aaa99b7445b60a2e220f0e0
	// c6e8d6971eaed8a22673cc0dd13868b31d9f3161d50c98a456865d669fcd48bf
	// 5bcf65c81f2ee888c3e029eca5ccf61f
}

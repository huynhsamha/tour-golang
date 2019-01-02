package main

import (
	"crypto/sha1"
	"fmt"
)

// SHA1 : SHA1 Hash
func SHA1(data string) string {
	hash := sha1.New()
	hash.Write([]byte(data))
	byteSlice := hash.Sum(nil)
	hex := fmt.Sprintf("%x", byteSlice)
	return hex
}

package main

import (
	"crypto/sha256"
	"fmt"
)

// SHA256 : SHA256 Hash
func SHA256(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	byteSlice := hash.Sum(nil)
	hex := fmt.Sprintf("%x", byteSlice)
	return hex
}

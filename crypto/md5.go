package main

import (
	"crypto/md5"
	"fmt"
)

// MD5 : MD5 Hash
func MD5(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	byteSlice := hash.Sum(nil)
	hex := fmt.Sprintf("%x", byteSlice)
	return hex
}

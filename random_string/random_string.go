package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	token := make([]byte, 16)
	rand.Read(token)
	tokenString := hex.EncodeToString(token)
	fmt.Println(tokenString)
}

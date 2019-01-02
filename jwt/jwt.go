package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
}	

func generateToken() string error {
	issuedAt := time.Now().Unix()
	notBefore := time.Now().Add(30 * time.Second).Unix()
	expire := time.Now().Add(30 * time.Second).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat":  issuedAt,
		"nbf":  notBefore,
		"exp":  expire,
		"data": "Hello World",
	})

	// Sign and get the complete encoded token as a string using the secret
	jwtSecret := "aewihr034v923u04912ujce092u"
	tokenString, err := token.SignedString([]byte(jwtSecret))

	return tokenString, err
}

package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("aewihr034v923u04912ujce092u")

func main() {
	token, err := generateToken()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(token)

	tryParse(token) // output: token is not valid yet

	time.Sleep(2 * time.Second) // after 2s -> now is valid
	tryParse(token)             // output: Successfully 1546417482 1546417484 1546417489 Hello Go lang JWT

	time.Sleep(1 * time.Second) // after 3s -> still valid
	tryParse(token)             // output: Successfully 1546417482 1546417484 1546417489 Hello Go lang JWT

	time.Sleep(1 * time.Second) // after 4s -> still valid
	tryParse(token)             // output: Successfully 1546417482 1546417484 1546417489 Hello Go lang JWT

	time.Sleep(3 * time.Second) // after 7s -> still valid
	tryParse(token)             // output: Successfully 1546417482 1546417484 1546417489 Hello Go lang JWT

	time.Sleep(1 * time.Second) // after 8s -> now is expired
	tryParse(token)             // output: token is expired by 1s
}

// MyCustomClaims :
type MyCustomClaims struct {
	jwt.StandardClaims
	Data string `json:"data"`
}

func tryParse(token string) {
	fmt.Println("Trying:")
	claims, err := parseToken(token)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully", claims.IssuedAt, claims.NotBefore, claims.ExpiresAt, claims.Data)
	}
}

func generateToken() (string, error) {
	claims := MyCustomClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(7 * time.Second).Unix(), // only valid in 5s after valid
			NotBefore: time.Now().Add(2 * time.Second).Unix(), // just valid after 2s
		},
		Data: "Hello Go lang JWT",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwtSecret)

	return tokenString, err
}

func parseToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

package authentication

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	User string `json:"user"`
	jwt.StandardClaims
}

func GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		"thisIsAnUserTest",
		jwt.StandardClaims{
			ExpiresAt: 3600,
			Issuer:    "app",
		},
	})

	mySecret := "summerReminder"
	signedToken, err := token.SignedString([]byte(mySecret))

	return signedToken, err
}

func ValidateToken(receivedToken string) (bool, error) {
	_, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		mySecret, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		log.Println(mySecret)
		return mySecret, nil
	})
	return false, err
}

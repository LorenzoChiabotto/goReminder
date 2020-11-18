package authentication

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type MyCustomClaims struct {
	User string `json:"user"`
	jwt.StandardClaims
}

func GenerateToken() (string, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		return "", fmt.Errorf("Error loading .env file: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		"thisIsAnUserTest",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			Issuer:    "endpointValidateGo",
		},
	})

	signedToken, err := token.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		return "", fmt.Errorf("Error signing the token: %v", err)
	}
	return signedToken, err
}

func ValidateToken(receivedToken string) error {
	err := godotenv.Load("./.env")
	splittedToken := strings.Split(receivedToken, " ")
	if len(splittedToken) != 2 {
		return fmt.Errorf("Invalid token structure: %v", receivedToken)
	}

	token, err := jwt.Parse(splittedToken[1], func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("secret")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Println(claims)
	}
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

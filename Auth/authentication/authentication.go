package authentication

import (
	"fmt"
	"time"

	"../database"
	"github.com/dgrijalva/jwt-go"
)

//MyCustomClaims defines claims for the token
type MyCustomClaims struct {
	User string `json:"user"`
	jwt.StandardClaims
}

//GenerateToken generates a token based on the user info.
func GenerateToken(user User) (*Token, error) {
	loggedUser, err := database.GetUser(user.Name, user.Password)
	if err != nil {
		return nil, err
	}

	expirationTime := time.Now().Add(time.Minute * 60).Unix()
	issuer := "authenticationGo"

	key := []byte("this_is_a_test_key_that_must_be_replaced_in_prod")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		loggedUser.Name,
		jwt.StandardClaims{
			ExpiresAt: expirationTime,
			Issuer:   issuer,
		},
	})

	signedToken, err := token.SignedString(key)
	if err != nil {
		return nil, fmt.Errorf("error signing the token: %v", err)
	}

	return &Token{
		Token:        signedToken,
		ExpiresIn:    expirationTime,
	}, err
}

//ValidateToken validates a received token
/*func ValidateToken(receivedToken string) error {
	splitToken := strings.Split(receivedToken, " ")
	if len(splitToken) != 2 {
		return fmt.Error("invalid token structure: %v", receivedToken)
	}

	_, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Error("unexpected signing method: %v", token.Header["alg"])
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		secret, err := database.GetSecret(claims["user"].(string))

		if err != nil {
			return nil, fmt.Error("invalid token: %v", err)
		}

		return []byte(*secret), nil
	})

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}*/

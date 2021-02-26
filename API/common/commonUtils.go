package commonUtils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
)

//ValidateToken validates token passed by parameter
func ValidateToken (tokenStr string) (*jwt.Token,error) {
	key := []byte("this_is_a_test_key_that_must_be_replaced_in_prod")

	//this is in case i use Auth0 (for a future)
	//cleanedToken := strings.ReplaceAll(tokenStr, "bearer ", "")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		log.Println("There was an error when parsing the token:", err)
		return nil, err
	}
	if token.Valid {
		return token, nil
	}
	return nil, fmt.Errorf("the token was not able to validate: %v", err)
}
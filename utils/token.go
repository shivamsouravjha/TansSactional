package utils

import (
	"github.com/dgrijalva/jwt-go"
)

func TokenParse(myToken string, myKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(myKey), nil
	})
	return token, err
}

func GenerateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("mySigningKey"))
	return tokenString, err
}

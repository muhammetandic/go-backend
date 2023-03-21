package jwtAuth

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("DenemeJWTAnahtarı")

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		message := fmt.Errorf("something went wrong: %s", err.Error())
		return "", message
	}

	return tokenString, nil
}

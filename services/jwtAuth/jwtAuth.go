package jwtAuth

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type UserDetails struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Uid      int    `json:"uid"`
	UserType string `json:"user_type"`
	jwt.RegisteredClaims
}

var secretKey = []byte("DenemeJWTAnahtarı")

func GenerateTokens(username string) (string, string, error) {
	claims := &UserDetails{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Minute * 30)),
			Subject:   "antpos",
		},
	}

	refreshClaims := &UserDetails{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * 24 * 180)),
			Subject:   "antpos",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretKey)
	if err != nil {
		message := fmt.Errorf("something went wrong: %s", err.Error())
		return "", "", message
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(secretKey)
	if err != nil {
		message := fmt.Errorf("something went wrong: %s", err.Error())
		return "", "", message
	}
	return token, refreshToken, nil
}

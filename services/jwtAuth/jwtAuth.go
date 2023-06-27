package jwtAuth

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/muhammetandic/go-backend/main/core/models"
)

type UserDetails struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	jwt.RegisteredClaims
}

var secretKey = []byte("DenemeJWTAnahtarı")

func GenerateTokens(username string) (*models.LoginResponse, error) {
	login := &models.LoginResponse{}
	login.Username = username

	claims := &UserDetails{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "antpos",
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Minute * 30)),
			Subject:   "auth",
		},
	}

	refreshClaims := &UserDetails{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "antpos",
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * 24 * 180)),
			Subject:   "auth",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretKey)
	if err != nil {
		message := fmt.Errorf("something went wrong: %s", err.Error())
		return nil, message
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(secretKey)
	if err != nil {
		message := fmt.Errorf("something went wrong: %s", err.Error())
		return nil, message
	}

	login.AccessToken = token
	login.AccessTokenExpiresAt = claims.ExpiresAt.Time
	login.RefreshToken = refreshToken
	login.RefreshTokenExpiresAt = refreshClaims.ExpiresAt.Time
	return login, nil
}

func ValidateToken(signedToken string) (*UserDetails, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&UserDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*UserDetails)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}
	if claims.ExpiresAt.Before(time.Now().Local()) {
		err = errors.New("token expired")
		return nil, err
	}
	return claims, nil
}

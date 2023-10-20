package jwtAuth

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/muhammetandic/go-backend/main/core/models"
)

const (
	issuer                 = "antpos"
	subject                = "auth"
	expiresAtThirtyMinutes = time.Minute * 30
	expiresAtSixMonths     = time.Hour * 24 * 180
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
			Issuer:    issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(expiresAtThirtyMinutes)),
			Subject:   subject,
		},
	}

	refreshClaims := &UserDetails{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(expiresAtSixMonths)),
			Subject:   subject,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretKey)
	if err != nil {
		message := fmt.Errorf("failed to generate access token: %s", err.Error())
		return nil, message
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(secretKey)
	if err != nil {
		message := fmt.Errorf("failed to generate refresh token: %s", err.Error())
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

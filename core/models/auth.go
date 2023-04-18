package models

import "time"

type Auth struct {
	Username string `json:"username" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Register struct {
	Username string `json:"username" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"fullname" binding:"required"`
}

type LoginResponse struct {
	Username              string    `json:"username"`
	AccessToken           string    `json:"accessToken"`
	AccessTokenExpiresAt  time.Time `json:"accessTokenExpiresAt"`
	RefreshToken          string    `json:"refreshToken"`
	RefreshTokenExpiresAt time.Time `json:"refreshTokenExpiresAt"`
}

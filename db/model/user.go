package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `json:"email"`
	Password     string `json:"password"`
	Fullname     string `json:"fullname"`
	RefreshToken string `json:"refreshToken"`
}

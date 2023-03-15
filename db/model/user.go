package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `json:"email" gorm:"unique; size: 100"`
	Password     string `json:"password" gorm:"size:100"`
	Fullname     string `json:"fullname" gorm:"size:200"`
	RefreshToken string `json:"refreshToken"`
}

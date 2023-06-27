package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model
	Email        string `json:"email" gorm:"unique; size:100" binding:"required"`
	Password     string `json:"password" gorm:"size:100" binding:"required"`
	Fullname     string `json:"fullname" gorm:"size:200" binding:"required"`
	RefreshToken string `json:"refreshToken"`
	Roles        []UserToRole
}

type UserDto struct {
	Email    string `json:"email" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

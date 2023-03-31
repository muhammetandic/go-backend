package services

import (
	"fmt"
	"log"

	"github.com/muhammetandic/go-backend/main/db"
	"github.com/muhammetandic/go-backend/main/db/model"
	"github.com/muhammetandic/go-backend/main/services/jwtAuth"
)

func Login(info model.Auth) (string, error) {
	var user model.User

	userRecord := db.Instance.Where("email= ?", info.Email).First(&user)
	if userRecord.Error != nil {
		return "", fmt.Errorf("user not found")
	}

	passwordError := user.CheckPassword(info.Password)
	if passwordError != nil {
		return "", fmt.Errorf("password incorrect")
	}

	token, err := jwtAuth.GenerateToken(info.Email)
	if err != nil {
		return "", fmt.Errorf("internal server error")
	}

	return token, nil
}

func Register(info model.Register) error {
	newUser := model.User{Email: info.Email, Fullname: info.FullName, Password: info.Password}

	if err := newUser.HashPassword(newUser.Password); err != nil {
		log.Println(err.Error())
		return fmt.Errorf("password couldn't hashed")
	}

	if err := db.Instance.Create(&newUser).Error; err != nil {
		log.Println(err.Error())
		return fmt.Errorf("couldn't create user")
	}

	return nil
}

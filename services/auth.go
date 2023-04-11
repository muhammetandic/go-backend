package services

import (
	"fmt"
	"log"

	"github.com/muhammetandic/go-backend/main/core/models"
	"github.com/muhammetandic/go-backend/main/db"
	"github.com/muhammetandic/go-backend/main/db/model"
	"github.com/muhammetandic/go-backend/main/services/jwtAuth"
)

func Login(info models.Auth) (*models.LoginResponse, error) {
	var user model.User

	userRecord := db.Instance.Where("email= ?", info.Username).First(&user)
	if userRecord.Error != nil {
		return nil, fmt.Errorf("user not found")
	}

	passwordError := user.CheckPassword(info.Password)
	if passwordError != nil {
		return nil, fmt.Errorf("password incorrect")
	}

	response, err := jwtAuth.GenerateTokens(info.Username)
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	return response, nil
}

func Register(info models.Register) error {
	newUser := model.User{Email: info.Username, Fullname: info.FullName, Password: info.Password}

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

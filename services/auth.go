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

	db, err := db.Connect()
	if err != nil {
		log.Println(err.Error())
		return "", fmt.Errorf("couldn't connect database")
	}

	if err := db.Where("email= ? AND password= ?", info.Email, info.Password).First(&user).Error; err != nil {
		return "", fmt.Errorf("login incorrect")
	}

	token, err := jwtAuth.GenerateToken(info.Email)
	if err != nil {
		return "", fmt.Errorf("internal server error")
	}

	return token, nil
}

func Register(info model.Register) error {
	newUser := model.User{Email: info.Email, Fullname: info.FullName, Password: info.Password}

	db, err := db.Connect()
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("couldn't connect database")
	}

	if err := db.Create(&newUser).Error; err != nil {
		log.Println(err.Error())
		return fmt.Errorf("couldn't create user")
	}

	return nil
}

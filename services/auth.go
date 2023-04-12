package services

import (
	"log"

	"github.com/muhammetandic/go-backend/main/core/enums/errors"
	"github.com/muhammetandic/go-backend/main/core/models"
	"github.com/muhammetandic/go-backend/main/db"
	"github.com/muhammetandic/go-backend/main/db/model"
	"github.com/muhammetandic/go-backend/main/services/jwtAuth"
)

func Login(info models.Auth) (*models.LoginResponse, *models.ErrorResponse) {
	var user model.User
	var errResponse models.ErrorResponse

	userRecord := db.Instance.Where("email= ?", info.Username).First(&user)
	if userRecord.Error != nil {
		errResponse.Code = errors.AuthenticationError
		errResponse.Status = errors.AuthenticationErrorStatus
		errResponse.Error = "user not found"
		return nil, &errResponse
	}

	passwordError := user.CheckPassword(info.Password)
	if passwordError != nil {
		errResponse.Code = errors.AuthenticationError
		errResponse.Status = errors.AuthenticationErrorStatus
		errResponse.Error = "password incorrect"
		return nil, &errResponse
	}

	response, err := jwtAuth.GenerateTokens(info.Username)
	if err != nil {
		log.Println(err.Error())
		errResponse.Code = errors.InternalServerError
		errResponse.Status = errors.InternalServerErrorStatus
		errResponse.Error = err.Error()
		return nil, &errResponse
	}

	return response, nil
}

func Register(info models.Register) *models.ErrorResponse {
	newUser := model.User{Email: info.Username, Fullname: info.FullName, Password: info.Password}
	var errResponse models.ErrorResponse

	if err := newUser.HashPassword(newUser.Password); err != nil {
		log.Println(err.Error())
		errResponse.Code = errors.InternalServerError
		errResponse.Status = errors.InternalServerErrorStatus
		errResponse.Error = err.Error()
		return &errResponse
	}

	if err := db.Instance.Create(&newUser).Error; err != nil {
		log.Println(err.Error())
		errResponse.Code = errors.InternalServerError
		errResponse.Status = errors.InternalServerErrorStatus
		errResponse.Error = err.Error()
		return &errResponse
	}

	return nil
}

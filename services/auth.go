package services

import (
	"context"
	"log"

	"github.com/muhammetandic/go-backend/main/core/models"
	"github.com/muhammetandic/go-backend/main/db"
	"github.com/muhammetandic/go-backend/main/db/model"
	"github.com/muhammetandic/go-backend/main/db/repository"
	"github.com/muhammetandic/go-backend/main/services/jwtAuth"
	"github.com/muhammetandic/go-backend/main/utils/helpers"
)

func Login(info models.Auth) (*models.LoginResponse, *models.ErrorResponse) {
	ctx := context.Background()
	userRepo := repository.NewUserRepo()

	user := userRepo.Get(&model.User{Email: info.Username}, ctx)

	if user == nil {
		errResponse := helpers.StatusUnauthenticated("user not found")
		return nil, &errResponse
	}

	passwordError := user.CheckPassword(info.Password)
	if passwordError != nil {
		errResponse := helpers.StatusUnauthenticated("password incorrect")
		return nil, &errResponse
	}

	response, err := jwtAuth.GenerateTokens(info.Username)
	if err != nil {
		log.Println(err.Error())
		errResponse := helpers.StatusInternalServerError(err.Error())
		return nil, &errResponse
	}

	user.RefreshToken = response.RefreshToken
	err = userRepo.Update(int(user.ID), user, ctx)
	if err != nil {
		log.Println(err.Error())
		errResponse := helpers.StatusInternalServerError(err.Error())
		return nil, &errResponse
	}

	return response, nil
}

func Register(info models.Register) *models.ErrorResponse {
	newUser := model.User{Email: info.Username, Fullname: info.FullName, Password: info.Password}

	if err := newUser.HashPassword(newUser.Password); err != nil {
		log.Println(err.Error())
		errResponse := helpers.StatusInternalServerError(err.Error())
		return &errResponse
	}

	if err := db.Instance.Create(&newUser).Error; err != nil {
		log.Println(err.Error())
		errResponse := helpers.StatusInternalServerError(err.Error())
		return &errResponse
	}

	return nil
}

func RefreshToken(token string) (*models.LoginResponse, *models.ErrorResponse) {
	ctx := context.Background()
	userRepo := repository.NewUserRepo()

	jwtUser, err := jwtAuth.ValidateToken(token)
	if err != nil {
		errResponse := helpers.StatusInvalidated(err.Error())
		return nil, &errResponse
	}
	user := userRepo.Get(&model.User{Email: jwtUser.Username}, ctx)
	if user == nil {
		errResponse := helpers.StatusUnauthenticated("user not found")
		return nil, &errResponse
	}

	if user.RefreshToken != token {
		errResponse := helpers.StatusUnauthenticated("token is invalid")
		return nil, &errResponse
	}

	tokens, err := jwtAuth.GenerateTokens(jwtUser.Username)
	if err != nil {
		errResponse := helpers.StatusInternalServerError(err.Error())
		return nil, &errResponse
	}

	user.RefreshToken = tokens.RefreshToken
	err = userRepo.Update(int(user.ID), user, ctx)
	if err != nil {
		log.Println(err.Error())
		errResponse := helpers.StatusInternalServerError(err.Error())
		return nil, &errResponse
	}

	return tokens, nil
}

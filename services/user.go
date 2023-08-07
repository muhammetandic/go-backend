package services

import (
	"github.com/muhammetandic/go-backend/main/db/model"
	"github.com/muhammetandic/go-backend/main/db/repository"
)

type UserService struct {
	*Service[model.User]
}

func NewUserService() *UserService {
	return &UserService{NewService[model.User](repository.NewUserRepo().Repository)}
}

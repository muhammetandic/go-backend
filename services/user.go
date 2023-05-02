package services

import (
	"github.com/muhammetandic/go-backend/main/db"
	"github.com/muhammetandic/go-backend/main/db/model"
	"github.com/muhammetandic/go-backend/main/db/repository"
)

func UserService() *Service[model.User] {
	userRepo := repository.NewRepository[model.User](db.Instance)
	service := NewService(userRepo)
	return service
}

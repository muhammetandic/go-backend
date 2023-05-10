package services

import (
	"github.com/muhammetandic/go-backend/main/db"
	"github.com/muhammetandic/go-backend/main/db/model"
	"github.com/muhammetandic/go-backend/main/db/repository"
)

func TodoService() *Service[model.Todo] {
	todoRepo := repository.NewRepository[model.Todo](db.Instance)
	service := NewService(todoRepo)
	return service
}

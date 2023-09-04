package services

import (
	"github.com/muhammetandic/go-backend/main/db/model"
	"github.com/muhammetandic/go-backend/main/db/repository"
)

type TodoService struct {
	*Service[model.Todo]
}

func NewTodoService() *TodoService {
	return &TodoService{NewService[model.Todo](repository.NewTodoRepo())}
}

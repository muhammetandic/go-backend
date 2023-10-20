package repository

import (
	"context"

	"github.com/muhammetandic/go-backend/main/db"
	"github.com/muhammetandic/go-backend/main/db/model"
)

type TodoRepo struct {
	*Repository[model.Todo]
}

func NewTodoRepo() *TodoRepo {
	repo := NewRepository[model.Todo](db.Instance)
	repository := repo.CreateRepository(context.Background())
	return &TodoRepo{repository}
}

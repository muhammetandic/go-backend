package repository

import (
	"github.com/muhammetandic/go-backend/main/db"
	"github.com/muhammetandic/go-backend/main/db/model"
)

type TodoRepo struct {
	*Repository[model.Todo]
}

func NewTodoRepo() *TodoRepo {
	return &TodoRepo{NewRepository[model.Todo](db.Instance)}
}

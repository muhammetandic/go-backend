package repository

import (
	"context"

	"github.com/muhammetandic/go-backend/main/db"
	"github.com/muhammetandic/go-backend/main/db/model"
)

type PrivilegeRepo struct {
	*Repository[model.Privilege]
}

func NewPrivilegeRepo() *PrivilegeRepo {
	repo := NewRepository[model.Privilege](db.Instance)
	repository := repo.CreateRepository(context.Background())
	return &PrivilegeRepo{repository}
}

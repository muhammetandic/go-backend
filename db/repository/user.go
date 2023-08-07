package repository

import (
	"github.com/muhammetandic/go-backend/main/db"
	"github.com/muhammetandic/go-backend/main/db/model"
)

type UserRepo struct {
	*Repository[model.User]
}

func NewUserRepo() *UserRepo {
	return &UserRepo{NewRepository[model.User](db.Instance)}
}

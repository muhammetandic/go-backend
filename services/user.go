package services

import (
	"github.com/muhammetandic/go-backend/main/db"
	"github.com/muhammetandic/go-backend/main/db/model"
	"github.com/muhammetandic/go-backend/main/db/repository"
)

func User() *repository.Repository[model.User] {
	repo := repository.NewRepository[model.User](db.Instance)
	return repo
}

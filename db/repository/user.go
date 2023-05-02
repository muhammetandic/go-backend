package repository

import (
	"gorm.io/gorm"

	"github.com/muhammetandic/go-backend/main/db/model"
)

func UserRepo(db *gorm.DB) *Repository[model.User] {
	repo := NewRepository[model.User](db)
	return repo
}

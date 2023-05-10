package repository

import (
	"gorm.io/gorm"

	"github.com/muhammetandic/go-backend/main/db/model"
)

func TodoRepo(db *gorm.DB) *Repository[model.Todo] {
	repo := NewRepository[model.Todo](db)
	return repo
}

package repository

import (
	"context"

	"github.com/muhammetandic/go-backend/main/db"
	"github.com/muhammetandic/go-backend/main/db/model"
)

type UserRepo struct {
	*Repository[model.User]
}

func NewUserRepo() *UserRepo {
	return &UserRepo{NewRepository[model.User](db.Instance)}
}

func (r *UserRepo) GetAll(ctx context.Context) (*[]model.User, error) {
	var users []model.User
	err := r.db.Preload("Roles.Role").WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *UserRepo) GetById(id int, ctx context.Context) (*model.User, error) {
	var user model.User
	err := r.db.Preload("Roles.Role").WithContext(ctx).Model(&user).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

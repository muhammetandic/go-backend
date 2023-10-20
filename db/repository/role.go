package repository

import (
	"context"

	"github.com/muhammetandic/go-backend/main/db"
	"github.com/muhammetandic/go-backend/main/db/model"
)

type RoleRepo struct {
	*Repository[model.Role]
}

func NewRoleRepo() *RoleRepo {
	repo := NewRepository[model.Role](db.Instance)
	repository := repo.CreateRepository(context.Background())
	return &RoleRepo{repository}
}

func (r *RoleRepo) GetAll(ctx context.Context) (*[]model.Role, error) {
	var roles []model.Role
	err := r.db.Preload("Privileges").WithContext(ctx).Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return &roles, nil
}

func (r *RoleRepo) GetById(id int, ctx context.Context) (*model.Role, error) {
	var role model.Role
	err := r.db.Preload("Privileges").WithContext(ctx).Model(&role).Where("id = ?", id).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

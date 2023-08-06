package services

import (
	"github.com/muhammetandic/go-backend/main/db/model"
	"github.com/muhammetandic/go-backend/main/db/repository"
)

type RoleService struct {
	*Service[model.Role]
}

func NewRoleService() *RoleService {
	return &RoleService{NewService[model.Role](repository.NewRoleRepo().Repository)}
}

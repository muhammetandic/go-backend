package services

import (
	"context"

	"github.com/muhammetandic/go-backend/main/db"
	"github.com/muhammetandic/go-backend/main/db/model"
	"github.com/muhammetandic/go-backend/main/db/repository"
)

type UserService struct {
	*Service[model.User]
}

func NewUserService() *UserService {
	return &UserService{NewService[model.User](repository.NewUserRepo())}
}

func (us *UserService) AddRole(role *model.UserToRoleDto, ctx context.Context) (model.UserToRole, error) {
	repo := repository.NewRepository[model.UserToRole](db.Instance)
	data := model.UserToRole{UserID: role.UserID, RoleID: role.RoleID}
	entity, err := repo.Add(&data, ctx)
	return *entity, err
}

func (us *UserService) RemoveRole(id int, ctx context.Context) error {
	repo := repository.NewRepository[model.UserToRole](db.Instance)
	err := repo.Delete(id, ctx)
	return err
}

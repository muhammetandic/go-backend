package services

import (
	"github.com/muhammetandic/go-backend/main/db/model"
	"github.com/muhammetandic/go-backend/main/db/repository"
)

type PrivilegeService struct {
	*Service[model.Privilege]
}

func NewPrivilegeService() *PrivilegeService {
	return &PrivilegeService{NewService[model.Privilege](repository.NewPrivilegeRepo())}
}

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
	return &PrivilegeRepo{(NewRepository[model.Privilege](db.Instance)).CreateRepository(context.Background())}
}

package services

import (
	"context"

	"github.com/muhammetandic/go-backend/main/db/repository"
)

type Service[T any] struct {
	repo *repository.Repository[T]
}

func NewService[T any](repo *repository.Repository[T]) *Service[T] {
	return &Service[T]{
		repo: repo,
	}
}

func (service *Service[T]) Add(entity *T, ctx context.Context) error {
	return service.repo.Add(entity, ctx)
}

func (service *Service[T]) GetAll(ctx context.Context) (*[]T, error) {
	entities, err := service.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return entities, nil
}

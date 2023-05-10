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

func (service *Service[T]) Add(entity *T, ctx context.Context) (*T, error) {
	return service.repo.Add(entity, ctx)
}

func (service *Service[T]) GetAll(ctx context.Context) (*[]T, error) {
	entities, err := service.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (service *Service[T]) Get(id int, ctx context.Context) (*T, error) {
	entity, err := service.repo.GetById(id, ctx)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (service *Service[T]) Update(id int, entity *T, ctx context.Context) error {
	err := service.repo.Update(id, entity, ctx)
	return err
}

func (service *Service[T]) Delete(id int, ctx context.Context) error {
	err := service.repo.Delete(id, ctx)
	return err
}

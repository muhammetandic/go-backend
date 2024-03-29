package repository

import (
	"context"

	"gorm.io/gorm"
)

type IRepository[T any] interface {
	CreateRepository(ctx context.Context) *Repository[T]
	Add(entity *T, ctx context.Context) (*T, error)
	AddAll(entities []*T, ctx context.Context) error
	GetAll(ctx context.Context) (*[]T, error)
	GetById(id int, ctx context.Context) (*T, error)
	Delete(id int, ctx context.Context) error
	Get(params *T, ctx context.Context) *T
	GetWithRelated(params *T, relate string, ctx context.Context) *T
	Update(id int, entity *T, ctx context.Context) error
	Where(params *T, ctx context.Context) (*[]T, error)
}

type Repository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) IRepository[T] {
	return &Repository[T]{
		db: db,
	}
}

func (f *Repository[T]) CreateRepository(ctx context.Context) *Repository[T] {
	return &Repository[T]{
		db: f.db.WithContext(ctx),
	}
}

func (r *Repository[T]) Add(entity *T, ctx context.Context) (*T, error) {
	err := r.db.WithContext(ctx).Create(&entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *Repository[T]) AddAll(entities []*T, ctx context.Context) error {
	return r.db.WithContext(ctx).Create(&entities).Error
}

func (r *Repository[T]) GetAll(ctx context.Context) (*[]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return &entities, nil
}

func (r *Repository[T]) GetById(id int, ctx context.Context) (*T, error) {
	var entity T
	err := r.db.WithContext(ctx).Model(&entity).Where("id = ?", id).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *Repository[T]) Get(params *T, ctx context.Context) *T {
	var entity T
	err := r.db.WithContext(ctx).Where(&params).First(&entity).Error
	if err != nil {
		return nil
	}
	return &entity
}

func (r *Repository[T]) GetWithRelated(params *T, relate string, ctx context.Context) *T {
	var entity T
	err := r.db.Preload(relate).WithContext(ctx).Where(&params).First(&entity).Error
	if err != nil {
		return nil
	}
	return &entity
}

func (r *Repository[T]) Where(params *T, ctx context.Context) (*[]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Where(&params).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return &entities, nil
}

func (r *Repository[T]) Update(id int, entity *T, ctx context.Context) error {
	var savedEntity T
	err := r.db.WithContext(ctx).Model(&savedEntity).Where("id = ?", id).First(&savedEntity).Error
	if err != nil {
		return err
	}
	return r.db.WithContext(ctx).Save(&entity).Error
}

func (r *Repository[T]) UpdateAll(entities []*T, ctx context.Context) error {
	return r.db.WithContext(ctx).Save(&entities).Error
}

func (r *Repository[T]) Delete(id int, ctx context.Context) error {
	var entity *T
	err := r.db.WithContext(ctx).First(&entity, id).Error
	if err != nil {
		return err
	}
	return r.db.WithContext(ctx).Delete(&entity).Error
}

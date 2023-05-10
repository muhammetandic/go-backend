package repository

import (
	"context"

	"gorm.io/gorm"
)

type Repository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) *Repository[T] {
	return &Repository[T]{
		db: db,
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
	return r.db.WithContext(ctx).Model(&entity).Updates(&entity).Error
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

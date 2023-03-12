package repository

import (
	"context"

	"gorm.io/gorm"
)

type repository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) *repository[T] {
	return &repository[T]{
		db: db,
	}
}

func (r *repository[T]) Add(entity *T, ctx context.Context) error {
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *repository[T]) AddAll(entities []*T, ctx context.Context) error {
	return r.db.WithContext(ctx).Create(&entities).Error
}

func (r *repository[T]) GetAll(ctx context.Context) (*[]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return &entities, nil
}

func (r *repository[T]) GetById(id int, ctx context.Context) (*T, error) {
	var entity T
	err := r.db.WithContext(ctx).Model(&entity).Where("id= ? AND is_active= ?", id, true).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *repository[T]) Get(params *T, ctx context.Context) *T {
	var entity T
	r.db.WithContext(ctx).Where(&params).First(&entity)
	return &entity
}

func (r *repository[T]) Where(params *T, ctx context.Context) (*[]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Where(&params).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return &entities, nil
}

func (r *repository[T]) Update(entity *T, ctx context.Context) error {
	return r.db.WithContext(ctx).Save(&entity).Error
}

func (r *repository[T]) UpdateAll(entities []*T, ctx context.Context) error {
	return r.db.WithContext(ctx).Save(&entities).Error
}

func (r *repository[T]) Delete(id int, ctx context.Context) error {
	var entity *T
	// err := r.db.WithContext(ctx).Model(&entity).Where("id = ? AND is_active = ?", id, true).FirstOrInit(&entity).Error
	// if err != nil {
	// 	return err
	// }
	return r.db.WithContext(ctx).First(&entity).UpdateColumn("is_active", false).Error
}

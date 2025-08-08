package repo

import (
	"context"

	"gorm.io/gorm"
)

type IGenericRepo[T any] interface {
	Create(context.Context, *T) error
	Delete(context.Context, uint) error
	GetList(context.Context) (*[]T, error)
	GetById(context.Context, uint) (*T, error)
}

type GenericRepo[T any] struct {
	db *gorm.DB
}

func (g *GenericRepo[T]) Create(ctx context.Context, record *T) error {
	err := gorm.G[T](g.db).Create(ctx, record)
	if err != nil {
		return err
	}
	return nil
}

func (g *GenericRepo[T]) Delete(ctx context.Context, id uint) error {
	_, err := gorm.G[T](g.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (g *GenericRepo[T]) GetList(ctx context.Context) (*[]T, error) {
	records, err := gorm.G[T](g.db).Find(ctx)
	if err != nil {
		return nil, err
	}
	return &records, nil
}

func (g *GenericRepo[T]) GetById(ctx context.Context, id uint) (*T, error) {
	record, err := gorm.G[T](g.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func NewGenericRepo[T any](db *gorm.DB) *GenericRepo[T] {
	return &GenericRepo[T]{
		db,
	}
}

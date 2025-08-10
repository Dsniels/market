package repo

import (
	"context"

	"gorm.io/gorm"
)

type IGeneric[T any] interface {
	Create(context.Context, *T) error
	Update(t *T) (*T, error)
	Delete(context.Context, uint) error
	GetList(context.Context) (*[]T, error)
	GetById(context.Context, uint) (*T, error)
}

type Generic[T any] struct {
	db *gorm.DB
}

func (g *Generic[T]) Create(ctx context.Context, record *T) error {
	err := gorm.G[T](g.db).Create(ctx, record)
	if err != nil {
		return err
	}
	return nil
}

func (g *Generic[T]) Delete(ctx context.Context, id uint) error {
	_, err := gorm.G[T](g.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (g *Generic[T]) GetList(ctx context.Context) (*[]T, error) {
	records, err := gorm.G[T](g.db).Find(ctx)
	if err != nil {
		return nil, err
	}
	return &records, nil
}

func (g *Generic[T]) GetById(ctx context.Context, id uint) (*T, error) {
	record, err := gorm.G[T](g.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (p *Generic[T]) Update(t *T) (*T, error) {
	err := p.db.Save(t).Error
	if err != nil {
		return nil, err
	}
	return t, err
}
func NewGeneric[T any](db *gorm.DB) *Generic[T] {
	return &Generic[T]{
		db,
	}
}

package services

import (
	"context"

	"github.com/dsniels/market/core/types"
	"github.com/dsniels/market/internal/repo"
)

type FormaPago struct {
	repo repo.IFormaPago
}

type IFormaPago interface {
	CreateFormaPago(ctx context.Context, categoria *types.FormaPago) error
	GetFormaPagos(ctx context.Context) (*[]types.FormaPago, error)
	GetFormaPago(ctx context.Context, id uint) (*types.FormaPago, error)
	DeleteFormaPago(ctx context.Context, id uint) error
}

func (c *FormaPago) CreateFormaPago(ctx context.Context, formaPago *types.FormaPago) error {
	err := c.repo.Create(ctx, formaPago)
	return err
}

func (c *FormaPago) GetFormaPagos(ctx context.Context) (*[]types.FormaPago, error) {
	list, err := c.repo.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return list, err
}

func (c *FormaPago) GetFormaPago(ctx context.Context, id uint) (*types.FormaPago, error) {
	categoria, err := c.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return categoria, nil
}

func (c *FormaPago) DeleteFormaPago(ctx context.Context, id uint) error {
	_, err := c.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func NewFormaPago(repo repo.IFormaPago) *FormaPago {
	return &FormaPago{
		repo: repo,
	}
}

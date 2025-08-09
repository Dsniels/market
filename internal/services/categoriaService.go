package services

import (
	"context"

	"github.com/dsniels/market/core/types"
	"github.com/dsniels/market/internal/repo"
)

type Categoria struct {
	repo repo.ICategoria
}

type ICategoria interface {
	CreateCategoria(ctx context.Context, categoria *types.Categoria) error
	GetCategorias(ctx context.Context) (*[]types.Categoria, error)
	GetCategoria(ctx context.Context, id uint) (*types.Categoria, error)
	DeleteCategoria(ctx context.Context, id uint) error
}

func (c *Categoria) CreateCategoria(ctx context.Context, categoria *types.Categoria) error {
	err := c.repo.Create(ctx, categoria)
	return err
}

func (c *Categoria) GetCategorias(ctx context.Context) (*[]types.Categoria, error) {
	list, err := c.repo.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return list, err
}

func (c *Categoria) GetCategoria(ctx context.Context, id uint) (*types.Categoria, error) {
	categoria, err := c.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return categoria, nil
}

func (c *Categoria) DeleteCategoria(ctx context.Context, id uint) error {
	_, err := c.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func NewCategoria(repo repo.ICategoria) *Categoria {
	return &Categoria{
		repo: repo,
	}
}

package services

import (
	"context"

	"github.com/dsniels/market/core/types"
	"github.com/dsniels/market/internal/repo"
)

type ProductoComp struct {
	repo repo.IProductoComp
}

type IProductoComp interface {
	CreateProductoComp(ctx context.Context, categoria *types.ProductoCompuesto) error
	GetProductoComps(ctx context.Context) (*[]types.ProductoCompuesto, error)
	GetProductoComp(ctx context.Context, id uint) (*types.ProductoCompuesto, error)
	DeleteProductoComp(ctx context.Context, id uint) error
}

func (c *ProductoComp) CreateProducto(ctx context.Context, prod *types.ProductoCompuesto) error {
	err := c.repo.Create(ctx, prod)
	return err
}

func (c *ProductoComp) GetProductos(ctx context.Context) (*[]types.ProductoCompuesto, error) {
	list, err := c.repo.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return list, err
}

func (c *ProductoComp) GetProducto(ctx context.Context, id uint) (*types.ProductoCompuesto, error) {
	categoria, err := c.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return categoria, nil
}

func (c *ProductoComp) DeleteProducto(ctx context.Context, id uint) error {
	_, err := c.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func NewProductoComp(repo repo.IProductoComp) *ProductoComp {
	return &ProductoComp{
		repo: repo,
	}
}

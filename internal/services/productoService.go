package services

import (
	"context"
	"fmt"

	"github.com/dsniels/market/core/types"
	"github.com/dsniels/market/internal/repo"
)

type Producto struct {
	repo         repo.IProducto
	categoriaSvc Categoria
}

type IProducto interface {
	CreateProducto(ctx context.Context, prod *types.Producto) error
	GetProductos(ctx context.Context) (*[]types.Producto, error)
	GetProducto(ctx context.Context, id uint) (*types.Producto, error)
	DeleteProducto(ctx context.Context, id uint) error
}

func (p *Producto) CreateProducto(ctx context.Context, prod *types.Producto) error {
	categoria, err := p.categoriaSvc.GetCategoria(ctx, prod.CategoriaID)
	if err != nil {
		return err
	}
	if categoria == nil {
		return fmt.Errorf("la categoria no existe")
	}
	err = p.repo.Create(ctx, prod)
	return err
}

func (c *Producto) GetProductos(ctx context.Context) (*[]types.Producto, error) {
	list, err := c.repo.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return list, err
}

func (c *Producto) GetProducto(ctx context.Context, id uint) (*types.Producto, error) {
	categoria, err := c.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return categoria, nil
}

func (c *Producto) DeleteProducto(ctx context.Context, id uint) error {
	_, err := c.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func NewProducto(repo repo.IProducto) *Producto {
	return &Producto{
		repo: repo,
	}
}

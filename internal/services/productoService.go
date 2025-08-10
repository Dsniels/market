package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/dsniels/market/core/types"
	"github.com/dsniels/market/internal/repo"
)

type Producto struct {
	repo         repo.IProducto
	categoriaSvc ICategoria
}

type IProducto interface {
	CreateProducto(ctx context.Context, prod *types.Producto) error
	GetProductos(ctx context.Context) (*[]types.Producto, error)
	GetProducto(ctx context.Context, id uint) (*types.Producto, error)
	DeleteProducto(ctx context.Context, id uint) error
}

func (p *Producto) CreateProducto(ctx context.Context, prod *types.Producto) error {
	_, err := p.categoriaSvc.GetCategoria(ctx, prod.CategoriaID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return fmt.Errorf("categoria not found")
		}
		return err
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
	prod, err := c.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return prod, nil
}

func (c *Producto) DeleteProducto(ctx context.Context, id uint) error {
	err := c.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func NewProducto(repo repo.IProducto, categoriaSvc ICategoria) *Producto {
	return &Producto{
		repo:         repo,
		categoriaSvc: categoriaSvc,
	}
}

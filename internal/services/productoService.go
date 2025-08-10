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
	UpdateProducto(ctx context.Context, producto *types.Producto, id uint) (*types.Producto, error)
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

func (p *Producto) GetProductos(ctx context.Context) (*[]types.Producto, error) {
	list, err := p.repo.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return list, err
}

func (p *Producto) GetProducto(ctx context.Context, id uint) (*types.Producto, error) {
	prod, err := p.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return prod, nil
}

func (p *Producto) DeleteProducto(ctx context.Context, id uint) error {
	err := p.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *Producto) UpdateProducto(ctx context.Context, producto *types.Producto, id uint) (*types.Producto, error) {
	_, err := p.GetProducto(ctx, id)
	if err != nil {
		return nil, err
	}
	producto.ID = id

	newProducto, err := p.repo.Update(producto)
	if err != nil {
		return nil, err
	}

	return newProducto, nil

}

func NewProducto(repo repo.IProducto, categoriaSvc ICategoria) *Producto {
	return &Producto{
		repo:         repo,
		categoriaSvc: categoriaSvc,
	}
}

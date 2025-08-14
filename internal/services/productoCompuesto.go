package services

import (
	"context"

	"github.com/dsniels/market/core/dto"
	"github.com/dsniels/market/core/types"
	"github.com/dsniels/market/internal/repo"
)

type ProductoComp struct {
	productoSvc IProducto
	repo        repo.IProductoComp
}

type IProductoComp interface {
	CreateProducto(ctx context.Context, producto *dto.ProductoComp) error
	GetProductos(ctx context.Context) (*[]types.ProductoCompuesto, error)
	GetProducto(ctx context.Context, id uint) (*types.ProductoCompuesto, error)
	DeleteProducto(ctx context.Context, id uint) error
}

func (p *ProductoComp) CreateProducto(ctx context.Context, prod *dto.ProductoComp) error {
	prodComp, err := p.productoSvc.GetProducto(ctx, prod.ProductoComponenteID)
	if err != nil {
		return err
	}

	prodPrincipal, err := p.productoSvc.GetProducto(ctx, prod.ProductoPrincipalID)
	if err != nil {
		return err
	}

	comp := new(types.ProductoCompuesto)
	comp.ProductoComponenteID = prodComp.ID
	comp.ProductoComponente = *prodComp
	comp.ProductoPrincipalID = prodPrincipal.ID
	comp.ProductoPrincipal = *prodPrincipal

	err = p.repo.Create(ctx, comp)
	if err != nil {
		return err
	}
	prod.ID = comp.ID

	return nil
}

func (c *ProductoComp) GetProductos(ctx context.Context) (*[]types.ProductoCompuesto, error) {
	list, err := c.repo.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return list, err
}

func (p *ProductoComp) GetProducto(ctx context.Context, id uint) (*types.ProductoCompuesto, error) {

	obj, err := p.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return obj, nil

}

func (c *ProductoComp) DeleteProducto(ctx context.Context, id uint) error {
	err := c.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func NewProductoComp(repo repo.IProductoComp, productoSvc IProducto) *ProductoComp {
	return &ProductoComp{
		productoSvc: productoSvc,
		repo:        repo,
	}
}

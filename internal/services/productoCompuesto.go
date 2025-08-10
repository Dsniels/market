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
	CreateProducto(ctx context.Context, producto *types.ProductoCompuesto) error
	GetProductos(ctx context.Context) (*[]types.ProductoCompuesto, error)
	GetProducto(ctx context.Context, id uint) (*dto.ProductoComp, error)
	DeleteProducto(ctx context.Context, id uint) error
}

func (p *ProductoComp) CreateProducto(ctx context.Context, prod *types.ProductoCompuesto) error {
	_, err := p.productoSvc.GetProducto(ctx, prod.ProductoComponenteID)
	if err != nil {
		return err
	}

	_, err = p.productoSvc.GetProducto(ctx, prod.ProductoPrincipalID)
	if err != nil {
		return err
	}

	err = p.repo.Create(ctx, prod)
	if err != nil {
		return err
	}

	return nil
}

func (c *ProductoComp) GetProductos(ctx context.Context) (*[]types.ProductoCompuesto, error) {
	list, err := c.repo.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return list, err
}

func (p *ProductoComp) GetProducto(ctx context.Context, id uint) (*dto.ProductoComp, error) {
	dto := new(dto.ProductoComp)
	head, err := p.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	principal, err := p.productoSvc.GetProducto(ctx, head.ProductoPrincipalID)
	if err != nil {
		return nil, err
	}

	comp, err := p.productoSvc.GetProducto(ctx, head.ProductoComponenteID)
	if err != nil {
		return nil, err
	}
	dto.ProductoCompuesto = comp
	dto.ProductoPrincipal = principal
	return dto, nil

}

func (c *ProductoComp) DeleteProducto(ctx context.Context, id uint) error {
	err := c.repo.Delete(ctx, id)
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

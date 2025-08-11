package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dsniels/market/core/dto"
	"github.com/dsniels/market/internal/services"
	"github.com/dsniels/market/pkg"
	"gorm.io/gorm"
)

type ProductoComp struct {
	productoSvc services.IProductoComp
}

// GetProductsHandler godoc
// @Summary Get all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} pkg.Body[[]types.ProductoCompuesto]
// @Router /api/ProductComp/GetAll [get]
func (p *ProductoComp) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := p.productoSvc.GetProductos(r.Context())
	if err != nil {
		pkg.BadRequest(err.Error())
	}
	pkg.Response(w, http.StatusOK, products)
}

// GetProductHandler godoc
// @Summary Get a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} pkg.Body[dto.ProductoComp]
// @Router /api/ProductComp/GetByID/{id} [get]
func (p *ProductoComp) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	id := pkg.GetIDFromUrl[uint](r)
	product, err := p.productoSvc.GetProducto(r.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			pkg.NotFound()
		}
		pkg.BadRequest(err.Error())
	}
	pkg.Response(w, http.StatusOK, product)
}

// CreateProductHandler godoc
// @Summary Create a new productComp
// @Tags products
// @Accept json
// @Produce json
// @Param product body dto.ProductoComp true "Product data"
// @Success 200 {object} pkg.Body[types.ProductoCompuesto]
// @Router /api/ProductComp/Create [post]
func (p *ProductoComp) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	product := new(dto.ProductoComp)
	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		pkg.BadRequest(err.Error())
	}
	err = p.productoSvc.CreateProducto(r.Context(), product)

	if err != nil {
		pkg.BadRequest(err.Error())
	}
	pkg.Response(w, http.StatusOK, product)
}

// DeleteProductHandler godoc
// @Summary Delete a productComp by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} pkg.Body[string]
// @Router /api/ProductComp/DeleteByID/{id} [post]
func (p *ProductoComp) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {

	id := pkg.GetIDFromUrl[uint](r)
	_, err := p.productoSvc.GetProducto(r.Context(), id)
	if err != nil {
		pkg.BadRequest(err.Error())
	}

	err = p.productoSvc.DeleteProducto(r.Context(), id)
	if err != nil {
		pkg.PanicException(500, err.Error())
	}

	pkg.Response(w, http.StatusOK, "ok")

}

func NewProductoComp(productoSvc services.IProductoComp) *ProductoComp {
	return &ProductoComp{
		productoSvc: productoSvc,
	}
}

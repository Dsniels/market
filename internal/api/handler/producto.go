package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dsniels/market/core/types"
	"github.com/dsniels/market/internal/services"
	"github.com/dsniels/market/pkg"
	"gorm.io/gorm"
)

type Producto struct {
	productoSvc services.IProducto
	productoCompSvc services.IProductoComp
}

// GetProductsHandler godoc
// @Summary Get all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} pkg.Body[[]types.Producto]
// @Router /api/Product/GetAll [get]
func (p *Producto) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
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
// @Success 200 {object} pkg.Body[types.Producto]
// @Router /api/Product/GetByID/{id} [get]
func (p *Producto) GetProductHandler(w http.ResponseWriter, r *http.Request) {
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

// UpdateProductHandler godoc
// @Summary Update a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body types.Producto true "Product data"
// @Success 200 {object} pkg.Body[types.Producto]
// @Router /api/Product/Update/{id} [post]
func (p *Producto) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	product := new(types.Producto)
	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		pkg.BadRequest(err.Error())
	}
	id := pkg.GetIDFromUrl[uint](r)
	product, err = p.productoSvc.UpdateProducto(r.Context(), product, id)
	if err != nil {
		pkg.BadRequest(err.Error())
	}
	pkg.Response(w, http.StatusOK, product)
}

// CreateProductHandler godoc
// @Summary Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body types.Producto true "Product data"
// @Success 200 {object} pkg.Body[types.Producto]
// @Router /api/Product/Create [post]
func (p *Producto) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	product := new(types.Producto)
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
// @Summary Delete a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} pkg.Body[string]
// @Router /api/Product/DeleteByID/{id} [post]
func (p *Producto) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {

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

func NewProducto(productoSvc services.IProducto) *Producto {
	return &Producto{
		productoSvc: productoSvc,
	}
}

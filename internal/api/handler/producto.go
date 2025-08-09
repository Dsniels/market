package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dsniels/market/core/types"
	"github.com/dsniels/market/internal/services"
	"github.com/dsniels/market/pkg"
)

type Producto struct {
	productoSvc services.IProducto
}

func (p *Producto) GetProductHandler(w http.ResponseWriter, r *http.Request) {

}

func (p *Producto) GetProductsHandler(w http.ResponseWriter, r *http.Request) {

}

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

	pkg.Response(w, http.StatusOK, struct{}{})

}

func NewProducto(productoSvc services.IProducto) *Producto {
	return &Producto{
		productoSvc: productoSvc,
	}
}

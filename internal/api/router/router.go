package router

import (
	"net/http"

	"github.com/dsniels/market/internal/api"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func InitRoutes(a *api.App) http.Handler {
	route := http.NewServeMux()

	route.HandleFunc("GET /Product/GetAll", a.Product.GetProductsHandler)
	route.HandleFunc("GET /Product/GetByID/{id}", a.Product.GetProductHandler)
	route.HandleFunc("POST /Product/Create", a.Product.CreateProductHandler)
	route.HandleFunc("POST /Product/Update/{id}", a.Product.UpdateProductHandler)
	route.HandleFunc("POST /Product/DeleteByID/{id}", a.Product.DeleteProductHandler)

	route.HandleFunc("GET /Sucursal/GetAll", a.Sucursal.GetSucursalsHandler)
	route.HandleFunc("GET /Sucursal/GetByID/{id}", a.Sucursal.GetSucursalHandler)
	route.HandleFunc("POST /Sucursal/Create", a.Sucursal.CreateSucursalHandler)
	route.HandleFunc("POST /Sucursal/Update/{id}", a.Sucursal.UpdateSucursalHandler)
	route.HandleFunc("POST /Sucursal/DeleteByID/{id}", a.Sucursal.DeleteSucursalHandler)

	route.HandleFunc("GET /Categoria/GetByID/{id}", a.Categoria.GetCategoriaHandler)
	route.HandleFunc("GET /Categoria/GetAll", a.Categoria.GetCategoriasHandler)
	route.HandleFunc("POST /Categoria/Create", a.Categoria.CreateCategoriaHandler)
	route.HandleFunc("POST /Categoria/DeleteByID/{id}", a.Categoria.DeleteCategoriaHandler)

	route.HandleFunc("GET /FormaPago/GetByID/{id}", a.FormaPago.GetFormaPagoHandler)
	route.HandleFunc("GET /FormaPago/GetAll", a.FormaPago.GetFormasPagoHandler)
	route.HandleFunc("POST /FormaPago/Create", a.FormaPago.CreateFormaPagoHandler)
	route.HandleFunc("POST /FormaPago/DeleteByID/{id}", a.FormaPago.DeleteFormaPagoHandler)

	route.HandleFunc("GET /ProductComp/GetAll", a.ProductComp.GetProductsHandler)
	route.HandleFunc("GET /ProductComp/GetByID/{id}", a.ProductComp.GetProductHandler)
	route.HandleFunc("POST /ProductComp/Create", a.ProductComp.CreateProductHandler)
	route.HandleFunc("POST /ProductComp/DeleteByID/{id}", a.ProductComp.DeleteProductHandler)

	router := http.NewServeMux()
	router.Handle("/api/", http.StripPrefix("/api", route))
	router.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)
	return router
}

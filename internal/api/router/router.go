package router

import (
	"net/http"

	"github.com/dsniels/market/internal/api"
)

func InitRoutes(a *api.App) http.Handler {

	router := http.NewServeMux()

	router.HandleFunc("GET /Product/GetAll", a.Product.GetProductsHandler)
	router.HandleFunc("GET /Product/GetByID/{id}", a.Product.GetProductHandler)
	router.HandleFunc("POST /Product/", a.Product.CreateProductHandler)
	router.HandleFunc("POST /Product/DeleteByID/{id}", a.Product.DeleteProductHandler)

	return router
}

package api

import (
	"log/slog"

	"github.com/dsniels/market/internal/api/handler"
	"github.com/dsniels/market/internal/database"
	"github.com/dsniels/market/internal/repo"
	"github.com/dsniels/market/internal/services"
)

type App struct {
	Product     *handler.Producto
	ProductComp *handler.ProductoComp
	Categoria   *handler.Categoria
	FormaPago   *handler.FormaPago
}

func NewApp() *App {

	db := database.Connect()
	err := database.Migrate(db)
	if err != nil {
		slog.Error("Migration Error", slog.String("msg", err.Error()))
		panic(err)
	}
	categoriaRepo := repo.NewCategoria(db)
	categoriaSvc := services.NewCategoria(categoriaRepo)
	categoriaHandler := handler.NewCategoria(categoriaSvc)

	productRepo := repo.NewProducto(db)
	productSvc := services.NewProducto(productRepo, categoriaSvc)
	productHandler := handler.NewProducto(productSvc)
	productCompRepo := repo.NewProductoComp(db)
	productCompSvc := services.NewProductoComp(productCompRepo)
	productCompHandler := handler.NewProductoComp(productCompSvc)

	formaPagoRepo := repo.NewFormaPago(db)
	formaPagoSvc := services.NewFormaPago(formaPagoRepo)
	formaPago := handler.NewFormaPago(formaPagoSvc)

	return &App{
		Product:     productHandler,
		Categoria:   categoriaHandler,
		ProductComp: productCompHandler,
		FormaPago:   formaPago,
	}
}

package api

import (
	"log/slog"

	"github.com/dsniels/market/internal/api/handler"
	"github.com/dsniels/market/internal/database"
	"github.com/dsniels/market/internal/repo"
	"github.com/dsniels/market/internal/services"
)

type App struct {
	Product *handler.Producto
}

func NewApp() *App {

	db := database.Connect()
	err := database.Migrate(db)
	if err != nil {
		slog.Error("Migration Error", slog.String("msg", err.Error()))
		panic(err)
	}
	productRepo := repo.NewProducto(db)
	productSvc := services.NewProducto(productRepo)
	productHandler := handler.NewProducto(productSvc)

	return &App{
		Product: productHandler,
	}
}

package database

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/dsniels/market/internal/types"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	var server = os.Getenv("DB_HOST")
	var port = 1433
	var user = os.Getenv("DB_USER")
	var password = os.Getenv("DB_PWD")
	var database = "market"

	url := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database,
	)
	db, err := gorm.Open(sqlserver.Open(url), &gorm.Config{})
	if err != nil {
		slog.Error("Open Database: %v", err.Error())
		panic(err)
	}

	slog.Info("Database connected")
	return db
}

func Migrate(db *gorm.DB) error {
	slog.Info("Starting Migrations")
	return db.AutoMigrate(&types.Sucursal{}, &types.Categoria{}, &types.Producto{}, &types.FormaPago{}, &types.ProductoCompuesto{})

}

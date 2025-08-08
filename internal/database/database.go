package database

import (
	"fmt"
	"log/slog"

	"github.com/dsniels/market/internal/types"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var server = "dbunilearn.database.windows.net"
var port = 1433
var user = "adminlearn"
var password = "UniLearn123"
var database = "market"

func Connect() *gorm.DB {
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

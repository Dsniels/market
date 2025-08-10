package database

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/dsniels/market/core/types"
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
	slog.Info("Connecting to database")
	db := connectDb(url)
	slog.Info("Database connected")
	return db
}

func Migrate(db *gorm.DB) error {
	slog.Info("Migrating database")
	return db.AutoMigrate(&types.Sucursal{}, &types.Categoria{}, &types.Producto{}, &types.FormaPago{}, &types.ProductoCompuesto{})
}

func connectDb(url string) *gorm.DB {

	d, err := sql.Open("sqlserver", url)
	if err != nil {
		log.Fatalln("Cannot connect db", err)
	}

	tries := 10

	for {
		err := d.Ping()
		if err != nil {
			if tries == 0 {
				log.Fatalln("something went wrong", err)
			}
			slog.Warn("Couldnt connect to db", slog.Int("retries left", tries))
			tries--
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}

	db, err := gorm.Open(sqlserver.Open(url))
	if err != nil {
		log.Fatalln("Cannot connect db", err)
	}

	return db
}

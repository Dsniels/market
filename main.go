package main

import (
	"flag"
	"log/slog"
	"net/http"

	"github.com/dsniels/market/internal/api/router"
	"github.com/dsniels/market/internal/database"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "80", "port to run server")
	flag.Parse()

	slog.Info("Starting Server")

	db := database.Connect()
	err := database.Migrate(db)
	if err != nil{
		slog.Error("Migrating database", slog.Any("error",err))
		panic(err)
	}

	router := router.InitRoutes()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	slog.Info("Server running", slog.String("port", port))
	if err := server.ListenAndServe(); err != nil {
		slog.Error("Running server: ")
		panic(err)
	}

}

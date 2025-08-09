package main

import (
	"flag"
	"log/slog"
	"net/http"

	"github.com/dsniels/market/internal/api"
	"github.com/dsniels/market/internal/api/router"
	"github.com/joho/godotenv"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "80", "port to run server")
	flag.Parse()
	godotenv.Load()

	panic(runServer(port))
}

func runServer(port string) error {
	app := api.NewApp()
	r := router.InitRoutes(app)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router.ErrorMiddleware(r),
	}
	slog.Info("Server running", slog.String("port", port))
	if err := server.ListenAndServe(); err != nil {
		slog.Error("Failed to start server: ")
		return err
	}

	return nil

}

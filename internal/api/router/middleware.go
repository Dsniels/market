package router

import (
	"net/http"

	pkg "github.com/dsniels/market/pkg"
)

func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				pkg.HandleException(w, err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

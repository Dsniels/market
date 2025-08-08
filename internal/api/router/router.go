package router

import "net/http"

func InitRoutes() http.Handler{

	router := http.NewServeMux()




	return  router
}
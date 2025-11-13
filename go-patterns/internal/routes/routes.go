package routes

import (
	"myapi/internal/handlers"
	"myapi/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.JsonContentType)
	r.HandleFunc("/api", handlers.IndexHandler).Methods("GET")
	SetupItemRoutes(r)
	SetupCategoryRoutes(r)
	return r
}

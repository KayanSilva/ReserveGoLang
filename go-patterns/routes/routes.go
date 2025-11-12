package routes

import (
	"myapi/internal/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api", handlers.IndexHandler).Methods("GET")
	SetupItemRoutes(r)
	SetupCategoryRoutes(r)
	return r
}

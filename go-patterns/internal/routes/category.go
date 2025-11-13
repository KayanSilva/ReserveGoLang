package routes

import (
	"myapi/internal/handlers"

	"github.com/gorilla/mux"
)

func SetupCategoryRoutes(r *mux.Router) {
	r.HandleFunc("/categorias", handlers.ListCategoriasHandler).Methods("GET")
	r.HandleFunc("/categorias/{id}", handlers.GetCategoriaHandler).Methods("GET")
	r.HandleFunc("/categorias", handlers.CreateCategoriaHandler).Methods("POST")
	r.HandleFunc("/categorias", handlers.UpdateCategoriaHandler).Methods("PUT")
	r.HandleFunc("/categorias", handlers.DeleteCategoriaHandler).Methods("DELETE")
}

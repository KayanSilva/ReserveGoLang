package routes

import (
	"myapi/internal/handlers"

	"github.com/gorilla/mux"
)

func SetupItemRoutes(r *mux.Router) {
	r.HandleFunc("/api/itens", handlers.ListItensHandler).Methods("GET")
	r.HandleFunc("/api/itens/{id}", handlers.GetItenHandler).Methods("GET")
	r.HandleFunc("/api/itens/code/{code}", handlers.GetItenByCodigoHandler).Methods("GET")
	r.HandleFunc("/api/itens", handlers.CreateItenHandler).Methods("POST")
	r.HandleFunc("/api/itens", handlers.UpdateItenHandler).Methods("PUT")
	r.HandleFunc("/api/itens/{id}", handlers.DeleteItenHandler).Methods("DELETE")
}

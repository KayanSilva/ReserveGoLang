package main

import (
	"log"
	"myapi/internal/config"
	"myapi/internal/routes"
	"net/http"
)

func main() {
	config.Connect()
	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", routes.SetupRoutes()))
}

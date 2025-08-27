package main

import (
	"net/http"

	"github.com/KayanSilva/ReserveGoLang/web-api/routes"
	_ "github.com/lib/pq"
)

func main() {
	routes.SetupRoutes()
	http.ListenAndServe(":8080", nil)
}

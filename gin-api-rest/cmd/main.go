package main

import (
	"github.com/KayanSilva/ReserveGoLang/gin-api-rest/database"
	"github.com/KayanSilva/ReserveGoLang/gin-api-rest/routes"
)

func main() {
	database.Connect()

	routes.HandleRoutes()
}

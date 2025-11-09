package main

import (
	"github.com/KayanSilva/ReserveGoLang/api-rest/database"
	"github.com/KayanSilva/ReserveGoLang/api-rest/routes"
)

func main() {
	database.Connect()
	routes.HandleRequest()
}

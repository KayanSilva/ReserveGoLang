package routes

import (
	"net/http"

	"github.com/KayanSilva/ReserveGoLang/web-api/controllers"
)

func SetupRoutes() {
	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/new", controllers.CreateProductPageHandler)
	http.HandleFunc("/insert", controllers.InsertNewProductHandler)
	http.HandleFunc("/delete", controllers.DeleteProductHandler)
	http.HandleFunc("/edit", controllers.EditProductPageHandler)
	http.HandleFunc("/update", controllers.UpdateProductHandler)
}

package routes

import (
	"log"
	"net/http"

	"github.com/KayanSilva/ReserveGoLang/api-rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalities", controllers.AllPersonalities).Methods("GET")
	r.HandleFunc("/personalities/{id}", controllers.GetPersonalityById).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

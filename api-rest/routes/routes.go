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
	r.HandleFunc("/personalities", controllers.GetPersonalities).Methods("GET")
	r.HandleFunc("/personalities/{id}", controllers.GetPersonalityById).Methods("GET")
	r.HandleFunc("/personalities", controllers.NewPersonality).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}

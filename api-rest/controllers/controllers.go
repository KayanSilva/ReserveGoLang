package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KayanSilva/ReserveGoLang/api-rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Home Page!"))
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}

func GetPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	for _, personality := range models.Personalities {
		if int(personality.ID) == id {
			json.NewEncoder(w).Encode(personality)
			return
		}
	}
	http.Error(w, "Personality not found", http.StatusNotFound)
}

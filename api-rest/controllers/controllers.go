package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KayanSilva/ReserveGoLang/api-rest/database"
	"github.com/KayanSilva/ReserveGoLang/api-rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Home Page!"))
}

func GetPersonalities(w http.ResponseWriter, r *http.Request) {
	var personalities []models.Personality
	database.DB.Find(&personalities)
	json.NewEncoder(w).Encode(&personalities)
}

func GetPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var personality models.Personality

	database.DB.First(&personality, id)
	if personality.ID == 0 {
		http.Error(w, "Personality not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(personality)
}

func NewPersonality(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	json.NewDecoder(r.Body).Decode(&personality)
	database.DB.Create(&personality)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(personality)
}

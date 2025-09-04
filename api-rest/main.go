package main

import (
	"github.com/KayanSilva/ReserveGoLang/api-rest/models"
	"github.com/KayanSilva/ReserveGoLang/api-rest/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{ID: 1, Name: "Adventurous", History: "Loves exploring new places and trying new activities."},
		{ID: 2, Name: "Creative", History: "Enjoys artistic pursuits and thinking outside the box."},
		{ID: 3, Name: "Analytical", History: "Prefers logical reasoning and data-driven decisions."},
	}

	routes.HandleRequest()
}

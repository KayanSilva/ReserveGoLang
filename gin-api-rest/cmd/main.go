package main

import (
	"github.com/KayanSilva/ReserveGoLang/gin-api-rest/database"
	"github.com/KayanSilva/ReserveGoLang/gin-api-rest/models"
	"github.com/KayanSilva/ReserveGoLang/gin-api-rest/routes"
)

func main() {
	database.Connect()
	models.Students = []models.Student{
		{Name: "John Doe", CPF: "12345678901", RG: "123456789"},
		{Name: "Jane Smith", CPF: "10987654321", RG: "987654321"},
		{Name: "Alice Johnson", CPF: "11223344556", RG: "556677889"},
		{Name: "Bob Brown", CPF: "66778899000", RG: "0011223344"},
	}
	routes.HandleRoutes()
}

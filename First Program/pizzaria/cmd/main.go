package main

import (
	"fmt"
	"pizzaria/internal/data"
	"pizzaria/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	nomePizzaria := "Pizzaria Go"
	instagram, telefone := "@pizzaria_go", 1199711
	fmt.Println(nomePizzaria, instagram, telefone)
	data.LoadPizzas()

	router := gin.Default()
	router.GET("/pizzas", handler.GetPizzas)
	router.GET("/pizzas/:id", handler.GetPizzasById)
	router.POST("pizzas", handler.PostPizza)
	router.DELETE("/pizzas/:id", handler.DeletePizzaById)
	router.PUT("/pizzas/:id", handler.UpdatePizzaById)
	router.Run()
}

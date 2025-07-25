package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas []models.Pizza

func main() {
	nomePizzaria := "Pizzaria Go"
	instagram, telefone := "@pizzaria_go", 1199711
	fmt.Println(nomePizzaria, instagram, telefone)
	loadPizzas()

	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.GET("/pizzas/:id", getPizzasById)
	router.POST("pizzas", postPizza)
	router.DELETE("/pizzas/:id", deletePizzaById)
	router.PUT("/pizzas/:id", updatePizzaById)
	router.Run()
}

func loadPizzas() {
	file, err := os.Open("dados/pizzas.json")

	if err != nil {
		fmt.Println("Erro no arquivo")
		return
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Println("Error deconding JSON:", err)
	}
}

func getPizzas(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}

func postPizza(ctx *gin.Context) {
	var newPizza models.Pizza

	if err := ctx.ShouldBindJSON(&newPizza); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error()})
	}

	newPizza.ID = len(pizzas) + 1
	pizzas = append(pizzas, newPizza)
	savePizza()
	ctx.Status(201)
}

func getPizzasById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error()})
		return
	}
	for _, p := range pizzas {
		if p.ID == id {
			c.JSON(200, p)
			return
		}
	}
	c.JSON(404, gin.H{"message": "Pizza not found"})
}

func savePizza() {
	file, err := os.Create("dados/pizzas.json")

	if err != nil {
		fmt.Println("Erro no arquivo")
		return
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pizzas); err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}

func deletePizzaById(c *gin.Context) {
	c.JSON(200, gin.H{"method": "delete"})
}

func updatePizzaById(c *gin.Context) {
	c.JSON(200, gin.H{"method": "update"})
}

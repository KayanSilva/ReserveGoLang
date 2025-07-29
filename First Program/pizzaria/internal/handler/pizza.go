package handler

import (
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"pizzaria/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPizzas(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"pizzas": data.Pizzas,
	})
}

func GetPizzasById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error()})
		return
	}
	for _, p := range data.Pizzas {
		if p.ID == id {
			c.JSON(200, p)
			return
		}
	}
	c.JSON(404, gin.H{"message": "Pizza not found"})
}

func PostPizza(ctx *gin.Context) {
	var newPizza models.Pizza

	if err := ctx.ShouldBindJSON(&newPizza); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error()})
		return
	}

	if err := service.ValidatePrice(&newPizza); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error()})
		return
	}

	newPizza.ID = len(data.Pizzas) + 1
	data.Pizzas = append(data.Pizzas, newPizza)
	data.SavePizza()
	ctx.Status(201)
}

func DeletePizzaById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error()})
		return
	}

	for i, p := range data.Pizzas {
		if p.ID == id {
			data.Pizzas = append(data.Pizzas[:i], data.Pizzas[i+1:]...)
			data.SavePizza()

			c.Status(204)
			return
		}
	}

	c.JSON(404, gin.H{"message": "pizza not found"})
}

func UpdatePizzaById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error()})
		return
	}

	var updatePizza models.Pizza
	if err := c.ShouldBindJSON(&updatePizza); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error()})
		return
	}

	if err := service.ValidatePrice(&updatePizza); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error()})
		return
	}

	for i, p := range data.Pizzas {
		if p.ID == id {
			data.Pizzas[i] = updatePizza
			data.Pizzas[i].ID = id
			data.SavePizza()

			c.Status(204)
			return
		}
	}

	c.JSON(404, gin.H{"message": "pizza not found"})
}

package handler

import (
	"net/http"
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"pizzaria/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostReview(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	var newReview models.Review
	if err := ctx.ShouldBindJSON(&newReview); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := service.ValidateReviewRating(newReview); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	for i, p := range data.Pizzas {
		if p.ID == id {
			data.Pizzas[i].Reviews = append(data.Pizzas[i].Reviews, newReview)
			data.SavePizza()
			ctx.Status(http.StatusCreated)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Pizza not found"})
}

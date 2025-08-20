package controllers

import (
	"net/http"

	"github.com/KayanSilva/ReserveGoLang/gin-api-rest/database"
	"github.com/KayanSilva/ReserveGoLang/gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, models.Students)
}

func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name parameter is required"})
		return
	}
	c.String(http.StatusOK, "Hello, %s!", name)
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&student)
	if student.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student"})
		return
	}

	c.JSON(http.StatusCreated, student)
}

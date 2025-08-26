package controllers

import (
	"fmt"
	"net/http"

	"github.com/KayanSilva/ReserveGoLang/gin-api-rest/database"
	"github.com/KayanSilva/ReserveGoLang/gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	var students []models.Student

	name := c.Query("name")
	cpf := c.Query("cpf")

	query := database.DB.Model(&models.Student{})

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if cpf != "" {
		query = query.Where("cpf LIKE ?", cpf)
	}

	page := 1
	limit := 10
	if p := c.Query("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}
	if l := c.Query("limit"); l != "" {
		fmt.Sscanf(l, "%d", &limit)
	}
	offset := (page - 1) * limit

	if err := query.Offset(offset).Limit(limit).Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve students"})
		return
	}

	totalCount := int64(0)
	database.DB.Model(&models.Student{}).Count(&totalCount)

	c.JSON(http.StatusOK, gin.H{
		"page":     page,
		"limit":    limit,
		"students": students,
		"total":    totalCount,
		"pages":    (totalCount + int64(limit) - 1) / int64(limit),
	})
}

func GetStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	if err := database.DB.First(&student, id).Error; err != nil {
		if err.Error() == "record not found" {
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve student"})
		return
	}

	c.JSON(http.StatusOK, student)
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
	if err := models.ValidateFields(&student); err != nil {
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

func DeleteStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := database.DB.Delete(&models.Student{}, id).Error; err != nil {
		if err.Error() == "record not found" {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student"})
		return
	}

	c.Status(http.StatusNoContent)
}

func UpdateStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	if err := database.DB.First(&student, id).Error; err != nil {
		if err.Error() == "record not found" {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve student"})
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidateFields(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	c.Status(http.StatusNoContent)
}

func GetIndexPage(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.HTML(http.StatusOK, "index.html", gin.H{"students": students})
}

func NotFoundPage(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}

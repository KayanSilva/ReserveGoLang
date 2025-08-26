package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/KayanSilva/ReserveGoLang/gin-api-rest/controllers"
	"github.com/KayanSilva/ReserveGoLang/gin-api-rest/database"
	"github.com/KayanSilva/ReserveGoLang/gin-api-rest/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRoutesTests() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	return r
}

func CreateStudentMock() {
	student := models.Student{
		Name: "Kayan Silva",
		CPF:  "12345678900",
		RG:   "199234567",
	}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestVerifyEndpointGreeting(t *testing.T) {
	r := SetupRoutesTests()
	r.GET("/:name", controllers.Greeting)
	req, _ := http.NewRequest("GET", "/Kayan", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
	mockResponse := `Hello, Kayan!`
	assert.Equal(t, mockResponse, response.Body.String(), "Devem ser iguais")
}

func TestAllStudentsHandler(t *testing.T) {
	database.Connect()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutesTests()
	r.GET("/students", controllers.GetStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetStudentWithCPFHandler(t *testing.T) {
	database.Connect()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutesTests()
	cpf := "12345678900"
	r.GET("/students", controllers.GetStudents)
	req, _ := http.NewRequest("GET", "/students?cpf="+cpf, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetStudantByIdHandler(t *testing.T) {
	database.Connect()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutesTests()
	r.GET("/students/:id", controllers.GetStudentById)
	req, _ := http.NewRequest("GET", "/students/"+strconv.Itoa(ID), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var studentMock models.Student
	json.Unmarshal(response.Body.Bytes(), &studentMock)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Kayan Silva", studentMock.Name, "Devem ser iguais")
}

func TestDeletedStudantHandler(t *testing.T) {
	database.Connect()
	CreateStudentMock()
	r := SetupRoutesTests()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	req, _ := http.NewRequest("DELETE", "/students/"+strconv.Itoa(ID), nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusNoContent, response.Code)
}

func TestUpdatedStudantHandler(t *testing.T) {
	database.Connect()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutesTests()
	r.PATCH("/students/:id", controllers.UpdateStudent)
	payload := models.Student{
		Name: "Kayan Silva",
		CPF:  "12345678900",
		RG:   "991234567",
	}
	bodyRequest, _ := json.Marshal(payload)
	req, _ := http.NewRequest("PATCH", "/students/"+strconv.Itoa(ID), bytes.NewBuffer(bodyRequest))
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusNoContent, response.Code)
}

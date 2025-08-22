package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KayanSilva/ReserveGoLang/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRoutesTests() *gin.Engine {
	r := gin.Default()
	return r
}

func TestVerifyEndpointGreeting(t *testing.T) {
	r := SetupRoutesTests()
	r.GET("/:name", controllers.Greeting)
	req, _ := http.NewRequest("GET", "/Kayan", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
	mockResponse := `"Hello, Kayan!"`
	assert.Equal(t, mockResponse, response.Body.String(), "Devem ser iguais")
}

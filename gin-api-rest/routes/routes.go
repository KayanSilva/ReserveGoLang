package routes

import (
	"github.com/KayanSilva/ReserveGoLang/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRoutes() {
	r := gin.Default()
	r.GET("students", controllers.GetStudents)
	r.GET(":name", controllers.Greeting)
	r.POST("students", controllers.CreateStudent)
	if err := r.Run(); err != nil {
		panic(err)
	}
}

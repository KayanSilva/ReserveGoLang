package routes

import (
	"github.com/KayanSilva/ReserveGoLang/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRoutes() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("../templates/*")
	r.Static("/assets", "../assets")
	r.GET(":name", controllers.Greeting)
	r.GET("students", controllers.GetStudents)
	r.GET("students/:id", controllers.GetStudentById)
	r.POST("students", controllers.CreateStudent)
	r.DELETE("students/:id", controllers.DeleteStudent)
	r.PATCH("students/:id", controllers.UpdateStudent)
	r.GET("/index", controllers.GetIndexPage)
	r.NoRoute(controllers.NotFoundPage)
	if err := r.Run(); err != nil {
		panic(err)
	}
}

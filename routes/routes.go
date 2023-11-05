package routes

import (
	"gin-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/api/students", controllers.ShowAllStudents)
	r.POST("/api/students", controllers.CreateStudent)
	r.Run()
}

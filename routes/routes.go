package routes

import (
	"gin-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/api/students", controllers.ShowAllStudents)
	r.GET("/api/students/:id", controllers.GetStudentById)
	r.POST("/api/students", controllers.CreateStudent)
	r.DELETE("/api/students/:id", controllers.DeleteStudent)
	r.PATCH("/api/students/:id", controllers.EditStudent)
	r.GET("/api/students/search/cpf/:cpf", controllers.SearchByCpf)
	r.Run()
}

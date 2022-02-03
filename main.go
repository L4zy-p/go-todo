package main

import (
	"go-todo/controllers"
	"go-todo/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", controllers.CheckStatus)
	r.GET("/todo", controllers.GetAllLists)
	r.POST("/todo", controllers.CreateTodoList)

	r.Run(":3001")
}

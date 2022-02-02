package main

import (
	"go-todo/controllers"
	"go-todo/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", controllers.GetAllLists)

	r.Run(":3001")
}

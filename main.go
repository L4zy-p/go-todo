package main

import (
	"go-todo/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", controllers.GetTodo)

	r.Run(":3001")
}

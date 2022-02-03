package main

import (
	"go-todo/controllers"
	"go-todo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", controllers.CheckStatus)
	r.GET("/todo", controllers.GetAllLists)
	r.GET("/todo/:id", controllers.GetTodoListById)
	r.POST("/todo", controllers.CreateTodoList)
	r.POST("/upload", controllers.Upload)
	r.DELETE("/todo/:id", controllers.DeleteTodoListById)

	r.GET("/user", controllers.GetTodoListByUsername)

	// ให้ path file แสดง file ทั้งหมดที่อยู่ใน public
	r.StaticFS("/file", http.Dir("public"))

	r.Run(":3001")
}

package controllers

import (
	"go-todo/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "server started",
	})
}

func GetAllLists(c *gin.Context) {
	var todoLists []models.Todo
	models.DB.Find(&todoLists)

	c.JSON(http.StatusOK, gin.H{
		"data": todoLists,
	})
}

func CreateTodoList(c *gin.Context) {
	var input models.TodoCreate

	// check ว่าข้อมูลที่อยู่ใน json ที่โยนขึ้นมาใช้ได้มั้ย
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	todoList := models.Todo{Username: input.Username, Title: input.Title, Message: input.Message}
	models.DB.Create(&todoList)

	c.JSON(http.StatusOK, gin.H{
		"staus": "ok",
		"data":  todoList,
	})
}

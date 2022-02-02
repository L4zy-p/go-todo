package controllers

import (
	"go-todo/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func GetAllLists(c *gin.Context) {
	var todoLists []models.Todo
	models.DB.Find(&todoLists)

	c.JSON(http.StatusOK, gin.H{
		"data": todoLists,
	})
}

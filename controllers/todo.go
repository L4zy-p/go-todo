package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

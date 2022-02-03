package controllers

import (
	"fmt"
	"go-todo/models"
	"io"
	"log"
	"os"

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

func GetTodoListById(c *gin.Context) {
	var todoList models.Todo

	// ให้ err เท่ากับการค้นหา todo จาก id ที่ param id มา แล้วเช็ค err ต่อว่า มี error หรือป่าว
	// First เป็นการหาตัวแรก ที่มี id นี้แค่ตัวเดียวเท่านั้น
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todoList).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todoList,
	})
}

func GetTodoListByUsername(c *gin.Context) {
	var todoList []models.Todo

	// ให้ err เท่ากับการค้นหา todo จาก username ที่ query username มา แล้วเช็ค err ต่อว่า มี error หรือป่าว
	// Find จะเป็นการหาทุกตัวที่มี Username นี้
	if err := models.DB.Where("username = ?", c.Query("username")).Find(&todoList).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todoList,
	})
}

func DeleteTodoListById(c *gin.Context) {
	var todoList models.Todo

	// ให้ err เท่ากับการค้นหา todo จาก id ที่ param id มา แล้วเช็ค err ต่อว่า มี error หรือป่าว
	// First เป็นการหาตัวแรก ที่มี id นี้เพื่อที่จะทำการลบต่อ
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todoList).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Record not found",
		})
		return
	}

	models.DB.Delete(&todoList)
	c.JSON(http.StatusOK, gin.H{
		"result": "Delete successfuly",
	})
}

func Upload(c *gin.Context) {
	// ดู field file จาก request formfile ที่ส่งมา ซึ่ง return ค่ามาให้ 3 ตัวคือ file, header, err
	file, header, err := c.Request.FormFile("file")

	// เช็คมาว่าที่ส่ง formfile มามี error หรือป่าว
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err %s", err.Error()))
		return
	}

	// get ชื่อไฟล์จาก header
	filename := header.Filename

	// ทำการสร้างไฟล์ให้อยู่ในโฟลเดอร์ public
	out, err := os.Create("public/" + filename)
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	// ทำการ copy file ที่ส่งมาจาก request แล้วใส่ไปใน path ที่เกิดจาก os.Create
	_, err = io.Copy(out, file)

	if err != nil {
		log.Fatal(err)
	}

	filepath := "http://localhost:3001/file/" + filename
	c.JSON(http.StatusOK, gin.H{
		"filepath": filepath,
	})
}

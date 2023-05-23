package controllers

import (
	"github.com/bookpanda/champ-eng-go-crud/database"
	"github.com/bookpanda/champ-eng-go-crud/models"
	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	database.DB.Find(&tasks)

	c.JSON(200, gin.H{
		"tasks": tasks,
	})
}

func GetTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	database.DB.First(&task, id)

	c.JSON(200, gin.H{
		"task": task,
	})
}

func CreateTask(c *gin.Context) {
	var body struct {
		Description string
		DueDate     string
		Order       int
	}
	c.Bind(&body)
	task := models.Task{Description: body.Description, DueDate: body.DueDate, Order: body.Order}
	database.DB.Create(&task)

	c.JSON(200, gin.H{
		"task": task,
	})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Description string
		DueDate     string
		Order       int
	}
	c.Bind(&body)
	var task models.Task
	database.DB.First(&task, id)
	database.DB.Model(&task).Updates(models.Task{Description: body.Description, DueDate: body.DueDate, Order: body.Order})

	c.JSON(200, gin.H{
		"task": task,
	})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	database.DB.First(&task, id)
	database.DB.Delete(&models.Task{}, id)

	c.JSON(200, gin.H{
		"task": task,
	})
}

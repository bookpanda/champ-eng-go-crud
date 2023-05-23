package controllers

import (
	services "github.com/bookpanda/champ-eng-go-crud/services/taskServices"
	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	services.GetTasks(c)
}

func GetTask(c *gin.Context) {
	services.GetTask(c)
}

func CreateTask(c *gin.Context) {
	services.CreateTask(c)
}

func UpdateTask(c *gin.Context) {
	services.UpdateTask(c)
}

func DeleteTask(c *gin.Context) {
	services.DeleteTask(c)
}

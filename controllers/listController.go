package controllers

import (
	services "github.com/bookpanda/champ-eng-go-crud/services/listServices"
	"github.com/gin-gonic/gin"
)

func GetLists(c *gin.Context) {
	services.GetLists(c)
}

func GetList(c *gin.Context) {
	services.GetList(c)
}

func CreateList(c *gin.Context) {
	services.CreateList(c)
}

func UpdateList(c *gin.Context) {
	services.UpdateList(c)
}

func DeleteList(c *gin.Context) {
	services.DeleteList(c)
}

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMedicine(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"data": ""})
}

func GetMedicine(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": ""})
}

func UpdateMedicine(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"data": id})
}

func DeleteMedicine(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"data": id})
}

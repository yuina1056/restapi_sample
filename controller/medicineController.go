package controller

import (
	"net/http"
	"test_restapi/db"
	"test_restapi/models"

	"github.com/gin-gonic/gin"
)

func CreateMedicine(c *gin.Context) {
	var medicineInput models.Medicine
	err := c.BindJSON(&medicineInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Validation Error"})
		return
	}

	insertMedicine := models.Medicine{
		Name: medicineInput.Name,
		Type: medicineInput.Type,
	}
	err = db.DB.Create(&insertMedicine).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "DB Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": insertMedicine})
}

func GetMedicine(c *gin.Context) {
	findMedicines := []models.Medicine{}
	err := db.DB.Find(&findMedicines).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "DB Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": findMedicines})
}

func UpdateMedicine(c *gin.Context) {
	id := c.Param("id")
	var medicineInput models.Medicine
	err := c.BindJSON(&medicineInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Validation Error"})
		return
	}

	findMedicine := models.Medicine{}
	err = db.DB.Where("id = ?", id).First(&findMedicine).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "DB Error"})
		return
	}

	updateMedicine := models.Medicine{
		Name: medicineInput.Name,
		Type: medicineInput.Type,
	}
	err = db.DB.Model(&findMedicine).Update(updateMedicine).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "DB Error"})
		return
	}

	err = db.DB.Where("id = ?", id).First(&findMedicine).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "DB Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": findMedicine})
}

func DeleteMedicine(c *gin.Context) {
	id := c.Param("id")

	findMedicine := models.Medicine{}
	err := db.DB.Where("id = ?", id).First(&findMedicine).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "DB Error"})
		return
	}

	err = db.DB.Delete(&findMedicine).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "DB Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

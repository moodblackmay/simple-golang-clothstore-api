package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"simple-golang-clothstore-api/models"
)

//GET
func FindClothes(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var clothes []models.Cloth
	db.Find(&clothes)

	c.JSON(http.StatusOK, gin.H{"data": clothes})
}

//GET
func FindCloth(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//Get model if exist
	var cloth models.Cloth
	if err := db.Where("id = ?", c.Param("id")).First(&cloth).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cloth})
}

//POST
func CreateCloth(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//Validate input
	var input models.CreateCloth
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	cloth := models.Cloth{Name: input.Name, Description: input.Description, Price: input.Price}
	db.Create(&cloth)

	c.JSON(http.StatusOK, gin.H{"data": cloth})
}

//UPDATE
func UpdateCloth(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)


	//Get model if exist
	var cloth models.Cloth
	if err := db.Where("id = ?", c.Param("id")).First(&cloth).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"not found"})
		return
	}

	//Validate input
	var input models.UpdateCloth
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err})
	}

	db.Model(&cloth).Update(input)

	c.JSON(http.StatusOK, gin.H{"data":cloth})
}

//DELETE
func DeleteCloth(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)


	//Get model if exist
	var cloth models.Cloth
	if err := db.Where("id = ?", c.Param("id")).First(&cloth).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"not found"})
		return
	}

	db.Delete(&cloth)
	c.JSON(http.StatusOK, gin.H{"data":"ok"})
}
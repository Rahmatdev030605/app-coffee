package handlers

import (
	"net/http"

	"github.com/Rahmatdev030605/app-coffee/db"
	"github.com/Rahmatdev030605/app-coffee/models"
	"github.com/gin-gonic/gin"
)

func GetCake(c *gin.Context) {
	var cakes []models.MenuCake
	db.DB.Find(&cakes)
	c.JSON(http.StatusOK, cakes)
}

func CreateCake(c *gin.Context)  {
	var cakeInput models.MenuCake
if err := c.ShouldBindJSON(&cakeInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPrice := cakeInput.PriceCake * float64(cakeInput.TotalCake)
	cakeInput.TotalPrice = totalPrice

	cake := models.MenuCake{
		NameCake: cakeInput.NameCake,
		PriceCake: cakeInput.PriceCake,
		TotalCake: cakeInput.TotalCake,
		TotalPrice: cakeInput.TotalPrice,
	}
	
	if err := db.DB.Create(&cake).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Create Cake Menu"})
		return
	}

	db.DB.First(&cake, cake.ID)
	c.JSON(http.StatusOK, cake)
}

func UpdateCake(c *gin.Context)  {
	id := c.Param("id")
	
	var cake models.MenuCake
	if err := db.DB.First(&cake, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"The Menu is not updating"})
		return
	}
	if err := c.ShouldBindJSON(&cake); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	totalPrice := cake.PriceCake * float64(cake.TotalCake)
	cake.TotalPrice = totalPrice

	db.DB.Save(&cake)
	c.JSON(http.StatusOK, cake)
}

func DeleteCake(c *gin.Context)  {
	id := c.Param("id")

	var cake models.MenuCake
	if err := db.DB.First(&cake, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "The Cake is not Delete"})
		return
	}
	db.DB.Delete(&cake)
	c.JSON(http.StatusOK, gin.H{"message": "Cake Deleted"})
}
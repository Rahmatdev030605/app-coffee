package handlers

import (
	"net/http"

	"github.com/Rahmatdev030605/app-coffee/db"
	"github.com/Rahmatdev030605/app-coffee/models"
	"github.com/gin-gonic/gin"
)

func GetSnack(c *gin.Context) {
	var snack []models.MenuSnack
	db.DB.Find(&snack)
	c.JSON(http.StatusOK, snack)
}

func CreateSnack(c *gin.Context)  {
	var snackInput models.MenuSnack
	if err := c.ShouldBindJSON(&snackInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPrice := snackInput.PriceSnack * float64(snackInput.TotalSnack)
	snackInput.TotalPrice = totalPrice

	snack := models.MenuSnack{
		NameSnack: snackInput.NameSnack,
		PriceSnack: snackInput.PriceSnack,
		TotalSnack: snackInput.TotalSnack,
		TotalPrice: snackInput.TotalPrice,
	}
	
	if err := db.DB.Create(&snack).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Create Snack Menu"})
		return
	}

	db.DB.First(&snack, snack.ID)
	c.JSON(http.StatusOK, snack)
}

func UpdateSnack(c *gin.Context)  {
	id := c.Param("id")
	
	var snack models.MenuSnack
	if err := db.DB.First(&snack, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"The Menu is not updating"})
		return
	}
	if err := c.ShouldBindJSON(&snack); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPrice := snack.PriceSnack * float64(snack.TotalSnack)
	snack.TotalPrice = totalPrice

	db.DB.Save(&snack)
	c.JSON(http.StatusOK, snack)
}

func DeleteSnack(c *gin.Context)  {
	id := c.Param("id")

	var snack models.MenuSnack
	if err := db.DB.First(&snack, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "The snack is not Delete"})
		return
	}
	db.DB.Delete(&snack)
	c.JSON(http.StatusOK, gin.H{"message": "Cake Deleted"})
}
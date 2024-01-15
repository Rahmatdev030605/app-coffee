package handlers

import (
	"net/http"

	"github.com/Rahmatdev030605/app-coffee/db"
	"github.com/Rahmatdev030605/app-coffee/models"
	"github.com/gin-gonic/gin"
)

func GetCoffeeTypes(c *gin.Context) {
	var coffeeTypes []models.CoffeeType
	db.DB.Find(&coffeeTypes)
	c.JSON(http.StatusOK, coffeeTypes)
}

func CreateCoffeeType(c *gin.Context) {
	var coffeeType models.CoffeeType
	if err := c.ShouldBindJSON(&coffeeType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&coffeeType)
	c.JSON(http.StatusOK, coffeeType)
}

func UpdateCoffeeType(c *gin.Context) {
	id := c.Param("id")
	var coffeeType models.CoffeeType
	if err := db.DB.First(&coffeeType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CoffeeType not found"})
		return
	}

	if err := c.ShouldBindJSON(&coffeeType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&coffeeType)
	c.JSON(http.StatusOK, coffeeType)
}

func DeleteCoffeeType(c *gin.Context) {
	id := c.Param("id")

	var coffeeType models.CoffeeType
	if err := db.DB.First(&coffeeType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CoffeeType not found"})
		return
	}

	db.DB.Delete(&coffeeType)
	c.JSON(http.StatusOK, gin.H{"message": "CoffeeType deleted"})
}

func GetAllCoffeeTypes(c *gin.Context) {
	var coffeeTypes []models.CoffeeType
	db.DB.Find(&coffeeTypes)
	c.JSON(http.StatusOK, coffeeTypes)
}

func GetCoffeeTypeByID(c *gin.Context) {
	id := c.Param("id")

	var coffeeType models.CoffeeType
	if err := db.DB.First(&coffeeType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CoffeeType not found"})
		return
	}

	c.JSON(http.StatusOK, coffeeType)
}

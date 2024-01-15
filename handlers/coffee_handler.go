package handlers

import (
	"net/http"

	"github.com/Rahmatdev030605/app-coffee/db"
	"github.com/Rahmatdev030605/app-coffee/models"
	"github.com/gin-gonic/gin"
)

func GetCoffees(c *gin.Context) {
	var coffees []models.Coffee

	db.DB.Preload("CoffeeType").Find(&coffees)

	c.JSON(http.StatusOK, coffees)
}


func CreateCoffee(c *gin.Context) {
	var coffeeInput models.Coffee
	if err := c.ShouldBindJSON(&coffeeInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	coffee := models.Coffee{
		Name: 			coffeeInput.Name,
		Price: 			coffeeInput.Price,
		CoffeeTypeID: 	coffeeInput.CoffeeTypeID,
	}

	if err := db.DB.Create(&coffee).Error;err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create coffee"})
		return
	}

	db.DB.Preload("CoffeeType").First(&coffee, coffee.ID)
	c.JSON(http.StatusOK, coffee)
}

func UpdateCoffee(c *gin.Context) {
	id := c.Param("id")

	var coffee models.Coffee
	if err := db.DB.First(&coffee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Coffee not found"})
		return
	}

	if err := c.ShouldBindJSON(&coffee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&coffee)
	c.JSON(http.StatusOK, coffee)
}

func DeleteCoffee(c *gin.Context) {
	id := c.Param("id")

	var coffee models.Coffee
	if err := db.DB.First(&coffee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Coffee not found"})
		return
	}

	db.DB.Delete(&coffee)
	c.JSON(http.StatusOK, gin.H{"message": "Coffee deleted"})
}

func GetAllCoffees(c *gin.Context) {
	var coffees []models.Coffee
	db.DB.Find(&coffees)
	c.JSON(http.StatusOK, coffees)
}

func GetCoffeeByID(c *gin.Context) {
	id := c.Param("id")

	var coffee models.Coffee
	if err := db.DB.Preload("CoffeeType").First(&coffee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Coffee not found"})
		return
	}

	c.JSON(http.StatusOK, coffee)
}

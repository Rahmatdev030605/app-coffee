package handlers

import (
	"net/http"

	"github.com/Rahmatdev030605/app-coffee/db"
	"github.com/Rahmatdev030605/app-coffee/models"
	"github.com/gin-gonic/gin"
)

func GetTables(c *gin.Context) {
	var tables []models.Table
	db.DB.Find(&tables)
	c.JSON(http.StatusOK, tables)
}

func CreateTable(c *gin.Context)  {
	var table models.Table
	if err := c.ShouldBindJSON(&table); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Create(&table)
	c.JSON(http.StatusOK, table)
}

func UpdateTable(c *gin.Context)  {
	id := c.Param("id")

	var tables models.Table
	if err := db.DB.First(&tables, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	db.DB.Save(&tables)
	c.JSON(http.StatusOK, tables)
}

func DeleteTable(c *gin.Context)  {
	id := c.Param("id")

	var tables models.Table
	if err := db.DB.First(&tables, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"Table not found"})
	}

	db.DB.Delete(&tables)
	c.JSON(http.StatusOK, gin.H{"message": "Table deleted"})
}

func GetTableById(c *gin.Context)  {
	id := c.Param("id")

	var tables models.Table
	if err := db.DB.First(&tables, id).Error;err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Table not found"})
	}
	c.JSON(http.StatusOK, tables)
}
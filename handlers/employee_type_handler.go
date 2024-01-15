package handlers

import (
	"net/http"

	"github.com/Rahmatdev030605/app-coffee/db"
	"github.com/Rahmatdev030605/app-coffee/models"
	"github.com/gin-gonic/gin"
)

func GetEmployeeTypes(c *gin.Context) {
	var employeeTypes []models.EmployeeType
	db.DB.Find(&employeeTypes)
	c.JSON(http.StatusOK, employeeTypes)
}

func CreateEmployeeType(c *gin.Context) {
	var employeeType models.EmployeeType
	if err := c.ShouldBindJSON(&employeeType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&employeeType)
	c.JSON(http.StatusOK, employeeType)
}

func UpdateEmployeeType(c *gin.Context) {
	id := c.Param("id")

	var employeeType models.EmployeeType
	if err := db.DB.First(&employeeType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "EmployeeType not found"})
		return
	}

	if err := c.ShouldBindJSON(&employeeType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&employeeType)
	c.JSON(http.StatusOK, employeeType)
}

func DeleteEmployeeType(c *gin.Context) {
	id := c.Param("id")

	var employeeType models.EmployeeType
	if err := db.DB.First(&employeeType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "EmployeeType not found"})
		return
	}

	db.DB.Delete(&employeeType)
	c.JSON(http.StatusOK, gin.H{"message": "EmployeeType deleted"})
}

func GetAllEmployeeTypes(c *gin.Context) {
	var employeeTypes []models.EmployeeType
	db.DB.Find(&employeeTypes)
	c.JSON(http.StatusOK, employeeTypes)
}

func GetEmployeeTypeByID(c *gin.Context) {
	id := c.Param("id")

	var employeeType models.EmployeeType
	if err := db.DB.First(&employeeType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "EmployeeType not found"})
		return
	}

	c.JSON(http.StatusOK, employeeType)
}
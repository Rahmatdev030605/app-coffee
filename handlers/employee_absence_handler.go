package handlers

import (
	"net/http"

	"github.com/Rahmatdev030605/app-coffee/db"
	"github.com/Rahmatdev030605/app-coffee/models"
	"github.com/gin-gonic/gin"
)

func GetAbsenceEmployee(c *gin.Context) {
	var absenceEmployee []models.EmployeeAbsence
	db.DB.Preload("Employee.EmployeeType").Find(&absenceEmployee)
	c.JSON(http.StatusOK, absenceEmployee)
}

func AddAbsenceEmployee(c *gin.Context)  {
	var absenceEmployee models.EmployeeAbsence
	if err := c.ShouldBindJSON(&absenceEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	absence := models.EmployeeAbsence{
		EmployeeID: 	absenceEmployee.EmployeeID,
		Comed: 			absenceEmployee.Comed,
	}

	if err := db.DB.Create(&absence).Error;err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Absence"})
	}

	db.DB.Preload("Employee").First(&absence, absence.ID)
	c.JSON(http.StatusOK, absence)
}

func GetAbsenceEmployeeByID(c *gin.Context)  {
	id := c.Param("id")

	var absence models.EmployeeAbsence
	if err := db.DB.Preload("Employee").First(&absence, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"ID employee not found"})
	}
	c.JSON(http.StatusOK, absence)
}
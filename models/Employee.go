package models

import (

	"github.com/jinzhu/gorm"
)

type Employee struct {
	gorm.Model
	Name           string       `json:"name"`
	Age            int          `json:"age"`
	Address        string       `json:"address"`
	Gender         string	    `json:"gender"`
	EmployeeTypeID uint         `json:"employee_type_id"`
	EmployeeType   EmployeeType `gorm:"foreignkey:EmployeeTypeID" json:"employee_type"`
}


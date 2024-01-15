package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type EmployeeAbsence struct {
	gorm.Model
	
	EmployeeID 		uint			`json:"employee_id"`
	Comed			time.Time		`json:"comed"`

	Employee 		Employee 		`gorm:"foreignkey:EmployeeID" json:"employee"`


}

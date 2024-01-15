package models

import "github.com/jinzhu/gorm"

type EmployeeType struct {
	gorm.Model
	Name string `json:"name`
}

func SeedEmployeeTypes(db *gorm.DB)  {
	types := []EmployeeType{
		{Name: "Barista"},
		{Name: "Cashier"},
		{Name: "Chef"},
		{Name: "Dishwasher"},
		{Name: "Resepsionis"},
		{Name: "Waiter"},
		{Name: "Graphic designer"},
	}
	for _, t:= range types{
		db.Create(&t)
	}
}
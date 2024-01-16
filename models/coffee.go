package models

import "github.com/jinzhu/gorm"	

type Coffee struct {
	gorm.Model
	Name 					string 			`json:="name"`
	Price 					float64 		`json:"price"`
	TotalCoffee				float64			`json:"total_coffee"`
	TotalPrice				float64			`json:"total_price"`
	CoffeeTypeID 			uint 			`json:"coffee_type_id"`
	CoffeeType 				CoffeeType 		`gorm:"foreignkey:CoffeeTypeID" json:"coffee_type"`
}

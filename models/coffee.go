package models

import "github.com/jinzhu/gorm"

type Coffee struct {
	gorm.Model
	Name 					string 			`json:="name"`
	Price 					float64 		`json:"price"`
	CoffeeTypeID 			uint 			`json:"coffee_type_id"`
	CoffeeType 				CoffeeType 		`gorm:"foreignkey:CoffeeTypeID" json:"coffee_type"`
}

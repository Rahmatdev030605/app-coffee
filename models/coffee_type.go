package models

import "github.com/jinzhu/gorm"

type CoffeeType struct {
	gorm.Model
	Name string `json:"name"`
}

func SeedCoffeeType(db *gorm.DB)  {
	types := []CoffeeType{
		{Name: "Espresso"},
		{Name: "Americano"},
		{Name: "Latte"},
		{Name: "Espresso"},
		{Name: "Long Black"},
		{Name: "Cappuchino"},
		{Name: "Caffé Latte"},
		{Name: "Mocha Latte"},
		{Name: "Affogato"},
		{Name: "Cold Brew"},
	}
	for _, t := range types{
		db.Create(&t)
}
}
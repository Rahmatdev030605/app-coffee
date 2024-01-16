package models

import "github.com/jinzhu/gorm"

type MenuSnack struct {
	gorm.Model
	NameSnack	SnackMethod	`json:"name_snack"`
	PriceSnack	float64	`json:"price_snack"`
	TotalSnack	float64	`json:"total_snack"`
	TotalPrice	float64	`json:"total_price"`
}

type SnackMethod string

const (
	OnionRings		SnackMethod	= "OnionRings"
	FrenchFries 	SnackMethod	= "FrenchFries"
	Resoles 		SnackMethod	= "Resoles"
	DeBananaFries	SnackMethod = "DeBananaFries"
)
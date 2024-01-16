package models

import "github.com/jinzhu/gorm"

type MenuCake struct {
	gorm.Model
	NameCake	CakeMethod	`json:"name_cake"`
	PriceCake 	float64		`json:"price_cake"`
	TotalCake 	float64		`json:"total_cake"`
	TotalPrice	float64		`json:"total_price`
}

type CakeMethod string

const (
	CaramelFudge 		CakeMethod	= "CaramelFudge" 
	ChocoCake			CakeMethod	= "ChocoCake"
	PeanutButterCake	CakeMethod	= "PeanutButterCake"
	RainbowCake 		CakeMethod	= "RainbowCake"
	RedVelvetCake 		CakeMethod	= "RedVelvetCake"
)
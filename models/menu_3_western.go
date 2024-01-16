package models

import "github.com/jinzhu/gorm"

type MenuWestern struct {
	gorm.Model
	NameWestern  WesternMethod `json:"name_western"`
	PriceWestern float64       `json:"price_western"`
	TotalWestern float64       `json:"total_western"`
	TotalPrice   float64       `json:"total_Price`
}

type WesternMethod string

const (
	ChickenCordonBlue    WesternMethod = "ChickenCordonBlue"
	BeefSteakBBQ         WesternMethod = "BeefSteakBBQ"
	ChickenSteakMushroom WesternMethod = "ChickenSteakMushroom"
	Spageti              WesternMethod = "Spageti"
)
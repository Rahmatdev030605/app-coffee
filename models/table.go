package models

import "github.com/jinzhu/gorm"

type Table struct {
	gorm.Model
	Number	uint `json:"number"`
}
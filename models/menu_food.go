package models

import "github.com/jinzhu/gorm"

type MenuFood struct {
	gorm.Model
	MenuTypeId 	uint 		`json:"menu_type_id"`
	Cake		MenuCake	`gorm:"foreignkey:MenuTypeID" json:"cake"`
	Snack 		MenuSnack	`gorm:"foreignkey:MenuTypeID" json:"snack"`
	Western 	MenuWestern	`gorm:"foreignkey:MenuTypeID" json:"western"`
}
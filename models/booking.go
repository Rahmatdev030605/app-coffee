package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Booking struct {
	gorm.Model
	UserID      	uint      			`json:"user_id"`
	NameBooking 	string    			`json:"name_booking"`
	CoffeeID    	uint      			`json:"coffee_id"`
	TableID     	uint      			`json:"table_id"`
	Schedule    	time.Time 			`json:"schedule"`
	histories		[]HistoryBooking	`gorm:"foreignkey:BookingID" json:"histories"`
	User        	User      			`gorm:"foreignkey:UserID" json:"user"`
	Coffee      	Coffee    			`gorm:"foreignkey:CoffeeID" json:"coffee"`
	Table      	 	Table     			`gorm:"foreignkey:TableID" json:"table"`
}


package models

import "github.com/jinzhu/gorm"

type HistoryBooking struct {
	gorm.Model

	BookingID uint   `json:"booking_id"`
	Status    string `json:"status"`
	Note      string `json:"note"`
	UpdatedBy uint   `json:"updated_by"`

	Booking   Booking `gorm:"foreignkey:BookingID" json:"booking"`
}

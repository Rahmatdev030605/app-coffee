package models

import "github.com/jinzhu/gorm"

type Payment struct {
	gorm.Model
	UserID         uint          `json:"user_id"`
	CoffeeID       uint          `json:"coffee_id"`
	MenuCakeID     uint          `json:"menu_cake_id"`
	MenuSnackID    uint          `json:"menu_snack_id"`
	MenuWesternID  uint          `json:"menu_western_id"`
	TotalPrice     float64       `json:"total_price"`
	Status         StatusMethod  `json:"status"`
	PaymentMethod  PaymentMethod `json:"payment_method"`
	User           User          `gorm:"foreignkey:UserID" json:"user"`
	Coffee         Coffee        `gorm:"foreignkey:CoffeeID" json:"coffee"`
	Cake           MenuCake      `gorm:"foreignkey:MenuCakeID" json:"cake"`
	Snack          MenuSnack     `gorm:"foreignkey:MenuSnackID" json:"snack"`
	Western        MenuWestern   `gorm:"foreignkey:MenuWesternID" json:"western"`
}


type PaymentMethod string

const (
	CreditCard PaymentMethod = "CreditCard"
	DebitCard  PaymentMethod = "DebitCard"
	PayPal     PaymentMethod = "Paypal"
)

type StatusMethod string

const (
	TransactionSuccessful	StatusMethod = "TransactionSuccessful"
	TransactionFailed		StatusMethod = "T	ransactionFailed"
)

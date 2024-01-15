package handlers

import (
	"net/http"
	"strconv"

	"github.com/Rahmatdev030605/app-coffee/db"
	"github.com/Rahmatdev030605/app-coffee/models"
	"github.com/gin-gonic/gin"
)

func GetBookings(c *gin.Context) {
	var bookings []models.Booking
	db.DB.Preload("User").Preload("Coffee.CoffeeType").Preload("Table").Find(&bookings)
	c.JSON(http.StatusOK, bookings)
}

func CreateBooking(c *gin.Context) {
	var bookingInput models.Booking
	if err := c.ShouldBindJSON(&bookingInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	booking := models.Booking{
		NameBooking: bookingInput.NameBooking,
		UserID:      bookingInput.UserID,
		CoffeeID:    bookingInput.CoffeeID,
		TableID:     bookingInput.TableID,
		Schedule:    bookingInput.Schedule,
	}

	if err := db.DB.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	history := models.HistoryBooking{
		BookingID: booking.ID,
		Status:    "Booked", 
		Note:      "Booking created",
		UpdatedBy: booking.UserID,
	}
	db.DB.Create(&history)

	
	db.DB.Preload("User").Preload("Coffee.CoffeeType").Preload("Table").First(&booking, booking.ID)
	c.JSON(http.StatusOK, booking)
}


func UpdateBooking(c *gin.Context)  {
	id := c.Param("id")

	var booking models.Booking
	if err := db.DB.First(&booking, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"The schedule you booked is not available."})
		return
	}
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&booking)
	c.JSON(http.StatusOK, booking)
}

func DeleteBooking(c *gin.Context)  {
	id := c.Param("id")
	
	var booking models.Booking
	if err := db.DB.First(&booking, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"The schedule you booked is not available."})
		return
	}
	db.DB.Delete(&booking)
	c.JSON(http.StatusOK, gin.H{"message":"Booked deleted"})
}

func GetBookingByID(c *gin.Context)  {
	id := c.Param("id")

	var booking models.Booking
	if err := db.DB.Preload("User").Preload("Coffee").Preload("Table").First(&booking, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booked not found"})
	}
	c.JSON(http.StatusOK, booking)
}


func GetBookingHistoriesByID(c *gin.Context) {
	bookingID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	var histories []models.HistoryBooking
	db.DB.Where("booking_id = ?", bookingID).Find(&histories)

	db.DB.Preload("Booking.User").Preload("Booking.Coffee.CoffeeType").Preload("Booking.Table").Find(&histories)

	c.JSON(http.StatusOK, histories)
}


// func GetBookingHistory(c *gin.Context)  {
// 	var booking []models.HistoryBooking
// 	db.DB.Preload("Booking.User.CoffeeType.Coffee.Table").Find(&booking)
// 	c.JSON(http.StatusOK, booking)
// }	

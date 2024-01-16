package handlers

import (
	"net/http"

	"github.com/Rahmatdev030605/app-coffee/db"
	"github.com/Rahmatdev030605/app-coffee/models"
	"github.com/gin-gonic/gin"
)

func GetPayment(c *gin.Context) {
	var payments []models.Payment
	db.DB.Preload("User").Preload("Coffee.CoffeeType")
	c.JSON(http.StatusOK, payments)
}

func CreatePayment(c *gin.Context) {
	var paymentInput models.Payment
	if err := c.ShouldBindJSON(&paymentInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPrice := calculateTotalPrice(paymentInput.Coffee, paymentInput.Cake, paymentInput.Snack, paymentInput.Western)

		paymentInput.Status = "Success Transaction"

		payment := models.Payment{
			UserID:        paymentInput.UserID,
			CoffeeID:      paymentInput.CoffeeID,
			MenuCakeID:    paymentInput.MenuCakeID,
			MenuSnackID:   paymentInput.MenuSnackID,
			MenuWesternID: paymentInput.MenuWesternID,
			TotalPrice:    totalPrice,
			Status:        paymentInput.Status,
			PaymentMethod: paymentInput.PaymentMethod,
		}
	if err := db.DB.Create(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment"})
		return
	}

	if err := db.DB.Preload("User").Preload("Coffee.CoffeeType").
		Preload("Cake").Preload("Snack").Preload("Western").
		First(&payment, payment.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to preload payment data"})
		return
	}

	c.JSON(http.StatusOK, payment)
}


func UpdatePayment(c *gin.Context)  {
	id := c.Param("id")

	var payment models.Payment
	if err := db.DB.First(&payment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to Update Payment"})
		return
	}
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	totalPrice := payment.Coffee.TotalPrice + payment.Cake.TotalPrice +
				  payment.Snack.TotalPrice + payment.Western.TotalPrice

				  payment.TotalPrice = totalPrice

	db.DB.Save(&payment)
	c.JSON(http.StatusOK, payment)
}

func DeletePayment(c *gin.Context)  {
	id := c.Param("id")

	var payment models.Payment
	if err := db.DB.First(&payment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "The Payment you is not available"})
		return
	}
	db.DB.Delete(&payment)
	c.JSON(http.StatusOK, gin.H{"message": "Payment Deleted"})
}

func GetPaymentById(c *gin.Context)  {
	id := c.Param("id")

	var payment models.Payment
	if err := db.DB.Preload("User").Preload("Coffee.CoffeeType").First(&payment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"payment not found"})
	}
	c.JSON(http.StatusOK, payment)
}

// func getPaymentHistoryByID()  {

// }

func calculateTotalPrice(coffee models.Coffee, cake models.MenuCake, snack models.MenuSnack, western models.MenuWestern) float64 {
    totalPrice := coffee.TotalPrice

    if cake != (models.MenuCake{}) {
        totalPrice += cake.TotalPrice
    }
    if snack != (models.MenuSnack{}) {
        totalPrice += snack.TotalPrice
    }
    if western != (models.MenuWestern{}) {
        totalPrice += western.TotalPrice
    }

    return totalPrice
}

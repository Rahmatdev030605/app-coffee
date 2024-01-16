package handlers

// import (
// 	"net/http"

// 	"github.com/Rahmatdev030605/app-coffee/models"
// 	"github.com/gin-gonic/gin"
// )

// func CreateMenuInput(c *gin.Context) {
// 	var foodMenuInput models.MenuFood
// 	if err := c.ShouldBindJSON(&foodMenuInput); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	cake := models.MenuCake{
// 		NameCake: "CaramelFudge",
// 		PriceCake: 40000,
// 		TotalPrice: ,
// 	}
// }
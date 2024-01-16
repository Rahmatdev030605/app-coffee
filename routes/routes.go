// routes/routes.go
package routes

import (
	"github.com/Rahmatdev030605/app-coffee/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	// CoffeeType routes
	api.GET("/coffee-types", handlers.GetCoffeeTypes)
	api.POST("/coffee-types-add", handlers.CreateCoffeeType)
	api.PUT("/coffee-types-update/:id", handlers.UpdateCoffeeType)
	api.DELETE("/coffee-types-delete/:id", handlers.DeleteCoffeeType)
	api.GET("/coffee-types/:id", handlers.GetCoffeeTypeByID)

	// EmployeeType routes
	api.GET("/employee-types", handlers.GetEmployeeTypes)
	api.POST("/employee-types-add", handlers.CreateEmployeeType)
	api.PUT("/employee-types-update/:id", handlers.UpdateEmployeeType)
	api.DELETE("/employee-types-delete/:id", handlers.DeleteEmployeeType)
	api.GET("/employee-types/:id", handlers.GetEmployeeTypeByID)

	// Coffee routes
	api.GET("/coffees", handlers.GetCoffees)
	api.POST("/coffees-add", handlers.CreateCoffee)
	api.PUT("/coffees-update/:id", handlers.UpdateCoffee)
	api.DELETE("/coffees-delete/:id", handlers.DeleteCoffee)
	api.GET("/coffees/:id", handlers.GetCoffeeByID)

	// Employee routes
	api.GET("/employees", handlers.GetEmployees)
	api.POST("/employees-add", handlers.CreateEmployee)
	api.PUT("/employees-update/:id", handlers.UpdateEmployee)
	api.DELETE("/employees-delete/:id", handlers.DeleteEmployee)
	api.GET("/employees/:id", handlers.GetEmployeeByID)

	//Booking
	api.GET("/bookings", handlers.GetBookings)
	api.POST("/bookings-add", handlers.CreateBooking)
	api.PUT("/bookings-update/:id", handlers.CreateBooking)
	api.DELETE("/bookings-delete/:id", handlers.CreateBooking)
	api.GET("/bookings/:id", handlers.CreateBooking)

	//Table
	api.GET("/tables", handlers.GetTables)
	api.POST("/tables-add", handlers.CreateTable)
	api.PUT("/tables-update", handlers.UpdateTable)
	api.DELETE("/tables-delete", handlers.DeleteTable)
	api.GET("/tables/:id", handlers.GetTableById)

	//History Status
	// api.GET("/histories", handlers.GetBookingHistory)
	api.GET("/:id/histories", handlers.GetBookingHistoriesByID)

	//Absence Employee
	api.GET("/absence-employee", handlers.GetAbsenceEmployee)
	api.POST("/absence-employee-add", handlers.AddAbsenceEmployee)
	api.GET("/absence-employee/:id", handlers.GetAbsenceEmployeeByID)


	//Payment
	api.GET("/payments", handlers.GetPayment)
	api.POST("/payments-add", handlers.CreatePayment)
	api.PUT("/payments-update/:id", handlers.UpdatePayment)
	api.DELETE("/payments-delete/:id", handlers.DeletePayment)
	api.GET("/payments/:id", handlers.GetPaymentById)

	//MENU

	//MenuCake
	api.GET("/cake", handlers.GetCake)
	api.POST("/cake-add", handlers.CreateCake)
	api.PUT("/cake-update/:id", handlers.UpdateCake)
	api.DELETE("/cake-delete/:id", handlers.DeleteCake)

	//MenuSnack
	api.GET("/snack", handlers.GetSnack)
	api.POST("/snack-add", handlers.CreateSnack)
	api.PUT("/snack-update/:id", handlers.UpdateSnack)
	api.DELETE("/snack-delete/:id", handlers.DeleteSnack)

	//MenuWestern
	api.GET("/western", handlers.GetWestern)
	api.POST("/western-add", handlers.CreateWestern)
	api.PUT("/western-update/:id", handlers.UpdateWestern)
	api.DELETE("/western-delete/:id", handlers.DeleteWestern)

	//User
	api.GET("/user", handlers.GetAllUsers)
	api.GET("/user/:id", handlers.GetUserByID)

	//Menu/Foodcategory
	// api.GET("/food-categories", handlers.GetFoodCategories)
	// api.POST("/food-categories", handlers.CreateFoodCategory)
	// api.GET("/foods", handlers.GetFoods)
	// api.POST("/foods", handlers.CreateFood)

}

func SetupAuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	auth.POST("/register", handlers.RegisterUser)
	auth.POST("/login", handlers.LoginUser)
}


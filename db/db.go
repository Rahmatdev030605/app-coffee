package db

import (
	"fmt"
	"log"
	"os"

	"github.com/Rahmatdev030605/app-coffee/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSL := os.Getenv("DB_SSLMODE")

	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbName, dbPassword, dbSSL)

	var errDB error
	DB, errDB = gorm.Open("postgres", dbURL)
	if errDB != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", errDB))
	}

	// Auto Migrate
	DB.AutoMigrate(
		&models.Coffee{},
		&models.CoffeeType{},
		&models.Employee{},
		&models.EmployeeType{},
		&models.User{},
		&models.Booking{},
		&models.Table{},
		&models.HistoryBooking{},
		&models.EmployeeAbsence{},
		&models.Payment{},
		// &models.MenuFood{},
		&models.MenuCake{},
		&models.MenuSnack{},
		&models.MenuWestern{},
		)

	// Seed Initial data
		// models.SeedCoffeeType(DB)`
		// models.SeedEmployeeTypes(DB)

		// Delete Initial data
		// DB.Unscoped().Delete(&models.CoffeeType{})
		// DB.Unscoped().Delete(&models.EmployeeType{})
}

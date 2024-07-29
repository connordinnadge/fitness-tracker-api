package database

import (
	"fitness-tracker-api/config"
	"fitness-tracker-api/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate the models
	err = db.AutoMigrate(&models.User{}, &models.Workout{}, &models.Exercise{}, &models.ExerciseLog{})
	if err != nil {
		log.Fatal("Failed to auto-migrate database schema:", err)
	}

	DB = db
	fmt.Println("Database connection established")
}

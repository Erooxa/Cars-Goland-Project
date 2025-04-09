package services

import (
	"Cars/internal/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase функциясы SQLite мәліметтер базасына қосылады
// және Car мен User модельдерін миграциялайды.
func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("cars.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Мәліметтер базасына қосылу сәтсіз болды: ", err)
	}

	// Барлық модельдерге арналған миграция
	DB.AutoMigrate(&models.Car{}, &models.User{})
}

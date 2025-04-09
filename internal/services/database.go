package services

import (
	"Cars/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	// PostgreSQL-дің DSN (Data Source Name)
	dsn := "host=localhost user=postgres password=21012023 dbname=Carss port=5432 sslmode=disable TimeZone=Asia/Almaty"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ ПостгреSQL базасына қосыла алмадық: ", err)
	}

	// Миграция
	err = DB.AutoMigrate(&models.Car{}, &models.User{})
	if err != nil {
		log.Fatal("❌ Миграция қатесі: ", err)
	}

	log.Println("✅ PostgreSQL базасына сәтті қосылдық және модельдер миграцияланды.")
}

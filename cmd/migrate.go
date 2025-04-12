package main

import (
	"Cars/internal/models"
	"Cars/internal/services"
	"log"
)

func main() {
	services.ConnectDatabase()

	// Удаляем текущую таблицу cars
	if err := services.DB.Migrator().DropTable("cars"); err != nil {
		log.Printf("Ошибка при удалении таблицы cars: %v", err)
	} else {
		log.Println("Таблица cars успешно удалена")
	}

	// Создаем новую таблицу с правильной структурой
	if err := services.DB.AutoMigrate(&models.Car{}); err != nil {
		log.Fatalf("Ошибка при создании таблицы cars: %v", err)
	} else {
		log.Println("Таблица cars успешно создана с новой структурой")
	}

	log.Println("Миграция успешно завершена")
}

package main

import (
	"Cars/internal/models"
	"Cars/internal/services"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {
	services.ConnectDatabase()

	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("user123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Ошибка при хешировании пароля: %v", err)
	}

	// Создаем пользователя
	user := models.User{
		Email:    "user@example.com",
		Password: string(hashedPassword),
		RoleID:   2,
	}

	// Проверяем, существует ли уже пользователь с таким email
	var existingUser models.User
	if result := services.DB.Where("email = ?", user.Email).First(&existingUser); result.Error == nil {
		// Пользователь уже существует, обновляем его роль
		log.Printf("Пользователь %s уже существует, обновляем его роль на user (role_id=2)", user.Email)
		existingUser.RoleID = 2
		services.DB.Save(&existingUser)
	} else {
		// Создаем нового пользователя
		result := services.DB.Create(&user)
		if result.Error != nil {
			log.Fatalf("Ошибка при создании пользователя: %v", result.Error)
		}
		log.Printf("Пользователь успешно создан: %s", user.Email)
	}

	// Проверяем роли
	var role models.Role
	if result := services.DB.Where("role_id = ?", 2).First(&role); result.Error != nil {
		// Создаем роль пользователя
		role = models.Role{
			ID:   2,
			Name: "user",
		}
		if err := services.DB.Create(&role).Error; err != nil {
			log.Fatalf("Ошибка при создании роли user: %v", err)
		}
		log.Println("Роль user успешно создана")
	} else {
		log.Printf("Роль user уже существует: %s", role.Name)
	}

	log.Println("Операция завершена успешно")
}

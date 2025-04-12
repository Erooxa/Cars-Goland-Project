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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Ошибка при хешировании пароля: %v", err)
	}

	// Создаем админа
	admin := models.User{
		Email:    "admin@example.com",
		Password: string(hashedPassword),
		RoleID:   1,
	}

	// Проверяем, существует ли уже пользователь с таким email
	var existingUser models.User
	if result := services.DB.Where("email = ?", admin.Email).First(&existingUser); result.Error == nil {
		// Пользователь уже существует, обновляем его роль
		log.Printf("Пользователь %s уже существует, обновляем его роль на admin (role_id=1)", admin.Email)
		existingUser.RoleID = 1
		services.DB.Save(&existingUser)
	} else {
		// Создаем нового пользователя
		result := services.DB.Create(&admin)
		if result.Error != nil {
			log.Fatalf("Ошибка при создании администратора: %v", result.Error)
		}
		log.Printf("Администратор успешно создан: %s", admin.Email)
	}

	// Проверяем роли
	var role models.Role
	if result := services.DB.Where("role_id = ?", 1).First(&role); result.Error != nil {
		// Создаем роль админа
		role = models.Role{
			ID:   1,
			Name: "admin",
		}
		if err := services.DB.Create(&role).Error; err != nil {
			log.Fatalf("Ошибка при создании роли admin: %v", err)
		}
		log.Println("Роль admin успешно создана")
	} else {
		log.Printf("Роль admin уже существует: %s", role.Name)
	}

	log.Println("Операция завершена успешно")
}

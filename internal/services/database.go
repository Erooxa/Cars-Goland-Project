package services

import (
	"Cars/internal/migrations"
	"Cars/internal/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

// CreateRoles создает роли в базе данных
func CreateRoles() {
	var count int64
	DB.Model(&models.Role{}).Count(&count)
	if count > 0 {
		log.Println("Роли уже существуют, пропускаем создание")
		return
	}

	// Создаем базовые роли
	roles := []models.Role{
		{
			ID:   1, // ID=1 для роли admin
			Name: "admin",
		},
		{
			ID:   2, // ID=2 для роли user
			Name: "user",
		},
	}

	// Добавляем роли в базу данных
	for _, role := range roles {
		if err := DB.Create(&role).Error; err != nil {
			log.Printf("Ошибка при создании роли %s: %v", role.Name, err)
		} else {
			log.Printf("Роль %s успешно создана", role.Name)
		}
	}
}

// CreateDefaultUsers создает начальных пользователей (admin и user)
func CreateDefaultUsers() {
	var count int64
	DB.Model(&models.User{}).Count(&count)
	if count > 0 {
		log.Println("Пользователи уже существуют, пропускаем создание")
		return
	}

	// Хешируем пароли
	adminPass, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	userPass, _ := bcrypt.GenerateFromPassword([]byte("user123"), bcrypt.DefaultCost)

	// Создаем базовых пользователей
	users := []models.User{
		{
			ID:       1, // ID=1 для админа
			Email:    "admin@example.com",
			Password: string(adminPass),
			RoleID:   1, // Роль admin
		},
		{
			ID:       2, // ID=2 для пользователя
			Email:    "user@example.com",
			Password: string(userPass),
			RoleID:   2, // Роль user
		},
	}

	// Добавляем пользователей в базу данных
	for _, user := range users {
		if err := DB.Create(&user).Error; err != nil {
			log.Printf("Ошибка при создании пользователя %s: %v", user.Email, err)
		} else {
			log.Printf("Пользователь %s успешно создан", user.Email)
		}
	}
}

func ConnectDatabase() {
	// Get database configuration from environment variables
	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "21012023")
	dbname := getEnv("DB_NAME", "Carss")
	port := getEnv("DB_PORT", "5432")
	timeZone := getEnv("DB_TIMEZONE", "Asia/Almaty")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		host, user, password, dbname, port, timeZone)
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ ПостгреSQL базасына қосыла алмадық: ", err)
	}

	// Gormigrate пайдалану
	m := migrations.GetMigrations(DB)

	if err := m.Migrate(); err != nil {
		log.Fatal("❌ Миграция қатесі: ", err)
	}

	// Создаем роли и пользователей
	CreateRoles()
	CreateDefaultUsers()

	log.Println("✅ PostgreSQL базасына сәтті қосылдық және миграциялар орындалды.")
}

func RollbackMigrations() {
	m := migrations.GetMigrations(DB)
	if err := m.RollbackLast(); err != nil {
		log.Fatal("❌ Миграцияны қайтару қатесі: ", err)
	}
	log.Println("✅ Миграциялар сәтті қайтарылды.")
}

// Helper function to get environment variables with default fallback
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

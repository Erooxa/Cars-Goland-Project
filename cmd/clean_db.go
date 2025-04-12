package main

import (
	"Cars/internal/services"
	"log"
)

func main() {
	services.ConnectDatabase()

	// Удаляем все refresh токены
	if err := services.DB.Exec("DELETE FROM refresh_tokens").Error; err != nil {
		log.Printf("Ошибка при удалении refresh токенов: %v", err)
	} else {
		log.Println("Refresh токены успешно удалены")
	}

	// Удаляем всех пользователей
	if err := services.DB.Exec("DELETE FROM users").Error; err != nil {
		log.Printf("Ошибка при удалении пользователей: %v", err)
	} else {
		log.Println("Пользователи успешно удалены")
	}

	// Удаляем все роли
	if err := services.DB.Exec("DELETE FROM role").Error; err != nil {
		log.Printf("Ошибка при удалении ролей: %v", err)
	} else {
		log.Println("Роли успешно удалены")
	}

	// Сбрасываем автоинкремент в таблицах
	if err := services.DB.Exec("ALTER SEQUENCE users_user_id_seq RESTART WITH 1").Error; err != nil {
		log.Printf("Ошибка при сбросе последовательности users: %v", err)
	}

	if err := services.DB.Exec("ALTER SEQUENCE role_role_id_seq RESTART WITH 1").Error; err != nil {
		log.Printf("Ошибка при сбросе последовательности roles: %v", err)
	}

	if err := services.DB.Exec("ALTER SEQUENCE refresh_tokens_id_seq RESTART WITH 1").Error; err != nil {
		log.Printf("Ошибка при сбросе последовательности refresh_tokens: %v", err)
	}

	log.Println("База данных успешно очищена")
}

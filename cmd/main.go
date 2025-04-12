package main

import (
	"Cars/internal/controllers"
	"Cars/internal/middleware"
	"Cars/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	services.ConnectDatabase()
	router := gin.Default()

	// Публичные маршруты
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)

	// Защищенные маршруты для пользователей (чтение)
	userRoutes := router.Group("/api")
	userRoutes.Use(middleware.AuthMiddleware())
	{
		userRoutes.GET("/cars", controllers.GetCars)
		userRoutes.GET("/cars/:id", controllers.GetCarByID)
		userRoutes.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Аутентификация прошла успешно"})
		})
	}

	// Защищенные маршруты для администраторов (полный CRUD)
	adminRoutes := router.Group("/api/admin")
	adminRoutes.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		adminRoutes.POST("/cars", controllers.CreateCar)
		adminRoutes.PUT("/cars/:id", controllers.UpdateCar)
		adminRoutes.DELETE("/cars/:id", controllers.DeleteCar)
	}

	router.Run(":8080")
}

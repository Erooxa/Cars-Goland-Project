package controllers

import (
	"Cars/internal/models"
	"Cars/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	router.POST("/signup", Signup)
	router.POST("/login", Login)
}

func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Устанавливаем роль по умолчанию, если она не указана
	if user.RoleID == 0 {
		user.RoleID = 2 // По умолчанию роль user
	}

	if err := services.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно зарегистрирован"})
}

func Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.LoginUser(loginData.Email, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := services.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка генерации токена"})
		return
	}

	// Получаем информацию о роли
	var role models.Role
	services.DB.First(&role, user.RoleID)

	c.JSON(http.StatusOK, gin.H{
		"token":     token,
		"user_id":   user.ID,
		"role_id":   user.RoleID,
		"role_name": role.Name,
	})
}

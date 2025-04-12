package middleware

import (
	"Cars/internal/models"
	"Cars/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoleMiddleware проверяет роль пользователя
func RoleMiddleware(roleID uint) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		var user models.User
		if err := services.DB.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// Проверяем, есть ли у пользователя нужная роль
		if user.RoleID != roleID {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// AdminMiddleware проверяет, что пользователь имеет роль администратора (role_id=1)
func AdminMiddleware() gin.HandlerFunc {
	return RoleMiddleware(1)
}

package controllers

import (
	"Cars/internal/models"
	"Cars/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Роуттарды тіркеу
func RegisterCarRoutes(router *gin.Engine) {
	router.GET("/cars", GetCars)
	router.GET("/cars/:id", GetCarByID)
	router.POST("/cars", CreateCar)
	router.PUT("/cars/:id", UpdateCar)
	router.DELETE("/cars/:id", DeleteCar)
}

// Барлық көліктерді алу
func GetCars(c *gin.Context) {
	cars := services.GetCars()
	c.JSON(http.StatusOK, cars)
}

// Көлікті ID бойынша алу
func GetCarByID(c *gin.Context) {
	id := c.Param("id")
	car, err := services.GetCarByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}
	c.JSON(http.StatusOK, car)
}

// Көлік қосу
func CreateCar(c *gin.Context) {
	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	services.CreateCar(car)
	c.JSON(http.StatusCreated, car)
}

// Көлікті жаңарту
func UpdateCar(c *gin.Context) {
	id := c.Param("id")
	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	services.UpdateCar(id, car)
	c.JSON(http.StatusOK, car)
}

// Көлікті өшіру
func DeleteCar(c *gin.Context) {
	id := c.Param("id")
	services.DeleteCar(id)
	c.JSON(http.StatusOK, gin.H{"message": "Car deleted"})
}

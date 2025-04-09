package controllers

import (
	"Cars/internal/models"
	"Cars/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterCarRoutes(router *gin.Engine) {
	router.GET("/cars", getCars)
	router.POST("/cars", createCar)
	router.PUT("/cars/:id", updateCar)
	router.DELETE("/cars/:id", deleteCar)
}

func getCars(c *gin.Context) {
	cars := services.GetCars()
	c.JSON(http.StatusOK, cars)
}

func createCar(c *gin.Context) {
	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	services.CreateCar(car)
	c.JSON(http.StatusCreated, car)
}

func updateCar(c *gin.Context) {
	var car models.Car
	id := c.Param("id")
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	services.UpdateCar(id, car)
	c.JSON(http.StatusOK, car)
}

func deleteCar(c *gin.Context) {
	id := c.Param("id")
	services.DeleteCar(id)
	c.JSON(http.StatusOK, gin.H{"message": "Car deleted"})
}

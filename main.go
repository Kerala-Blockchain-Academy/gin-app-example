package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarId        string `json:"carId" binding:"required"`
	Make         string `json:"make" binding:"required"`
	Model        string `json:"model" binding:"required"`
	Color        string `json:"color" binding:"required"`
	Date         string `json:"dateOfManufacture" binding:"required"`
	Manufacturer string `json:"manufacturerName" binding:"required"`
}

func main() {
	router := gin.Default()
	router.Static("/public", "./public")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Auto App",
		})
	})

	router.POST("/api/car", func(ctx *gin.Context) {
		var res Car
		if err := ctx.ShouldBind(&res); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		ctx.JSON(http.StatusOK, res)
	})

	router.GET("/api/car/:id", func(ctx *gin.Context) {
		carId := ctx.Param("id")
		car := Car{CarId: carId, Make: "Make", Model: "Model", Color: "Color", Date: "Date", Manufacturer: "Manufacturer"}
		ctx.JSON(http.StatusOK, gin.H{"data": car})
	})

	router.Run("localhost:8080")
}

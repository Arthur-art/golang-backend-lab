package main

import (
	"go-api/controller"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	//Usecase
	ProductUsecase := usecase.ProductUsecase{}
	// Products Controller
	ProductController := controller.NewProductController(ProductUsecase)
	server.GET("/products", ProductController.GetAllProducts)

	// User Controller
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Run(":8000")
}
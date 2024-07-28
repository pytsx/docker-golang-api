package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	usecase "go-api/useCase"

	"github.com/gin-gonic/gin"
)


func main() {

	server := gin.Default()

	bdConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(bdConnection)
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})


	server.GET("/products", ProductController.GetProducts)

	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)

	server.Run(":8000")
}
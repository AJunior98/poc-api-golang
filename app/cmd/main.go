package main

import (
	"go-api/config"
	"go-api/controller"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	//Camada usecase
	ProductUseCase := usecase.NewProductUSecase(ProductRepository)

	//Camada Controller
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/products", ProductController.GetProducts)
	server.POST("/products", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)

	server.Run(":8000")
}

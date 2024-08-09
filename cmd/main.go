package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)

	ProductUsecase := usecase.NewProductUseCase(ProductRepository)

	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/products/:id_product", ProductController.GetProductById)

	server.GET("/products", ProductController.GetProducts)

	server.POST("/products", ProductController.CreateProduct)

	server.PUT("/products/:id_product", ProductController.UpdateProduct)

	server.DELETE("/products/:id_product", ProductController.DeleteProduct)

	server.Run(":8000")
}

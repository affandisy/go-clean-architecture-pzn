package main

import (
	"go-clean-architecture-pzn/config"
	"go-clean-architecture-pzn/controller"
	"go-clean-architecture-pzn/exception"
	"go-clean-architecture-pzn/repository"
	"go-clean-architecture-pzn/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configuration := config.New()
	database := config.NewMongoDatabase(configuration)

	productRepository := repository.NewProductRepository(database)
	productService := service.NewProductService(&productRepository)
	productController := controller.NewProductController(&productService)

	app := fiber.New()
	productController.Route(app)

	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}

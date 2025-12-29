package main

import (
	"go-clean-architecture-pzn/config"
	"go-clean-architecture-pzn/controller"
	"go-clean-architecture-pzn/exception"
	"go-clean-architecture-pzn/repository"
	"go-clean-architecture-pzn/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Setup Configuration
	configuration := config.New()
	database := config.NewMongoDatabase(configuration)

	// Setup Repository
	productRepository := repository.NewProductRepository(database)
	// Setup Service
	productService := service.NewProductService(&productRepository)
	// Setup Controller
	productController := controller.NewProductController(&productService)

	// app := fiber.New()

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	productController.Route(app)

	// Start App
	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}

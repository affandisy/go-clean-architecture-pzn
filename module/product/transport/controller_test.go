package controller

import (
	configMongo "go-clean-architecture-pzn/app"
	productRepo "go-clean-architecture-pzn/module/product/datasource"
	productSvc "go-clean-architecture-pzn/module/product/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func createTestApp() *fiber.App {
	var app = fiber.New(configMongo.NewFiberConfig())
	app.Use(recover.New())
	productController.Route(app)

	return app
}

var configuration = configMongo.NewConfigMongo("../.env.test")

var database = configMongo.NewMongoDatabase(configuration)

var productRepository = productRepo.NewProductRepository(database)

var productService = productSvc.NewProductService(&productRepository)

var productController = NewProductController(&productService)

var app = createTestApp()

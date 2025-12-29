package main

import (
	"context"
	"fmt"
	"go-clean-architecture-pzn/app"
	"go-clean-architecture-pzn/domain/entity"
	"go-clean-architecture-pzn/exception"
	productRepo "go-clean-architecture-pzn/module/product/datasource"
	controller "go-clean-architecture-pzn/module/product/transport"
	productSvc "go-clean-architecture-pzn/module/product/usecase"
	"os"

	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// go:embed migrations
// var fs embed.FS

func main() {
	config, err := app.NewConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	db, err := app.NewDatabase(config)
	if err != nil {
		panic(fmt.Errorf("Fatal error database: %w \n", err))
	}

	// Run migrate if argument is migrate
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		err := RunMigration(db)
		if err != nil {
			panic(fmt.Errorf("Error run migration: %w \n", err))
		}
		return
	}

	transaction := db.Begin()
	err = SaveUser(transaction)
	if err != nil {
		fmt.Printf("Error save user: %s \n", err.Error())
	}

	connection, _ := db.DB()
	connection.Close()

	// Product Mongo Setup
	cfgMongo := app.NewConfigMongo(".env")

	dbMongo := app.NewMongoDatabase(cfgMongo)
	defer dbMongo.Client().Disconnect(context.Background())

	fiberApp := setupHTTPMongoProduct(dbMongo)

	exception.PanicIfNeeded(fiberApp.Listen(":3000"))
}

func setupHTTPMongoProduct(db *mongo.Database) *fiber.App {
	// Setup Repository
	productRepository := productRepo.NewProductRepository(db)
	// Setup Service
	productService := productSvc.NewProductService(&productRepository)
	// Setup Controller
	productController := controller.NewProductController(&productService)

	// Setup Fiber
	app := fiber.New(app.NewFiberConfig())
	app.Use(recover.New())

	productController.Route(app)

	return app
}

func RunMigration(db *gorm.DB) error {
	return nil
}

func SaveUser(db *gorm.DB) error {
	defer db.Rollback()

	result := db.Create(&entity.User{
		ID:       "2",
		Username: "rahasia",
		Password: "rahasia",
		Name:     "belajar",
	})
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Save user 1 success")

	result = db.Create(&entity.User{
		ID:       "1",
		Username: "rahasia belajar",
		Password: "rahasia",
		Name:     "belajar golang",
	})
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Save user 2 success")

	result = db.Commit()
	if result.Error != nil {
		return result.Error
	}

	return nil
}

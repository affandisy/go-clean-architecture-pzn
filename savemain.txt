package main

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"go-clean-architecture-pzn/app"
	"go-clean-architecture-pzn/domain/entity"
	"go-clean-architecture-pzn/exception"
	productRepo "go-clean-architecture-pzn/module/product/datasource"
	controller "go-clean-architecture-pzn/module/product/transport"
	productSvc "go-clean-architecture-pzn/module/product/usecase"

	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// go:embed migrations
var fs embed.FS

func main() {
	config, err := app.NewConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	// db, err := app.NewDatabase(config)
	// if err != nil {
	// 	panic(fmt.Errorf("Fatal error database: %w \n", err))
	// }

	log := app.NewLogger(config)
	log.Info("Start Application")

	db, err := app.NewDatabase(config, log)
	if err != nil {
		panic(fmt.Errorf("Fatal Error Database: %w\n", err))
	}

	// Run migrate if argument is migrate
	// if len(os.Args) > 1 && os.Args[1] == "migrate" {
	// 	err := RunMigration(db)
	// 	if err != nil {
	// 		panic(fmt.Errorf("Error run migration: %w \n", err))
	// 	}
	// 	return
	// }

	err = RunMigration(db)
	if err != nil {
		panic(fmt.Errorf("Error run migration: %w", err))
	}

	transaction := db.Begin()
	// err = SaveUser(transaction)
	// if err != nil {
	// 	fmt.Printf("Error save user: %s \n", err.Error())
	// }

	users, err := FindUserWithContact(transaction)
	if err != nil {
		fmt.Printf("Error find user: %s \n", err.Error())
	}
	log.Info(users)

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
	dbSql, err := db.DB()
	if err != nil {
		return err
	}

	location, err := iofs.New(fs, "migrations")
	if err != nil {
		return err
	}

	driver, err := mysql.WithInstance(dbSql, &mysql.Config{})
	if err != nil {
		return err
	}

	migration, err := migrate.NewWithInstance("iofs", location, "mysql", driver)
	if err != nil {
		return err
	}

	err = migration.Up()
	if err != nil {
		// return err
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		} else {
			return err
		}
	}

	return nil
}

func FindUserWithContact(db *gorm.DB) ([]entity.User, error) {
	var users []entity.User
	err := db.Model(&entity.User{}).Preload("Contacts").Find(&users).Error
	return users, err
}

func SaveContact(db *gorm.DB) error {
	defer db.Rollback()

	err := db.Create(&entity.Contact{
		ID:        "1",
		FirstName: "Affandi",
		LastName:  "Syihabuddin",
		Email:     "affandi@gmail.com",
		Phone:     "080181008",
		UserId:    "1",
	}).Error
	if err != nil {
		return err
	}

	err = db.Create(&entity.Contact{
		ID:        "2",
		FirstName: "Affandiss",
		LastName:  "Syihabuddinss",
		Email:     "sss@gmail.com",
		Phone:     "080181008",
		UserId:    "2",
	}).Error
	if err != nil {
		return err
	}

	err = db.Commit().Error
	if err != nil {
		return err
	}

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

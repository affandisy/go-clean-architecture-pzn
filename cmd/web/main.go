package main

import (
	"fmt"
	"go-clean-architecture-pzn/config"
	"go-clean-architecture-pzn/controller"
	"go-clean-architecture-pzn/internal"
	"go-clean-architecture-pzn/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	viperConfig, err := config.New()
	if err != nil {
		panic(fmt.Errorf("Fatal error viperConfig file: %w", err))
	}

	log := internal.NewLogger(viperConfig)
	log.Info("Start application")

	db, err := internal.NewDatabase(viperConfig, log)
	if err != nil {
		panic(fmt.Errorf("Fatal error database: %w", err))
	}

	validator := internal.NewValidator(viperConfig)

	webPort := viperConfig.GetInt("web.port")
	app := NewFiber(viperConfig)

	// regis route
	// err = route.User(app)
	// if err != nil {
	// 	panic(fmt.Errorf("Fatal error route user: %w", err))
	// }

	// err = route.Contact(app)
	// if err != nil {
	// 	panic(fmt.Errorf("Fatal error route contact: %w", err))
	// }

	// err = route.Address(app)
	// if err != nil {
	// 	panic(fmt.Errorf("Fatal error route address: %w", err))
	// }

	// register controller
	userController := controller.NewUserController(db, validator, log)
	// userController.Routes(app)
	contactController := controller.NewContactController(db, validator, log)
	// contactController.Routes(app)
	addressController := controller.NewAddressController(db, validator, log)
	// addressController.Routes(app)

	// guest routes
	app.Post("/api/users", userController.Register)
	app.Post("/api/users/_login", userController.Login)

	// auth routes
	app.Use(middleware.NewAuth(db, log))
	app.Delete("/api/users", userController.Logout)
	app.Patch("/api/users/_current", userController.Update)
	app.Get("/api/users/_current", userController.Current)

	app.Get("/api/contacts", contactController.List)
	app.Post("/api/contacts", contactController.Create)
	app.Put("/api/contacts", contactController.Update)
	app.Get("/api/contacts/:contactId", contactController.Get)
	app.Delete("/api/contacts", contactController.Delete)

	app.Get("/api/contacts/:contactId/addresses", addressController.List)
	app.Post("/api/contacts/:contactId/addresses", addressController.Create)
	app.Put("/api/contacts/:contactId/addresses/:addressId", addressController.Update)
	app.Get("/api/contacts/:contactId/addresses/:addressId", addressController.Get)
	app.Delete("/api/contacts/:contactId/addresses/addressId", addressController.Delete)

	// Start server
	err = app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		panic(err)
	}
}

func NewFiber(config *viper.Viper) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName: config.Get("app.name").(string),
	})

	return app
}

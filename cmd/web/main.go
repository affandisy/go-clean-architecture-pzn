package main

import (
	"fmt"
	"go-clean-architecture-pzn/controller"
	"go-clean-architecture-pzn/internal"
	"go-clean-architecture-pzn/middleware"
	"go-clean-architecture-pzn/usecase"
)

func main() {
	// viperConfig, err := internal.New()
	config, err := internal.NewViper()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w", err))
	}

	log := internal.NewLogger(config)
	log.Info("Start application")

	db, err := internal.NewDatabase(config, log)
	if err != nil {
		panic(fmt.Errorf("Fatal error database: %w", err))
	}

	// validator := internal.NewValidator(viperConfig)
	validate := internal.NewValidator(config)

	webPort := config.GetInt("web.port")

	app := internal.NewFiber(config)

	routeConfig := internal.RouteConfig{
		App: app,
		// UserController:    controller.NewUserController(db, validate, log),
		UserController: controller.NewUserController(usecase.NewUserUseCase(db, log, validate), log),
		// ContactController: controller.NewContactController(db, validate, log),
		ContactController: controller.NewContactController(usecase.NewContactUseCase(db, log, validate), log),
		// AddressController: controller.NewAddressController(db, validate, log),
		AddressController: controller.NewAddressController(usecase.NewAddressUseCase(db, log, validate), log),
		AuthMiddleware:    middleware.NewAuth(db, log),
	}

	routeConfig.Setup()

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
	// userController := controller.NewUserController(db, validator, log)
	// userController.Routes(app)
	// contactController := controller.NewContactController(db, validator, log)
	// contactController.Routes(app)
	// addressController := controller.NewAddressController(db, validator, log)
	// addressController.Routes(app)

	// guest routes
	// app.Post("/api/users", userController.Register)
	// app.Post("/api/users/_login", userController.Login)

	// // auth routes
	// app.Use(middleware.NewAuth(db, log))
	// app.Delete("/api/users", userController.Logout)
	// app.Patch("/api/users/_current", userController.Update)
	// app.Get("/api/users/_current", userController.Current)

	// app.Get("/api/contacts", contactController.List)
	// app.Post("/api/contacts", contactController.Create)
	// app.Put("/api/contacts", contactController.Update)
	// app.Get("/api/contacts/:contactId", contactController.Get)
	// app.Delete("/api/contacts", contactController.Delete)

	// app.Get("/api/contacts/:contactId/addresses", addressController.List)
	// app.Post("/api/contacts/:contactId/addresses", addressController.Create)
	// app.Put("/api/contacts/:contactId/addresses/:addressId", addressController.Update)
	// app.Get("/api/contacts/:contactId/addresses/:addressId", addressController.Get)
	// app.Delete("/api/contacts/:contactId/addresses/:addressId", addressController.Delete)

	// Start server
	err = app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		// panic(err)
		log.Fatal(err)
	}
}

// func NewFiber(config *viper.Viper) *fiber.App {
// 	var app = fiber.New(fiber.Config{
// 		// AppName:      config.Get("app.name").(string),
// 		AppName:      config.GetString("app.name"),
// 		ErrorHandler: NewErrorHandler(),
// 		Prefork:      config.GetBool("web.prefork"),
// 	})

// 	return app
// }

// func NewErrorHandler() fiber.ErrorHandler {
// 	return func(ctx *fiber.Ctx, err error) error {
// 		code := fiber.StatusInternalServerError
// 		if e, ok := err.(*fiber.Error); ok {
// 			code = e.Code
// 		}

// 		return ctx.Status(code).JSON(fiber.Map{
// 			"errors": err.Error(),
// 		})
// 	}
// }

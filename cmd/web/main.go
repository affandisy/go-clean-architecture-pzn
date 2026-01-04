package main

import (
	"fmt"

	config "go-clean-architecture-pzn/internal/config"
)

func main() {
	// viperConfig, err := internal.New()
	// viperConfig, err := config.NewViper()
	// if err != nil {
	// 	panic(fmt.Errorf("Fatal error config file: %w", err))
	// }

	viperConfig := config.NewViper()

	log := config.NewLogger(viperConfig)
	// log.Info("Start application")

	// db, err := config.NewDatabase(viperConfig, log)
	// if err != nil {
	// 	panic(fmt.Errorf("Fatal error database: %w", err))
	// }
	db := config.NewDatabase(viperConfig, log)

	// validator := internal.NewValidator(viperConfig)
	validate := config.NewValidator(viperConfig)

	// webPort := viperConfig.GetInt("web.port")

	app := config.NewFiber(viperConfig)

	// producer, err := config.NewKafkaProducer(viperConfig)
	// if err != nil {
	// 	panic(fmt.Errorf("Fatal error kafka producer: %w", err))
	// }

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   viperConfig,
		// Producer: producer,
	})

	// routeConfig := route.RouteConfig{
	// 	App: app,
	// 	// UserController:    controller.NewUserController(db, validate, log),
	// 	UserController: http.NewUserController(usecase.NewUserUseCase(db, log, validate), log),
	// 	// ContactController: controller.NewContactController(db, validate, log),
	// 	ContactController: http.NewContactController(usecase.NewContactUseCase(db, log, validate), log),
	// 	// AddressController: controller.NewAddressController(db, validate, log),
	// 	AddressController: http.NewAddressController(usecase.NewAddressUseCase(db, log, validate), log),
	// 	AuthMiddleware:    middleware.NewAuth(db, log),
	// }

	// routeConfig.Setup()

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
	// webPort := viperConfig.GetInt("web.port")
	// err := app.Listen(fmt.Sprintf(":%d", webPort))
	webPort := viperConfig.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		// panic(err)
		// log.Fatal(err)
		log.Fatalf("Failed to start server: %v", err)
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

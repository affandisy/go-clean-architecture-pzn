package test

import (
	"fmt"
	config "go-clean-architecture-pzn/internal/config"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var app *fiber.App

var db *gorm.DB

var viperConfig *viper.Viper

var log *logrus.Logger

var validate *validator.Validate

func init() {
	var err error

	viperConfig, err = config.NewViper()
	if err != nil {
		panic(fmt.Errorf("Fatal error viperConfig file: %w", err))
	}

	log = config.NewLogger(viperConfig)
	validate = config.NewValidator(viperConfig)
	app = config.NewFiber(viperConfig)

	db, err = config.NewDatabase(viperConfig, log)
	if err != nil {
		panic(fmt.Errorf("Fatal error database: %w", err))
	}

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

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   viperConfig,
	})

	// routeConfig.Setup()
}

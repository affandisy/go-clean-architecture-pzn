package test

import (
	"fmt"
	"go-clean-architecture-pzn/internal"
	http "go-clean-architecture-pzn/internal/delivery/http"
	"go-clean-architecture-pzn/internal/delivery/http/middleware"
	"go-clean-architecture-pzn/internal/usecase"

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

	viperConfig, err = internal.NewViper()
	if err != nil {
		panic(fmt.Errorf("Fatal error viperConfig file: %w", err))
	}

	log = internal.NewLogger(viperConfig)
	validate = internal.NewValidator(viperConfig)
	app = internal.NewFiber(viperConfig)

	db, err = internal.NewDatabase(viperConfig, log)
	if err != nil {
		panic(fmt.Errorf("Fatal error database: %w", err))
	}

	routeConfig := internal.RouteConfig{
		App: app,
		// UserController:    controller.NewUserController(db, validate, log),
		UserController: http.NewUserController(usecase.NewUserUseCase(db, log, validate), log),
		// ContactController: controller.NewContactController(db, validate, log),
		ContactController: http.NewContactController(usecase.NewContactUseCase(db, log, validate), log),
		// AddressController: controller.NewAddressController(db, validate, log),
		AddressController: http.NewAddressController(usecase.NewAddressUseCase(db, log, validate), log),
		AuthMiddleware:    middleware.NewAuth(db, log),
	}

	routeConfig.Setup()
}

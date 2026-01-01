package main

import (
	"fmt"
	"go-clean-architecture-pzn/cmd/web/route"
	"go-clean-architecture-pzn/config"
	"go-clean-architecture-pzn/internal"

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

	_, err = internal.NewDatabase(viperConfig, log)
	if err != nil {
		panic(fmt.Errorf("Fatal error database: %w", err))
	}

	webPort := viperConfig.GetInt("web.port")
	app := NewFiber(viperConfig)

	// regis route
	err = route.User(app)
	if err != nil {
		panic(fmt.Errorf("Fatal error route user: %w", err))
	}

	err = route.Contact(app)
	if err != nil {
		panic(fmt.Errorf("Fatal error route contact: %w", err))
	}

	err = route.Address(app)
	if err != nil {
		panic(fmt.Errorf("Fatal error route address: %w", err))
	}

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

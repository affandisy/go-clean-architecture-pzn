package transport

import (
	"go-clean-architecture-pzn/module/contact"

	"github.com/gofiber/fiber/v2"
)

type ContactTransportHTTP struct {
	ContactUsecase contact.ContactUseCase
}

func NewContactTransportHTTP(contactUsecase contact.ContactUseCase) *ContactTransportHTTP {
	return &ContactTransportHTTP{
		ContactUsecase: contactUsecase,
	}
}

func (receiver *ContactTransportHTTP) Create(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "Hello world",
	})
}

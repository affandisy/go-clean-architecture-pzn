package controller

import (
	"go-clean-architecture-pzn/internal/exception"
	"go-clean-architecture-pzn/internal/model"
	productSvc "go-clean-architecture-pzn/module/product/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductController struct {
	ProductService productSvc.ProductService
}

func NewProductController(productService *productSvc.ProductService) ProductController {
	return ProductController{ProductService: *productService}
}

func (controller *ProductController) Route(app *fiber.App) {
	app.Post("/api/products", controller.Create)
	app.Get("/api/products", controller.List)
}

func (controller *ProductController) Create(c *fiber.Ctx) error {
	var request model.CreateProductRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.ProductService.Create(request)
	return c.JSON(model.WebResponseHTTP{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *ProductController) List(c *fiber.Ctx) error {
	responses := controller.ProductService.List()

	return c.JSON(model.WebResponseHTTP{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}

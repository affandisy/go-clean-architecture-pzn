package usecase

import (
	"go-clean-architecture-pzn/domain/entity"
	"go-clean-architecture-pzn/model"
	productRepo "go-clean-architecture-pzn/module/product"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func NewProductService(productRepository *productRepo.ProductRepository) ProductService {
	return &productServiceImpl{
		ProductRepository: *productRepository,
	}
}

type productServiceImpl struct {
	ProductRepository productRepo.ProductRepository
}

func (service *productServiceImpl) Create(request model.CreateProductRequest) (response model.CreateProductResponse) {
	validation.Validate(request)

	product := entity.Product{
		Id:       request.Id,
		Name:     request.Name,
		Price:    request.Price,
		Quantity: request.Quantity,
	}

	service.ProductRepository.Insert(product)

	response = model.CreateProductResponse{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}

	return response
}

func (service *productServiceImpl) List() (responses []model.GetProductResponse) {
	products := service.ProductRepository.FindAll()

	for _, product := range products {
		responses = append(responses, model.GetProductResponse{
			Id:       product.Id,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
		})
	}

	return responses
}

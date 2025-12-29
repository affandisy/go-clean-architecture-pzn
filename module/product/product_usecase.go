package product

import "go-clean-architecture-pzn/model"

type ProductUseCase interface {
	Create(request model.CreateProductRequest) (response model.CreateProductResponse)
	List() (responses []model.GetProductResponse)
}

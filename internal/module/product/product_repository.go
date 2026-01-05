package product

import "go-clean-architecture-pzn/internal/entity"

type ProductRepository interface {
	Insert(product entity.Product)
	FindAll() (products []entity.Product)
	DeleteAll()
}

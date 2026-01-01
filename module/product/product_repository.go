package product

import "go-clean-architecture-pzn/entity"

type ProductRepository interface {
	Insert(product entity.Product)
	FindAll() (products []entity.Product)
	DeleteAll()
}

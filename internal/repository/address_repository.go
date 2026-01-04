package repository

import (
	"go-clean-architecture-pzn/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AddressRepository struct {
	Log *logrus.Logger
}

func (r *AddressRepository) Create(db *gorm.DB, address *entity.Address) error {
	return db.Create(address).Error
}

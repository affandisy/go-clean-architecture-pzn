package repository

import (
	"go-clean-architecture-pzn/internal/entity"

	"github.com/sirupsen/logrus"
)

type ContactRepository struct {
	// DB  *gorm.DB
	Repository[entity.Contact]
	Log *logrus.Logger
}

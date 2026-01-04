package repository

import (
	"go-clean-architecture-pzn/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// UserRepository is an interface for user repository contract
// type UserRepository interface {
// 	// FindById(ctx context.Context, id string) (*entity.User, error)
// 	// FindAll(ctx context.Context) ([]entity.User, error)
// 	// Save(ctx context.Context, user *entity.User) error
// 	// Update(ctx context.Context, user *entity.User) error
// 	// Delete(ctx context.Context, user *entity.User) error
// 	// DeleteById(ctx context.Context, id string) error
// }

type UserRepository struct {
	// DB  *gorm.DB
	Repository[entity.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}

func (r *UserRepository) FindByToken(db *gorm.DB, user *entity.User, token string) error {
	return db.Where("token = ?", token).First(user).Error
}

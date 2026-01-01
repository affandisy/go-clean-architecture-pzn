package usecase

import (
	"context"
	"go-clean-architecture-pzn/module/user"

	"github.com/go-playground/validator/v10"
)

type userUserCase struct {
	Validate       *validator.Validate
	UserRepository user.UserRepository
}

func (u *userUserCase) Create(ctx context.Context) error {
	return nil
}

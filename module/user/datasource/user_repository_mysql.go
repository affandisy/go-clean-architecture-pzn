package datasource

import (
	"context"
	"go-clean-architecture-pzn/domain/entity"
	"go-clean-architecture-pzn/module/user"
)

type userRepositoryMySQL struct {
}

func NewMySQL() user.UserRepository {
	return &userRepositoryMySQL{}
}

func (u *userRepositoryMySQL) FindById(ctx context.Context, id string) (*entity.User, error) {
	return nil, nil
}

func (u *userRepositoryMySQL) FindAll(ctx context.Context) (*[]entity.User, error) {
	return nil, nil
}

func (u *userRepositoryMySQL) Save(ctx context.Context, user *entity.User) error {
	return nil
}

func (u *userRepositoryMySQL) Update(ctx context.Context, user *entity.User) error {
	return nil
}

func (u *userRepositoryMySQL) Delete(ctx context.Context, user *entity.User) error {
	return nil
}

func (u *userRepositoryMySQL) DeleteById(ctx context.Context, id string) error {
	return nil
}

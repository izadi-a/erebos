package user

import (
	"context"

	"erebos/internal/domain/user"
)

type UserUseCases struct {
	service *user.UserService
}

func NewUserUseCases(service *user.UserService) *UserUseCases {
	return &UserUseCases{service: service}
}

func (uc *UserUseCases) CreateUser(ctx context.Context, name, email, password string) (*user.User, error) {
	return uc.service.Create(ctx, name, email, password)
}

func (uc *UserUseCases) FindUserByID(ctx context.Context, id string) (*user.User, error) {
	return uc.service.FindByID(ctx, id)
}

func (uc *UserUseCases) FindAllUsers(ctx context.Context) ([]*user.User, error) {
	return uc.service.FindAll(ctx)
}

func (uc *UserUseCases) ChangePassword(ctx context.Context, id, password string) error {
	return uc.service.ChangePassword(ctx, id, password)
}

func (uc *UserUseCases) DeleteUser(ctx context.Context, id string) error {
	return uc.service.Delete(ctx, id)
}

package user

import (
	"context"

	"erebos/internal/application/base"
	"erebos/internal/domain/user"
)

type UserService struct {
	base.BaseService
	service *user.UserService
}

func NewUserUseCases(service *user.UserService) *UserService {
	return &UserService{service: service}
}

func (uc *UserService) CreateUser(ctx context.Context, name, email, password string) (*user.User, error) {
	uc.Log("Creating user: " + name)
	return uc.service.Create(ctx, name, email, password)
}

func (uc *UserService) FindUserByID(ctx context.Context, id string) (*user.User, error) {
	uc.Log("Finding user by ID: " + id)
	return uc.service.FindByID(ctx, id)
}

func (uc *UserService) FindAllUsers(ctx context.Context) ([]*user.User, error) {
	uc.Log("Finding all users")
	return uc.service.FindAll(ctx)
}

func (uc *UserService) ChangePassword(ctx context.Context, id, password string) error {
	uc.Log("Changing password for user ID: " + id)
	return uc.service.ChangePassword(ctx, id, password)
}

func (uc *UserService) DeleteUser(ctx context.Context, id string) error {
	uc.Log("Deleting user ID: " + id)
	return uc.service.Delete(ctx, id)
}

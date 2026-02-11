package user

import (
	"context"

	"erebos/internal/application/base"
	"erebos/internal/domain/user"
	"erebos/internal/ports"
)

type CreateUserUseCase struct {
	base.BaseService
	Tx      base.TxManager
	service *user.UserService
}

type CreateUserDTO struct {
	Id    string
	Name  string
	Email string
}

func NewUserUseCases(service *user.UserService) *CreateUserUseCase {
	return &CreateUserUseCase{service: service}
}

func (uc *CreateUserUseCase) CreateUser(ctx context.Context, name, email, password string) (*CreateUserDTO, error) {
	uc.Log("Creating user: " + name)
	var result *CreateUserDTO

	err := uc.Tx.WithTransaction(ctx, ports.Write, func(txCtx context.Context) error {
		u, err := uc.service.Create(txCtx, name, email, password)
		if err != nil {
			return err
		}

		result = &CreateUserDTO{
			Id:    u.ID,
			Name:  u.Name,
			Email: u.Email}
		return nil
	})
	return result, err
}

func (uc *CreateUserUseCase) FindUserByID(ctx context.Context, id string) (*CreateUserDTO, error) {
	uc.Log("Finding user by ID: " + id)

	var result *CreateUserDTO

	err := uc.Tx.WithTransaction(ctx, ports.Read, func(txCtx context.Context) error {
		u, err := uc.service.FindByID(txCtx, id)
		if err != nil {
			return err
		}

		result = &CreateUserDTO{
			Id:    u.ID,
			Name:  u.Name,
			Email: u.Email}
		return nil
	})
	return result, err
}

func (uc *CreateUserUseCase) FindAllUsers(ctx context.Context) ([]*CreateUserDTO, error) {
	uc.Log("Finding all users")

	var result []*CreateUserDTO

	err := uc.Tx.WithTransaction(ctx, ports.Read, func(txCtx context.Context) error {
		users, err := uc.service.FindAll(ctx)
		if err != nil {
			return err
		}

		result = make([]*CreateUserDTO, 0, len(users))
		for _, u := range users {
			result = append(result, &CreateUserDTO{
				Id:    u.ID,
				Name:  u.Name,
				Email: u.Email,
			})
		}
		return nil
	})

	return result, err
}

func (uc *CreateUserUseCase) ChangePassword(ctx context.Context, id, password string) error {
	uc.Log("Changing password for user ID: " + id)

	return uc.Tx.WithTransaction(ctx, ports.Write, func(ctx context.Context) error {
		return uc.service.ChangePassword(ctx, id, password)
	})
}

func (uc *CreateUserUseCase) DeleteUser(ctx context.Context, id string) error {
	uc.Log("Deleting user ID: " + id)

	return uc.Tx.WithTransaction(ctx, ports.Write, func(ctx context.Context) error {
		return uc.service.Delete(ctx, id)
	})
}

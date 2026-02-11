package user

import (
	"context"
	"errors"

	domainuser "erebos/internal/domain/user"
)

type CreateUserUseCase struct {
	svc *domainuser.UserService
}

type CreateUserCommandDTO struct {
	Name     string
	Email    string
	Password string
}

func NewCreateUseCase(svc *domainuser.UserService) *CreateUserUseCase {
	return &CreateUserUseCase{svc: svc}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, cmd *CreateUserCommandDTO) (*domainuser.User, error) {
	if cmd.Name == "" || cmd.Email == "" || cmd.Password == "" {
		return nil, errors.New("name, email and password must not be empty")
	}
	return uc.svc.Create(ctx, cmd.Name, cmd.Email, cmd.Password)
}

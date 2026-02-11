package user

import (
	"context"
	"errors"

	domainuser "erebos/internal/domain/user"
)

type DeleteUserUseCase struct {
	svc *domainuser.UserService
}

type DeleteUserCommandDTO struct {
	ID string
}

func NewDeleteUseCase(svc *domainuser.UserService) *DeleteUserUseCase {
	return &DeleteUserUseCase{svc: svc}
}

func (uc *DeleteUserUseCase) Execute(ctx context.Context, cmd *DeleteUserCommandDTO) error {
	if cmd.ID == "" {
		return errors.New("ID must not be empty")
	}
	return uc.svc.Delete(ctx, cmd.ID)
}

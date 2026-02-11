package user

import (
	"context"

	domainuser "erebos/internal/domain/user"
)

type FindAllUsersUseCase struct {
	svc *domainuser.UserService
}

func NewFindAllUsersUseCase(svc *domainuser.UserService) *FindAllUsersUseCase {
	return &FindAllUsersUseCase{svc: svc}
}

func (uc *FindAllUsersUseCase) Execute(ctx context.Context) ([]*domainuser.User, error) {
	return uc.svc.FindAll(ctx)
}

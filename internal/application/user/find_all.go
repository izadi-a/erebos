package user

import (
	"context"
	domainuser "erebos/internal/domain/user"
)

type FindAllUseCase struct {
	svc *domainuser.UserService
}

func NewFindAllUseCase(svc *domainuser.UserService) *FindAllUseCase {
	return &FindAllUseCase{svc: svc}
}

func (uc *FindAllUseCase) Execute(ctx context.Context) ([]*domainuser.User, error) {
	return uc.svc.FindAll(ctx)
}

package user

import (
	"context"
	domainuser "erebos/internal/domain/user"
)

type FindByIDUseCase struct {
	svc *domainuser.UserService
}

func NewFindByIDUseCase(svc *domainuser.UserService) *FindByIDUseCase {
	return &FindByIDUseCase{svc: svc}
}

func (uc *FindByIDUseCase) Execute(ctx context.Context, id string) (*domainuser.User, error) {
	return uc.svc.FindByID(ctx, id)
}

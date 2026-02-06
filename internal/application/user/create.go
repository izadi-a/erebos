package user

import (
	"context"
	domainuser "erebos/internal/domain/user"
)

type CreateUseCase struct {
	svc *domainuser.UserService
}

func NewCreateUseCase(svc *domainuser.UserService) *CreateUseCase {
	return &CreateUseCase{svc: svc}
}

func (uc *CreateUseCase) Execute(ctx context.Context, name, email, pass string) (*domainuser.User, error) {
	return uc.svc.Create(ctx, name, email, pass)
}

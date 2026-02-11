package user

import (
	"context"

	domainuser "erebos/internal/domain/user"
)

type FindUserByIDUseCase struct {
	svc *domainuser.UserService
}

type FindUserByIdQueryDTO struct {
	ID string
}

func NewFindUserByIDUseCase(svc *domainuser.UserService) *FindUserByIDUseCase {
	return &FindUserByIDUseCase{svc: svc}
}

func (uc *FindUserByIDUseCase) Execute(ctx context.Context, qry FindUserByIdQueryDTO) (*domainuser.User, error) {
	return uc.svc.FindByID(ctx, qry.ID)
}

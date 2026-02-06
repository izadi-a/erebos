package user

import (
	"context"
	domainuser "erebos/internal/domain/user"
)

type DeleteUseCase struct {
	svc *domainuser.UserService
}

func NewDeleteUseCase(svc *domainuser.UserService) *DeleteUseCase {
	return &DeleteUseCase{svc: svc}
}

func (uc *DeleteUseCase) Execute(ctx context.Context, id string) error {
	return uc.svc.Delete(ctx, id)
}

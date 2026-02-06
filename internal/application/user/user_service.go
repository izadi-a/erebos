package user

import (
	"context"

	"erebos/internal/domain/user"
)

type UserService struct {
	repo user.Repository
}

func NewUserService(repo user.Repository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) ChangePassword(ctx context.Context, id, newPassword string) error {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if err := u.ChangePassword(newPassword); err != nil {
		return err
	}
	return s.repo.Update(ctx, u)
}

func (s *UserService) ChangeName(ctx context.Context, id, newName string) error {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	u.ChangeName(newName)
	return s.repo.Update(ctx, u)
}

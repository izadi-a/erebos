package user

import (
	"context"
)

type UserService struct {
	userRepo Repository
}

func NewService(userRepo Repository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) ChangePassword(ctx context.Context, id, newPassword string) error {
	u, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if err := u.ChangePassword(newPassword); err != nil {
		return err
	}
	u.SetModified("admin")
	return s.userRepo.Update(ctx, u)
}

func (s *UserService) ChangeName(ctx context.Context, id, newName string) error {
	u, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	u.ChangeName(newName)
	u.SetModified("admin")
	return s.userRepo.Update(ctx, u)
}

func (s *UserService) Create(ctx context.Context, name, email, password string) (*User, error) {
	u := &User{
		Name:         name,
		Email:        email,
		PasswordHash: password,
	}
	// if err := validateEmail(u.Email); err != nil {
	// 	return nil, err
	// }
	u.SetCreated("admin")
	return u, s.userRepo.Create(ctx, u)
}

func (s *UserService) FindByID(ctx context.Context, id string) (*User, error) {
	return s.userRepo.FindByID(ctx, id)
}

func (s *UserService) FindAll(ctx context.Context) ([]*User, error) {
	return s.userRepo.FindAll(ctx)
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	return s.userRepo.Delete(ctx, id)
}

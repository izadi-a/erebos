package user

import (
	"context"

	"your-module/internal/domain/user"
)

type CreateUser struct {
	repo user.Repository
}

func NewCreateUser(repo user.Repository) *CreateUser {
	return &CreateUser{repo: repo}
}

func (uc *CreateUser) Execute(ctx context.Context, email string) error {
	u := &user.User{
		ID:    "uuid-here",
		Email: email,
	}
	return uc.repo.Save(ctx, u)
}

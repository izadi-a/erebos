package user

import (
	"context"
	"testing"

	"erebos/internal/adapters/outbound/memory"
)

func TestCreateUser(t *testing.T) {
	repo := memory.NewUserRepository()
	uc := NewCreateUser(repo)

	err := uc.Execute(context.Background(), "test@mail.com")
	if err != nil {
		t.Fatal(err)
	}
}

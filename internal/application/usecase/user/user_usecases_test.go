package user_test

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	// repo := memory.NewUserRepository()
	// usecase := app.NewCreateUser(repo)

	// email := "test@example.com"
	// u, err := usecase.Execute(context.Background(), email, "test")
	// if err != nil {
	// 	t.Fatalf("failed to create user: %v", err)
	// }

	// if u.Email != email {
	// 	t.Errorf("expected email %s, got %s", email, u.Email)
	// }

	// بررسی اینکه در Repository ذخیره شده
	// found, err := repo.FindByID(context.Background(), u.ID)
	// if err != nil {
	// 	t.Fatalf("failed to find user: %v", err)
	// }
	// if found == nil {
	// 	t.Fatal("user not found in repository")
	// }
	// if found.Email != email {
	// 	t.Errorf("expected email %s, got %s", email, found.Email)
	// }
}

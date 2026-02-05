package memory

import (
	"context"
	"sync"

	"erebos/internal/domain/user"
)

type UserRepository struct {
	mu    sync.Mutex
	store map[string]*user.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		store: make(map[string]*user.User),
	}
}

func (r *UserRepository) Create(ctx context.Context, u *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.store[u.ID] = u
	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
	return r.store[id], nil
}

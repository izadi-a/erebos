package inbound

import "context"

type CreateUser interface {
	Execute(ctx context.Context, email string) error
}

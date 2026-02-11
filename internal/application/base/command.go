package base

import "context"

type Command interface {
	Execute(ctx context.Context) error
}

type BaseCommand struct {
}

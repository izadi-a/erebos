package ports

import "context"

type TxMode int

const (
	Read TxMode = iota
	Write
)

type TransactionManager interface {
	WithTransaction(ctx context.Context, mode TxMode, fn func(ctx context.Context) error) error
}

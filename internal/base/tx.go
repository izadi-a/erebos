// internal/application/base/tx.go
package base

import (
	"context"
	"erebos/internal/ports"
)

type TxManager interface {
	WithTransaction(ctx context.Context, mode ports.TxMode, fn func(ctx context.Context) error) error
}

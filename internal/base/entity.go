package base

import (
	"erebos/internal/util"
	"time"
)

type BaseEntity struct {
	ID         string
	CreatedAt  time.Time
	CreatedBy  string
	ModifiedAt time.Time
	ModifiedBy string
}

func (b *BaseEntity) SetCreated(userID string) {
	b.ID = util.GenerateID()
	b.CreatedAt = time.Now()
	b.CreatedBy = userID
	b.ModifiedAt = time.Now()
	b.ModifiedBy = userID
}

func (b *BaseEntity) SetModified(userID string) {
	b.ModifiedAt = time.Now()
	b.ModifiedBy = userID
}

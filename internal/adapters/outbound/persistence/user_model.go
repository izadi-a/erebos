package persistence

import (
	"erebos/internal/domain/user"
)

type UserModel struct {
	ID           string `gorm:"primaryKey"`
	Name         string
	Email        string `gorm:"uniqueIndex"`
	PasswordHash string
}

func ToDomain(model *UserModel) *user.User {
	return &user.User{
		ID:           model.ID,
		Name:         model.Name,
		Email:        model.Email,
		PasswordHash: model.PasswordHash,
	}
}

func ToModel(u *user.User) *UserModel {
	return &UserModel{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
	}
}

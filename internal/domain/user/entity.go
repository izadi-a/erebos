package user

import "errors"

type User struct {
	ID           string
	Name         string
	Email        string
	PasswordHash string
}

func (u *User) ChangePassword(newPassword string) error {
	if len(newPassword) < 6 {
		return errors.New("password too short")
	}
	u.PasswordHash = newPassword // ساده سازی، میتونه hash باشه
	return nil
}

func (u *User) ChangeName(newName string) {
	u.Name = newName
}

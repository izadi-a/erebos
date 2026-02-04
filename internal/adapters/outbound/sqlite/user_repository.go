package sqlite

import (
	"database/sql"

	"erebos/internal/domain/user"
)

type SQLiteUserRepository struct {
	db *sql.DB
}

func NewSQLiteUserRepository(db *sql.DB) *SQLiteUserRepository {
	return &SQLiteUserRepository{db: db}
}

func (r *SQLiteUserRepository) Create(user *user.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users(name, email) VALUES (?, ?)",
		user.Name, user.Email,
	)
	return err
}

func (r *SQLiteUserRepository) GetByID(id int64) (*user.User, error) {
	row := r.db.QueryRow(
		"SELECT id, name, email FROM users WHERE id = ?",
		id,
	)

	var u user.User
	if err := row.Scan(&u.ID, &u.Name, &u.Email); err != nil {
		return nil, err
	}
	return &u, nil
}

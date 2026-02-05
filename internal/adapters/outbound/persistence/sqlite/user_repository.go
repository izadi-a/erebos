package sqlite

import (
	"context"
	"database/sql"

	"erebos/internal/domain/user"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteUserRepository struct {
	db *sql.DB
}

func NewSQLiteUserRepository(db *sql.DB) *SQLiteUserRepository {
	return &SQLiteUserRepository{db: db}
}

func (r *SQLiteUserRepository) Create(ctx context.Context, user *user.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users(name, email) VALUES (?, ?)",
		user.Name, user.Email,
	)
	return err
}

func (r *SQLiteUserRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
	row := r.db.QueryRowContext(ctx,
		"SELECT id, name, email FROM users WHERE id = ?",
		id,
	)

	var u user.User
	if err := row.Scan(&u.ID, &u.Name, &u.Email); err != nil {
		return nil, err
	}
	return &u, nil
}

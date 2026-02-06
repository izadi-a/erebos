package sqlite

import (
	"context"
	"database/sql"

	"erebos/internal/domain/user"

	_ "github.com/mattn/go-sqlite3"

	"log"
)

type SQLiteUserRepository struct {
	db     *sql.DB
	logger *log.Logger
}

func NewSQLiteUserRepository(db *sql.DB, logger *log.Logger) *SQLiteUserRepository {
	return &SQLiteUserRepository{db: db, logger: logger}
}

func (r *SQLiteUserRepository) Create(ctx context.Context, user *user.User) error {
	r.logger.Println("Creating user", user.ID)
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (id,name,email,password) VALUES (?,?,?,?)",
		user.ID, user.Name, user.Email, user.PasswordHash)
	if err != nil {
		r.logger.Printf("Error creating user: %v", err)
	}
	return err
}

func (r *SQLiteUserRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id,name,email,password FROM users WHERE id=?", id)
	u := &user.User{}
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *SQLiteUserRepository) FindAll(ctx context.Context) ([]*user.User, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id,name,email,password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*user.User
	for rows.Next() {
		u := &user.User{}
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *SQLiteUserRepository) Update(ctx context.Context, u *user.User) error {
	r.logger.Println("Updating user", u.ID)
	_, err := r.db.ExecContext(ctx, "UPDATE users SET name=?, email=?, password=? WHERE id=?",
		u.Name, u.Email, u.PasswordHash, u.ID)
	return err
}

func (r *SQLiteUserRepository) Delete(ctx context.Context, id string) error {
	r.logger.Println("Deleting user", id)
	_, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id=?", id)
	return err
}

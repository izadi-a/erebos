package sqlite

import (
	"context"

	"gorm.io/gorm"

	"erebos/internal/adapters/outbound/persistence"
	"erebos/internal/domain/user"

	_ "github.com/mattn/go-sqlite3"

	"log"
)

type SQLiteUserRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewSQLiteUserRepository(db *gorm.DB, logger *log.Logger) *SQLiteUserRepository {
	return &SQLiteUserRepository{db: db, logger: logger}
}

func (repository *SQLiteUserRepository) Create(ctx context.Context, user *user.User) error {
	//_, err := r.db.ExecContext(ctx, "INSERT INTO users (id,name,email,password) VALUES (?,?,?,?)",
	// 	user.ID, user.Name, user.Email, user.PasswordHash)
	// if err != nil {
	// 	r.logger.Printf("Error creating user: %v", err)
	// }
	// return err
	repository.logger.Println("Creating user", user.ID)
	return repository.db.WithContext(ctx).Create(persistence.ToModel(user)).Error
}

func (repository *SQLiteUserRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
	// row := repository.db.QueryRowContext(ctx, "SELECT id,name,email,password FROM users WHERE id=?", id)
	// u := &user.User{}
	// if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash); err != nil {
	// 	return nil, err
	// }
	var model persistence.UserModel
	if err := repository.db.WithContext(ctx).First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return persistence.ToDomain(&model), nil
}

func (repository *SQLiteUserRepository) FindAll(ctx context.Context) ([]*user.User, error) {
	// rows, err := repository.db.QueryContext(ctx, "SELECT id,name,email,password FROM users")
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()
	// var users []*user.User
	// for rows.Next() {
	// 	u := &user.User{}
	// 	if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash); err != nil {
	// 		return nil, err
	// 	}
	// 	users = append(users, u)
	// }
	// return users, nil
	var models []persistence.UserModel

	if err := repository.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}

	users := make([]*user.User, len(models))
	for i, m := range models {
		users[i] = persistence.ToDomain(&m)
	}

	return users, nil
}

func (repository *SQLiteUserRepository) Update(ctx context.Context, u *user.User) error {
	// repository.logger.Println("Updating user", u.ID)
	// _, err := repository.db.ExecContext(ctx, "UPDATE users SET name=?, email=?, password=? WHERE id=?",
	// 	u.Name, u.Email, u.PasswordHash, u.ID)
	// return err
	var model persistence.UserModel

	return repository.db.WithContext(ctx).Model(&model).Where("id = ?", u.ID).Updates(persistence.ToModel(u)).Error
}

func (repository *SQLiteUserRepository) Delete(ctx context.Context, id string) error {
	// repository.logger.Println("Deleting user", id)
	// _, err := repository.db.ExecContext(ctx, "DELETE FROM users WHERE id=?", id)
	// return err

	return repository.db.WithContext(ctx).Delete(&persistence.UserModel{}, "id = ?", id).Error
}

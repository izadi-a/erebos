package database

import (
	"database/sql"
	"erebos/internal/adapters/outbound/persistence"

	"gorm.io/gorm"
)

func RunMigrations(db *sql.DB) error {
	query1 := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		name TEXT,
		email TEXT,
		password TEXT
	);`

	if _, err := db.Exec(query1); err != nil {
		return err
	}

	// // ستون جدید name
	// query2 := `
	// ALTER TABLE users ADD COLUMN name TEXT;`

	// // اجرا و نادیده گرفتن خطا اگر ستون از قبل وجود داشته باشد
	// _, _ = db.Exec(query2)

	return nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&persistence.UserModel{})
}

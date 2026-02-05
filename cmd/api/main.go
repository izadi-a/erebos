package main

import (
	"database/sql"
	"log"
	"net/http"

	httpAdapter "erebos/internal/adapters/inbound/http"
	sqliteRepo "erebos/internal/adapters/outbound/persistence/sqlite"

	// redisAdapter "erebos/internal/adapters/outbound/redis"
	app "erebos/internal/application/user"
	"erebos/internal/infrastructure/database"
	"erebos/internal/infrastructure/logger"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Logger
	lg, closeLog, err := logger.New("app.log")
	if err != nil {
		log.Fatal(err)
	}
	defer closeLog()

	// DB (SQLite)
	db, err := sql.Open("sqlite3", "file:app.db?_foreign_keys=on")
	if err != nil {
		lg.Fatal(err)
	}
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	defer db.Close()

	if err := database.RunMigrations(db); err != nil {
		log.Fatal(err)
	}

	// Redis
	// rdb := redisAdapter.New("localhost:6379")
	// if err := redisAdapter.Ping(context.Background(), rdb); err != nil {
	// 	lg.Println("redis not connected:", err)
	// }

	// Wire
	userRepo := sqliteRepo.NewSQLiteUserRepository(db)
	createUser := app.NewCreateUser(userRepo)
	userHandler := httpAdapter.NewUserHandler(createUser)

	// Routes
	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", userHandler.Create)

	lg.Println("server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

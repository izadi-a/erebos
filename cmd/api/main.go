package main

import (
	"database/sql"
	"log"
	"net/http"

	httpAdapter "erebos/internal/adapters/inbound/http"
	sqlite "erebos/internal/adapters/outbound/persistence/sqlite"
	"erebos/internal/infrastructure/database"

	// redisAdapter "erebos/internal/adapters/outbound/redis"
	appuser "erebos/internal/application/user"
	domainuser "erebos/internal/domain/user"
	loginfra "erebos/internal/infrastructure/logger"

	_ "erebos/cmd/api/docs"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Logger
	// lg, closeLog, err := loginfra.New("app.log")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer closeLog()

	logger := loginfra.NewFileLogger("./app.log")

	// DB (SQLite)
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		logger.Fatal(err)
	}
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	defer db.Close()

	if err := database.RunMigrations(db); err != nil {
		log.Fatal(err)
	}

	// // Redis
	// // rdb := redisAdapter.New("localhost:6379")
	// // if err := redisAdapter.Ping(context.Background(), rdb); err != nil {
	// // 	lg.Println("redis not connected:", err)
	// // }

	// --- Outbound Adapters (Repositories) ---
	userRepo := sqlite.NewSQLiteUserRepository(db, logger)
	// taskRepo := sqlite.NewTaskRepository(db, logger)

	// --- Domain Services ---
	userService := domainuser.NewService(userRepo)
	// taskSvc := domaintask.NewService(taskRepo)

	// --- Application UseCases ---
	userCreateUC := appuser.NewCreateUseCase(userService)
	userFindByIDUC := appuser.NewFindByIDUseCase(userService)
	userFindAllUC := appuser.NewFindAllUseCase(userService)
	userDeleteUC := appuser.NewDeleteUseCase(userService)

	// --- Inbound Adapters (HTTP Handlers) ---
	userHandler := httpAdapter.NewUserHandler(userCreateUC, userFindByIDUC, userFindAllUC, userDeleteUC)
	// taskHandler := httpadapter.NewTaskHandler(taskCreateUC, taskFindUC, taskDeleteUC)

	// --- Router (net/http + middleware دستی) ---
	router := httpAdapter.NewRouter(userHandler /*, taskHandler */)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("HTTP server running on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

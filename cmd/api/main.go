package main

import (
	"context"
	"fmt"
	db "inkspire/internal/db/gen"
	"inkspire/internal/handler"
	"inkspire/internal/repository"
	"inkspire/internal/router"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func main() {
	// General ctx
	ctx := context.Background()

	// Load env variables
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// Getting env variables from .env
	db_url := os.Getenv("DB_URL")

	// Creating pgx pool (will be used in new sqlc db)
	pool, err := pgxpool.New(ctx, db_url)
	if err != nil {
		log.Fatal(err)
	}
	// close the pool
	defer pool.Close()

	// This queries will be used in sqlc repos
	queries := db.New(pool)

	// Creating repositories (SQLC Repos) but in handlers we use repo interface
	userRepo := repository.NewUserRepoSQLC(queries)
	postRepo := repository.NewPostRepositorySQLC(queries)
	// Create Handlers
	userHandler := handler.NewUserHandler(userRepo)
	postHandler := handler.NewPostHandler(postRepo)

	// give handlers as paramater to router.New()
	r := router.New(userHandler, postHandler)

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", r)
}

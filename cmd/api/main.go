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

	"github.com/jackc/pgx/v5/pgxpool"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func main() {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, "postgres://postgres:sudosu@localhost:5432/inkspire?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

	queries := db.New(pool)

	// Creating repositories
	userRepo := repository.NewUserRepoSQLC(queries)

	// Create Handlers
	userHandler := handler.NewUserHandler(userRepo)

	// give handlers as paramater to router.New()
	r := router.New(userHandler)

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", r)
}

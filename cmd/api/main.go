package main

import (
	"context"
	"fmt"
	db "inkspire/internal/db/gen"
	"inkspire/internal/handler"
	"inkspire/internal/repository"
	"inkspire/internal/router"
	"inkspire/internal/service"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func main() {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, "postgres://inkspire_user:sudosu@localhost:5432/inkspire?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

	queries := db.New(pool)

	// Creating repositories
	userRepo := repository.NewUserRepoSQLC(queries)

	// Create Services
	userService := service.NewUserService(userRepo)
	postService := service.NewPostService()

	// Create Handlers
	userHandler := handler.NewUserHandler(userService)
	postHandler := handler.NewPostHandler(postService)

	// give handlers as paramater to router.New()
	r := router.New(userHandler, postHandler)

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", r)
}

package main

import (
	"fmt"
	"inkspire/internal/handler"
	"inkspire/internal/router"
	"inkspire/internal/service"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

/**
RESPONSE HELPERS
*/

/*
VALIDATOR HELPERS
*/

func main() {
	// Create Services
	userService := service.NewUserService()
	// Create Handlers
	userHandler := handler.NewUserHandler(userService)

	r := router.New(userHandler)

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", r)
}

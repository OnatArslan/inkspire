package router

import (
	"inkspire/internal/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func New(userHandler *handler.UserHandler) http.Handler {
	r := chi.NewRouter()

	// USER ROUTES
	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
	})

	// POST ROUTES
	r.Route("/posts", func(r chi.Router) {

	})

	return r
}

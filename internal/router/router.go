package router

import (
	"inkspire/internal/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func New(userHandler *handler.UserHandler, postHandler *handler.PostHandler, commentHandler *handler.CommentHandler) http.Handler {
	r := chi.NewRouter()

	// USER ROUTES
	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
	})

	// POST ROUTES
	r.Route("/posts", func(r chi.Router) {
		r.Post("/", postHandler.CreatePost)
		r.Get("/", postHandler.GetAllPosts)
		r.Get("/{uuid}", postHandler.GetUserById)
	})

	// COMMENT ROUTES
	r.Route("/comments", func(r chi.Router) {
		r.Post("/", commentHandler.CreateComment)
	})
	return r
}

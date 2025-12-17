package handler

import (
	"context"
	"encoding/json"

	"inkspire/internal/helper/response"
	"inkspire/internal/model"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// We will use this repo in handlers
type PostRepository interface {
	CreatePost(ctx context.Context, title, content string) (*model.Post, error)
	GetAllPosts(ctx context.Context) ([]model.Post, error)
	GetPostById(ctx context.Context, id string) (*model.Post, error)
}

// Common use is is handler.repo.xyz
type PostHandler struct {
	repo PostRepository
}

func NewPostHandler(repo PostRepository) *PostHandler {
	return &PostHandler{
		repo: repo,
	}
}

type CreatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var req CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	post, err := h.repo.CreatePost(r.Context(), req.Title, req.Content)
	if err != nil {
		response.WriteError(w, http.StatusConflict, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusCreated, *post)
}

func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	// We don't need to decode anything
	users, err := h.repo.GetAllPosts(r.Context())
	if err != nil {
		response.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, users)
}

func (h *PostHandler) GetPostById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	post, err := h.repo.GetPostById(r.Context(), id)
	if err != nil {
		response.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, *post)
}

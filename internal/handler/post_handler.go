package handler

import (
	"encoding/json"
	"inkspire/internal/httpx"
	"inkspire/internal/repository"
	"net/http"
)

type PostHandler struct {
	repo repository.PostRepository
}

func NewPostHandler(repo repository.PostRepository) *PostHandler {
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
		httpx.WriteError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	post, err := h.repo.CreatePost(r.Context(), req.Title, req.Content)
	if err != nil {
		httpx.WriteError(w, http.StatusConflict, err.Error())
		return
	}

	httpx.WriteJSON(w, http.StatusCreated, *post)
}

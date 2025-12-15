package handler

import (
	"encoding/json"
	"errors"
	"inkspire/internal/httpx"
	"inkspire/internal/service"
	"net/http"
)

type PostHandler struct {
	PostService *service.PostService
}

func NewPostHandler(postService *service.PostService) *PostHandler {
	return &PostHandler{
		PostService: postService,
	}
}

type CreatePostRequest struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	PageNum *int   `json:"page_number"`
}

var (
	ErrIdRequired      = errors.New("id can not be empty")
	ErrTitleRequired   = errors.New("title can not be empty")
	ErrInvalidPageNum  = errors.New("invalid page number given")
	ErrPageNumRequired = errors.New("page number can not be empty")
)

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var req CreatePostRequest
	// Decode request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "Invalid json format")
	}

	if req.PageNum == nil {
		httpx.WriteError(w, http.StatusBadRequest, "pageNumber is required")
		return
	}
	if req.Id == "" {
		httpx.WriteError(w, http.StatusBadRequest, "id is required")
		return
	}
	if req.Title == "" {
		httpx.WriteError(w, http.StatusBadRequest, "title is required")
		return
	}

	// Create user via service
	post, err := h.PostService.CreatePost(r.Context(), req.Id, req.Title, *req.PageNum)
	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, err.Error())
	}
	httpx.WriteJSON(w, http.StatusCreated, post)
}

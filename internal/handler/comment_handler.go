package handler

import (
	"context"
	"encoding/json"
	"inkspire/internal/helper/response"
	"inkspire/internal/helper/validator"
	"inkspire/internal/model"
	"net/http"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, content string) (*model.Comment, error)
}

type CommentHandler struct {
	repo CommentRepository
}

func NewCommentHandler(repo CommentRepository) *CommentHandler {
	return &CommentHandler{
		repo: repo,
	}
}

type CreateCommentRequest struct {
	Content string `json:"content" validate:"required,min=5,max=400"`
}

func (h CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var req CreateCommentRequest
	// parse request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	// validate request
	if err := validator.Validate.Struct(req); err != nil {
		response.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	comment, err := h.repo.CreateComment(r.Context(), req.Content)
	if err != nil {
		response.WriteError(w, http.StatusConflict, err.Error())
	}

	response.WriteJSON(w, http.StatusCreated, comment)
}

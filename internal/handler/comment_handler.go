package handler

import (
	"context"
	"encoding/json"
	"inkspire/internal/helper/response"
	"inkspire/internal/helper/validator"
	"inkspire/internal/model"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, content string) (*model.Comment, error)
	GetAllComments(ctx context.Context) ([]model.Comment, error)
	GetCommentById(ctx context.Context, id string) (*model.Comment, error)
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

func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
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
		return
	}

	response.WriteJSON(w, http.StatusCreated, comment)
}

func (h *CommentHandler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.repo.GetAllComments(r.Context())
	if err != nil {
		response.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, comments)
}

func (h *CommentHandler) GetCommentById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	comment, err := h.repo.GetCommentById(r.Context(), id)
	if err != nil {
		response.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, comment)

}

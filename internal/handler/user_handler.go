package handler

import (
	"encoding/json"
	"inkspire/internal/httpx"
	"inkspire/internal/repository"
	"net/http"
)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

/*
USER HANDLERS --- ---
*/

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	email, err := h.repo.Create(r.Context(), req.Email, req.Password)

	if err != nil {
		httpx.WriteError(w, http.StatusConflict, err.Error())
		return
	}

	httpx.WriteJSON(w, http.StatusCreated, email)
}

package handler

import (
	"context"
	"encoding/json"
	"inkspire/internal/helper/response"
	"inkspire/internal/model"
	"net/http"
)

/*
* Real repository is UserRepositorySQLC and this repository implements this interface
* We are doing this because accessing repo(sqlc) directly with handler is bad practice
* We can use this user repo interface on handler
* Layer is like this sqlcRepo -> repo interface -> handler
 */
type UserRepository interface {
	Create(ctx context.Context, email, password string) (string, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

type UserHandler struct {
	repo UserRepository
}

func NewUserHandler(repo UserRepository) *UserHandler {
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
		response.WriteError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	email, err := h.repo.Create(r.Context(), req.Email, req.Password)

	if err != nil {
		response.WriteError(w, http.StatusConflict, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusCreated, email)
}

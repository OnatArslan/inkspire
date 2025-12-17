package repository

import (
	"context"
	db "inkspire/internal/db/gen"
	"inkspire/internal/model"
)

type UserRepoSQLC struct {
	q *db.Queries
}

func NewUserRepoSQLC(q *db.Queries) *UserRepoSQLC {
	return &UserRepoSQLC{q: q}
}

func (r *UserRepoSQLC) Create(ctx context.Context, email, password string) (string, error) {
	email, err := r.q.CreateUser(ctx, db.CreateUserParams{
		Email:    email,
		Password: password,
	})
	return email, err
}

func (r *UserRepoSQLC) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	db_user, err := r.q.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:        db_user.ID.String(),
		Email:     db_user.Email,
		Password:  db_user.Password,
		CreatedAt: db_user.CreatedAt.Time,
	}, nil
}

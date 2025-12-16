package repository

import (
	"context"
	db "inkspire/internal/db/gen"
	"inkspire/internal/model"
)

type PostRepository interface {
	CreatePost(ctx context.Context, title, content string) (*model.Post, error)
}

type PostRepositorySQLC struct {
	q *db.Queries
}

func NewPostRepositorySQLC(q *db.Queries) *PostRepositorySQLC {
	return &PostRepositorySQLC{
		q: q,
	}
}

func (r *PostRepositorySQLC) CreatePost(ctx context.Context, title, content string) (*model.Post, error) {
	db_post, err := r.q.CreatePost(ctx, db.CreatePostParams{
		Title:   title,
		Content: content,
	})
	if err != nil {
		return nil, err
	}

	return &model.Post{
		ID:        db_post.ID.String(),
		Title:     db_post.Title,
		Content:   db_post.Content,
		CreatedAt: db_post.CreatedAt.Time,
	}, nil
}

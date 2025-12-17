package repository

import (
	"context"
	db "inkspire/internal/db/gen"
	"inkspire/internal/model"
)

type CommentRepositorySQLC struct {
	q *db.Queries
}

func NewCommentRepositorySQLC(q *db.Queries) *CommentRepositorySQLC {
	return &CommentRepositorySQLC{
		q: q,
	}
}

func (r *CommentRepositorySQLC) CreateComment(ctx context.Context, content string) (*model.Comment, error) {
	db_comment, err := r.q.CreateComment(ctx, content)
	if err != nil {
		return &model.Comment{}, err
	}

	return &model.Comment{
		Id:         db_comment.ID.String(),
		Content:    db_comment.Content,
		Created_at: db_comment.CreatedAt.Time,
	}, nil

}

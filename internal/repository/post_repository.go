package repository

import (
	"context"
	db "inkspire/internal/db/gen"
	"inkspire/internal/model"

	"github.com/jackc/pgx/v5/pgtype"
)

type PostRepositorySQLC struct {
	q *db.Queries
}

func NewPostRepositorySQLC(q *db.Queries) *PostRepositorySQLC {
	return &PostRepositorySQLC{
		q: q,
	}
}

type PostRepository interface {
	CreatePost(ctx context.Context, title, content string) (*model.Post, error)
	GetAllPosts(ctx context.Context) ([]model.Post, error)
	GetPostById(ctx context.Context, id string) (*model.Post, error)
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

func (r *PostRepositorySQLC) GetAllPosts(ctx context.Context) ([]model.Post, error) {
	db_posts, err := r.q.GetAllPosts(ctx)
	if err != nil {
		return nil, err
	}
	var posts = make([]model.Post, len(db_posts))

	for _, v := range db_posts {
		post := model.Post{
			ID:        v.ID.String(),
			Title:     v.Title,
			Content:   v.Content,
			CreatedAt: v.CreatedAt.Time,
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *PostRepositorySQLC) GetPostById(ctx context.Context, id string) (*model.Post, error) {
	var uuid pgtype.UUID
	if err := uuid.Scan(id); err != nil {
		return &model.Post{}, err
	}

	db_post, err := r.q.GetPostById(ctx, uuid)
	if err != nil {
		return &model.Post{}, err
	}

	return &model.Post{
		ID:        db_post.ID.String(),
		Title:     db_post.Title,
		Content:   db_post.Content,
		CreatedAt: db_post.CreatedAt.Time,
	}, nil
}

package service

import (
	"context"
	"inkspire/internal/model"
)

type PostService struct{}

func NewPostService() *PostService {
	return &PostService{}
}

func (s *PostService) CreatePost(ctx context.Context, id, title string, pageNum int) (*model.Post, error) {

	return &model.Post{
		Id:      id,
		Title:   title,
		PageNum: pageNum,
	}, nil
}

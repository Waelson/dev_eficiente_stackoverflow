package service

import (
	"context"
	"github.com/Waelson/internal/model"
	"github.com/Waelson/internal/repository"
	"github.com/Waelson/internal/util/api"
	"time"
)

type PostService interface {
	Save(ctx context.Context, post *model.Post) (*model.Post, api.Error)
}

type authorService struct {
	postRepository      repository.PostRepository
	searchEngineService SearchEngineService
}

func (a *authorService) Save(ctx context.Context, post *model.Post) (*model.Post, api.Error) {
	post.Status = model.PostStatusOpened
	post.CreateAt = time.Now()

	_, err := a.postRepository.Save(ctx, post)

	if err != nil {
		return nil, err
	}

	go func() {
		a.searchEngineService.Index(ctx, post)
	}()

	return post, nil
}

func NewPostService(postRepository repository.PostRepository, searchEngineService SearchEngineService) PostService {
	return &authorService{
		postRepository:      postRepository,
		searchEngineService: searchEngineService,
	}
}

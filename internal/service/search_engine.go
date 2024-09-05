package service

import (
	"context"
	"github.com/Waelson/internal/model"
)

type SearchEngineService interface {
	Index(ctx context.Context, post *model.Post) error
}

type searchEngineService struct {
}

func (s *searchEngineService) Index(ctx context.Context, post *model.Post) error {
	return nil
}

func NewSearchEngineService() SearchEngineService {
	return &searchEngineService{}
}
